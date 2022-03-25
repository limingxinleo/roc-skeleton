FROM golang:1.18 as builder

LABEL maintainer="limx <l@hyperf.io>"
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/builder

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o main main.go

FROM scratch

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/builder /

EXPOSE 9501

ENTRYPOINT ["/main"]