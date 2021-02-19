# builder image
FROM golang:1.14 as builder
LABEL maintainer="mav-MWP-Engg-All@mavenir.com"
WORKDIR /src
ADD . .
COPY main.go go.* /src
RUN CGO_ENABLED=0 go build

# helm image
FROM alpine:3.13.0 as helm
WORKDIR /src
RUN apk add curl && curl -LO https://get.helm.sh/helm-v3.4.2-linux-amd64.tar.gz
RUN tar -zxvf helm-v3.4.2-linux-amd64.tar.gz

# final image
FROM alpine:3.13.0
WORKDIR /src
COPY --from=builder /src/yala /usr/local/bin/yala
COPY --from=helm /src/linux-amd64/helm /usr/local/bin/helm
RUN apk add --update docker openrc
RUN rc-update add docker boot
ENTRYPOINT ["yala"]
