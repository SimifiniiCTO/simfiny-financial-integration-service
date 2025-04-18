#!/usr/bin/env bash

if ! command -v dlv &> /dev/null
then
    echo "[builder] delve could not be found, it will be installed"
    go install github.com/go-delve/delve/cmd/dlv@latest
else
    echo "[builder] delve already installed, skipping..."
fi

NOW=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
REVISION=$(git rev-parse HEAD)

LDFLAGS=""
LDFLAGS="${LDFLAGS} -X main.buildType=development"
LDFLAGS="${LDFLAGS} -X main.buildHost=${BUILD_HOST:-`hostname`}"
LDFLAGS="${LDFLAGS} -X main.buildTime=${NOW}"
LDFLAGS="${LDFLAGS} -X main.buildRevision=${REVISION}"

TAGS="development,mini,noui,icons"
[ -d "pkg/icons/sources/simple-icons" ] && TAGS="${TAGS},simple_icons"

echo "[builder] building the financial integration service now with tags (${TAGS})..."
go build -ldflags "${LDFLAGS}" -tags=${TAGS} -o /usr/bin/simfiny github.com/SimifiniiCTO/simfiny-financial-integration-service/cmd/podinfo
