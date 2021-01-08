package main

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

type DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

type ClientConfig struct {
	HTTPClient    *HTTPClientConfig
	HTTPTransport *HTTPTransportConfig
	TLSConfig     *TLSConfig
	NetDialer     *NetDialerConfig
}

type HTTPClientConfig struct {
	Timeout time.Duration
}

type HTTPTransportConfig struct {
	TLSHandshakeTimeout    time.Duration
	DisableKeepAlives      bool
	MaxIdleConns           int
	MaxIdleConnsPerHost    int
	MaxConnsPerHost        int
	IdleConnTimeout        time.Duration
	ResponseHeaderTimeout  time.Duration
	ExpectContinueTimeout  time.Duration
	MaxResponseHeaderBytes int64
	WriteBufferSize        int
	ReadBufferSize         int
	HTTP2Enabled           bool
}

type TLSConfig struct {
	PreferServerCipherSuites bool
}

type NetDialerConfig struct {
	Timeout   time.Duration
	KeepAlive time.Duration
}

var config = &ClientConfig{
	HTTPClient: &HTTPClientConfig{
		Timeout: 0 * time.Millisecond,
	},
	HTTPTransport: &HTTPTransportConfig{
		TLSHandshakeTimeout:    0 * time.Millisecond,
		DisableKeepAlives:      false,
		MaxIdleConns:           0,
		MaxIdleConnsPerHost:    2,
		MaxConnsPerHost:        2,
		IdleConnTimeout:        10 * time.Second,
		ResponseHeaderTimeout:  0 * time.Millisecond,
		ExpectContinueTimeout:  0 * time.Millisecond,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		HTTP2Enabled:           false,
	},
	TLSConfig: &TLSConfig{
		PreferServerCipherSuites: false,
	},
	NetDialer: &NetDialerConfig{
		Timeout:   0 * time.Millisecond,
		KeepAlive: 0 * time.Millisecond,
	},
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	client := buildHTTPClient(config)

	var waitGroup sync.WaitGroup

	for i := 0; i < 2; i++ {
		waitGroup.Add(1)

		contextLogger := log.WithFields(log.Fields{
			"requesterId": i,
		})

		go func(logger *log.Entry) {
			defer waitGroup.Done()

			for j := 0; j < 2; j++ {

				startTime := time.Now()

				err := ping(client)

				stopTime := time.Now()
				elapsedTime := stopTime.Sub(startTime)

				logger.Printf("Start: [%s], Stop: [%s], Elapsed: [%s], Err: [%v]\n", startTime, stopTime, elapsedTime, err)

				time.Sleep(5 * time.Second)
			}
		}(contextLogger)
	}

	waitGroup.Wait()
}

func buildHTTPClient(config *ClientConfig) *http.Client {
	dialContext := newDialContext(config.NetDialer)
	tlsClientConfig := newTLSClientConfig(config.TLSConfig)
	transport := newHTTPTransport(dialContext, tlsClientConfig, config.HTTPTransport)
	return newHTTPClient(transport, config.HTTPClient)
}

func newDialContext(config *NetDialerConfig) DialContext {
	return (&net.Dialer{
		Timeout:   config.Timeout,
		KeepAlive: config.KeepAlive,
	}).DialContext
}

func newTLSClientConfig(config *TLSConfig) *tls.Config {
	cfg := &tls.Config{
		PreferServerCipherSuites: config.PreferServerCipherSuites,
	}
	return cfg
}

func newHTTPTransport(dialContext DialContext, tlsClientConfig *tls.Config, config *HTTPTransportConfig) http.RoundTripper {
	httpTransport := &http.Transport{
		Proxy:                  http.ProxyFromEnvironment,
		DialContext:            dialContext,
		TLSClientConfig:        tlsClientConfig,
		TLSHandshakeTimeout:    config.TLSHandshakeTimeout,
		DisableKeepAlives:      config.DisableKeepAlives,
		MaxIdleConns:           config.MaxIdleConns,
		MaxIdleConnsPerHost:    config.MaxIdleConnsPerHost,
		MaxConnsPerHost:        config.MaxConnsPerHost,
		IdleConnTimeout:        config.IdleConnTimeout,
		ResponseHeaderTimeout:  config.ResponseHeaderTimeout,
		ExpectContinueTimeout:  config.ExpectContinueTimeout,
		MaxResponseHeaderBytes: config.MaxResponseHeaderBytes,
		WriteBufferSize:        config.WriteBufferSize,
		ReadBufferSize:         config.ReadBufferSize,
	}

	if !config.HTTP2Enabled {
		return httpTransport
	}

	err := http2.ConfigureTransport(httpTransport)
	if err != nil {
		panic(err)
	}

	return httpTransport
}

func newHTTPClient(transport http.RoundTripper, config *HTTPClientConfig) *http.Client {
	return &http.Client{
		Transport: transport,
		Timeout:   config.Timeout,
	}
}

func ping(client *http.Client) error {
	resp, err := client.Get("http://localhost:9090/ping")

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return nil
}
