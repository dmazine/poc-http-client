## http.client

### Timeout

## http.transport

### MaxConnsPerHost

Documentation:
```
MaxConnsPerHost optionally limits the total number of  connections per host, including connections in the dialing,
active, and idle states. On limit violation, dials will block. Zero means no limit.
```

Tests:

It works like a standard pool. Each go routine will get a connection from the pool, send the request, read the response
and release the connection back to the pool.

Test #1:

Configuration:
```
MaxConnsPerHost: 1
4 go routines
1 request / go routine
```

Result:
```
INFO[2021-01-08T07:44:15-03:00] Start: [2021-01-08 07:44:05.808962 -0300 -03 m=+0.000593718], Stop: [2021-01-08 07:44:15.81745 -0300 -03 m=+10.008989257], Elapsed: [10.008395539s], Err: [<nil>]  requesterId=3
INFO[2021-01-08T07:44:25-03:00] Start: [2021-01-08 07:44:05.809051 -0300 -03 m=+0.000682864], Stop: [2021-01-08 07:44:25.825225 -0300 -03 m=+20.016671437], Elapsed: [20.015988573s], Err: [<nil>]  requesterId=2
INFO[2021-01-08T07:44:35-03:00] Start: [2021-01-08 07:44:05.808979 -0300 -03 m=+0.000610771], Stop: [2021-01-08 07:44:35.836619 -0300 -03 m=+30.027973441], Elapsed: [30.02736267s], Err: [<nil>]  requesterId=1
INFO[2021-01-08T07:44:45-03:00] Start: [2021-01-08 07:44:05.808951 -0300 -03 m=+0.000582788], Stop: [2021-01-08 07:44:45.843826 -0300 -03 m=+40.035087504], Elapsed: [40.034504716s], Err: [<nil>]  requesterId=0
```

Test #2:

Configuration:
```
MaxConnsPerHost: 2
4 go routines
1 request / go routine
```

Result:
```
INFO[2021-01-08T07:45:41-03:00] Start: [2021-01-08 07:45:31.040509 -0300 -03 m=+0.000539523], Stop: [2021-01-08 07:45:41.046845 -0300 -03 m=+10.006783440], Elapsed: [10.006243917s], Err: [<nil>]  requesterId=3
INFO[2021-01-08T07:45:41-03:00] Start: [2021-01-08 07:45:31.040519 -0300 -03 m=+0.000549715], Stop: [2021-01-08 07:45:41.05151 -0300 -03 m=+10.011448402], Elapsed: [10.010898687s], Err: [<nil>]  requesterId=0
INFO[2021-01-08T07:45:51-03:00] Start: [2021-01-08 07:45:31.040508 -0300 -03 m=+0.000538737], Stop: [2021-01-08 07:45:51.051392 -0300 -03 m=+20.011236784], Elapsed: [20.010698047s], Err: [<nil>]  requesterId=1
INFO[2021-01-08T07:45:51-03:00] Start: [2021-01-08 07:45:31.040512 -0300 -03 m=+0.000542994], Stop: [2021-01-08 07:45:51.055338 -0300 -03 m=+20.015183477], Elapsed: [20.014640483s], Err: [<nil>]  requesterId=2
```

Test #3:

Configuration:
```
MaxConnsPerHost: 4
4 go routines
1 request / go routine
```

Result:
```
INFO[2021-01-08T07:47:49-03:00] Start: [2021-01-08 07:47:39.444648 -0300 -03 m=+0.000652179], Stop: [2021-01-08 07:47:49.455671 -0300 -03 m=+10.011582883], Elapsed: [10.010930704s], Err: [<nil>]  requesterId=2
INFO[2021-01-08T07:47:49-03:00] Start: [2021-01-08 07:47:39.444537 -0300 -03 m=+0.000541534], Stop: [2021-01-08 07:47:49.456227 -0300 -03 m=+10.012139169], Elapsed: [10.011597635s], Err: [<nil>]  requesterId=0
INFO[2021-01-08T07:47:49-03:00] Start: [2021-01-08 07:47:39.444582 -0300 -03 m=+0.000586757], Stop: [2021-01-08 07:47:49.456694 -0300 -03 m=+10.012605965], Elapsed: [10.012019208s], Err: [<nil>]  requesterId=1
INFO[2021-01-08T07:47:49-03:00] Start: [2021-01-08 07:47:39.444548 -0300 -03 m=+0.000552800], Stop: [2021-01-08 07:47:49.456365 -0300 -03 m=+10.012277291], Elapsed: [10.011724491s], Err: [<nil>]  requesterId=3
```

### MaxIdleConns

Documentation:
```
MaxIdleConns controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.
```

Test #1:
```
MaxConnsPerHost: 2
MaxIdleConnsPerHost: 1
IdleConnTimeout: 10 * time.Second,
2 go routines
2 requests / go routines with sleep of 5 secongs between each one
```

Results:
``` 
INFO[2021-01-08T13:56:22-03:00] Start: [2021-01-08 13:56:12.75651 -0300 -03 m=+0.000575922], Stop: [2021-01-08 13:56:22.779403 -0300 -03 m=+10.023285118], Elapsed: [10.022709196s], Err: [<nil>]  requesterId=0
INFO[2021-01-08T13:56:22-03:00] Start: [2021-01-08 13:56:12.756509 -0300 -03 m=+0.000575069], Stop: [2021-01-08 13:56:22.779435 -0300 -03 m=+10.023316873], Elapsed: [10.022741804s], Err: [<nil>]  requesterId=1
INFO[2021-01-08T13:56:37-03:00] Start: [2021-01-08 13:56:27.785193 -0300 -03 m=+15.028983078], Stop: [2021-01-08 13:56:37.790927 -0300 -03 m=+25.034532367], Elapsed: [10.005549289s], Err: [<nil>]  requesterId=0
INFO[2021-01-08T13:56:37-03:00] Start: [2021-01-08 13:56:27.785103 -0300 -03 m=+15.028893025], Stop: [2021-01-08 13:56:37.794343 -0300 -03 m=+25.037948811], Elapsed: [10.009055786s], Err: [<nil>]  requesterId=1
```

At the first, 2 connections are opened to handle the requests.
```
tcp6       0      0  ::1.62576              ::1.9090               ESTABLISHED
tcp6       0      0  ::1.62575              ::1.9090               ESTABLISHED
tcp46      0      0  *.9090                 *.*                    LISTEN   
```

It is possible to confirm that one of them is closed, and the other kept opened after the first request of each go
routine completes.
```
tcp6       0      0  ::1.62576              ::1.9090               ESTABLISHED
```

The remaining connection is reused to send the second request of one requester, and a new one is created to send the
second request of the other requester.

```
tcp6       0      0  ::1.62587              ::1.9090               ESTABLISHED
tcp6       0      0  ::1.62576              ::1.9090               ESTABLISHED
```

After both are completed, one connection is closed the other kept opened (idle) until the idle timeout expired.

```
tcp6       0      0  ::1.62576              ::1.9090               ESTABLISHED
```

Test #2:
```
MaxConnsPerHost: 2
MaxIdleConnsPerHost: 2
IdleConnTimeout: 10 * time.Second,
2 go routines
2 requests / go routines with sleep of 5 secongs between each one
```

Results:
```
INFO[2021-01-08T14:21:56-03:00] Start: [2021-01-08 14:21:46.516237 -0300 -03 m=+0.000713981], Stop: [2021-01-08 14:21:56.524434 -0300 -03 m=+10.008829360], Elapsed: [10.008115379s], Err: [<nil>]  requesterId=1
INFO[2021-01-08T14:21:56-03:00] Start: [2021-01-08 14:21:46.516234 -0300 -03 m=+0.000710903], Stop: [2021-01-08 14:21:56.527243 -0300 -03 m=+10.011638530], Elapsed: [10.010927627s], Err: [<nil>]  requesterId=0
INFO[2021-01-08T14:22:11-03:00] Start: [2021-01-08 14:22:01.527663 -0300 -03 m=+15.012006219], Stop: [2021-01-08 14:22:11.533241 -0300 -03 m=+25.017467740], Elapsed: [10.005461521s], Err: [<nil>]  requesterId=0
INFO[2021-01-08T14:22:11-03:00] Start: [2021-01-08 14:22:01.52601 -0300 -03 m=+15.010353640], Stop: [2021-01-08 14:22:11.533289 -0300 -03 m=+25.017514912], Elapsed: [10.007161272s], Err: [<nil>]  requesterId=1
```

In this case it is possible to confirm that 2 connections are created to handle the requests.
```
tcp6       0      0  ::1.63127              ::1.9090               ESTABLISHED
tcp6       0      0  ::1.63126              ::1.9090               ESTABLISHED
```

Both of them are kept opened and result to handle the next to requests.
```
tcp6       0      0  ::1.63127              ::1.9090               ESTABLISHED
tcp6       0      0  ::1.63126              ::1.9090               ESTABLISHED
```

Once the requests are complete, both of them are kept idle until the IdleConnTimeout expires. After that they are closed.

### MaxIdleConnsPerHost

Documentation:
```
MaxIdleConnsPerHost, if non-zero, controls the maximum idle (keep-alive) connections to keep per-host.
If zero, DefaultMaxIdleConnsPerHost is used.
```

### IdleConnTimeout

Documentation:
```
IdleConnTimeout is the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself.
Zero means no limit.
```

### TLSHandshakeTimeout
### DisableKeepAlives
### ResponseHeaderTimeout
### ExpectContinueTimeout
### MaxResponseHeaderBytes
### WriteBufferSize
### ReadBufferSize
### HTTP2Enabled

## tls.Config
### PreferServerCipherSuites

## net.Dialer

### Timeout
### KeepAlive
