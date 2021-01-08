#! /bin/bash

# Get the script directory
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
BASE_DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Java command
JAVA=java

# Wire mock directory
WIREMOCK_DIR="${BASE_DIR}/third_party/wiremock"

# Wire mock port
WIREMOCK_PORT="9090"

# Wire mock HTTPS port
WIREMOCK_HTTPS_PORT="9443"

# Execute wiremock
"${JAVA}" -jar "${WIREMOCK_DIR}/wiremock-jre8-standalone-2.27.2.jar" --root-dir "${WIREMOCK_DIR}" --port "${WIREMOCK_PORT}" --https-port "${WIREMOCK_HTTPS_PORT}" --async-response-enabled true
