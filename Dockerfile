# builder image
FROM golang:1.14 as builder
LABEL maintainer="mav-MWP-Engg-All@mavenir.com"
WORKDIR /yala
ADD . .
RUN pwd && ls -lha
RUN go env
RUN go mod vendor
RUN go build

 # helm image
FROM alpine:3.13.0 as helm
WORKDIR /yala
CMD pwd && ls -la /yala
RUN pwd && ls -la /yala
RUN apk add curl && curl -LO https://get.helm.sh/helm-v3.4.2-linux-amd64.tar.gz
RUN tar -zxvf helm-v3.4.2-linux-amd64.tar.gz

 # final image
FROM alpine:3.13.0
WORKDIR /yala
CMD pwd && ls -la /yala
RUN pwd && ls -la /yala
COPY --from=builder /yala/yala /usr/local/bin/yala
COPY --from=helm /yala/linux-amd64/helm /usr/local/bin/helm
RUN apk add --update docker openrc
RUN rc-update add docker boot
ENTRYPOINT ["yala"]
