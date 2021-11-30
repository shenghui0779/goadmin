FROM golang:1.16.3-alpine as builder

WORKDIR /go/goadmin
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOFLAGS=-mod=vendor go build -o goadmin main.go

# =============================================================================
FROM centos AS final

# RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app/
COPY . .
COPY --from=builder /go/goadmin/goadmin .
COPY --from=builder /go/goadmin/yiigo.toml .
COPY --from=builder /go/goadmin/favicon.ico .

COPY  ./views/ ./views/
COPY ./assets/ ./assets/

RUN chmod +x goadmin

# ENTRYPOINT ["/app/goadmin"]
ENTRYPOINT ["top"]
