FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x
COPY . .

RUN go build -o ./bin/auth cmd/auth/main.go
RUN go build -o ./bin/migrator cmd/migrator/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/auth /
COPY --from=builder /usr/local/src/bin/migrator /

# COPY config/.env /.env
COPY scripts/startup.sh /
COPY migrations /migrations

EXPOSE 50051

CMD ["/startup.sh"]