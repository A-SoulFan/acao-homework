FROM golang:1.17-alpine3.15 as builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

WORKDIR /work

COPY go.mod go.sum /work/

ENV GOPROXY=https://goproxy.io,direct

RUN go mod download

COPY ./ /work

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -ldflags '-w -s' -o support_api ./cmd/support_api/main.go

FROM alpine as prod

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

COPY --from=builder /work/support_api ./

CMD ["./support_api"]
