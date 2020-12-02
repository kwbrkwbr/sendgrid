FROM golang:1.15.2 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
#ENV GOPROXY="https://proxy.golang.org"

WORKDIR /go/src/sendgrid
ENV GO111MODULE=on

# 先にvendorを作成する
COPY go.mod ./
RUN go mod download

COPY . .
RUN make build

FROM alpine
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /go/src/sendgrid/main /app/main
CMD ["/app/main"]
