FROM golang:1.16.5 AS builder

WORKDIR /goadmin

COPY . .

RUN go env -w GOPROXY="https://goproxy.cn"
RUN go mod tidy
RUN sh ent.sh

RUN CGO_ENABLED=0 go build -o ./bin/main ./cmd

FROM scratch

WORKDIR /data

COPY --from=builder /goadmin/bin/main .
COPY --from=builder /goadmin/cmd/assets ./assets
COPY --from=builder /goadmin/cmd/html ./html

EXPOSE 8000

ENTRYPOINT ["./main"]

CMD ["--envfile", "/data/config/.env"]