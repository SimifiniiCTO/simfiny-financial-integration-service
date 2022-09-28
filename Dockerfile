FROM golang:1.18-alpine as builder

ARG REVISION

RUN apk update

RUN apk add git

RUN git config --global url."https://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:x-oauth-basic@github.com/SimifiniiCTO".insteadOf "https://github.com/SimifiniiCTO"

RUN mkdir -p /podinfo/

WORKDIR /podinfo

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags \
    						"-s -w -X github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/version.REVISION=${REVISION}" \
							-a -o bin/podinfo cmd/podinfo/*

RUN CGO_ENABLED=0 go build -ldflags \
    						"-s -w -X github.com/SimifiniiCTO/simfinii/src/backend/services//financial-integration-service/pkg/version.REVISION=${REVISION}" \
							-a -o bin/podcli cmd/podcli/*

FROM alpine

ARG BUILD_DATE
ARG VERSION
ARG REVISION

LABEL maintainer="yoanyomba"

RUN addgroup -S app \
    && adduser -S -G app app \
    && apk --no-cache add \
    ca-certificates curl netcat-openbsd

WORKDIR /home/app

COPY --from=builder /podinfo/bin/podinfo .
COPY --from=builder /podinfo/bin/podcli /usr/local/bin/podcli
COPY ./proto ./proto
COPY ./ui ./ui
COPY ./boot.yaml .
COPY ./buf.yaml .
COPY ./buf.gen.yaml .

RUN set -ex && apk --no-cache add sudo
RUN apk update && apk add bash

RUN curl -Ls https://download.newrelic.com/install/newrelic-cli/scripts/install.sh | bash && sudo NEW_RELIC_API_KEY=NRAK-FGBHISKJUMWR9DXMIV4Q4EFZJTK NEW_RELIC_ACCOUNT_ID=3270596 /usr/local/bin/newrelic install -n logs-integration
RUN chown -R app:app ./

USER app

CMD ["./podinfo"]
