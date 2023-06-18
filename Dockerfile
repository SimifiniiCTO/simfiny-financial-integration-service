FROM golang:1.20-alpine as builder

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

RUN chown -R app:app ./

USER app

CMD ["./podinfo"]
