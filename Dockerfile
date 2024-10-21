# Stage 1
FROM quay.io/fedora/fedora AS builder

RUN dnf install -y golang

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -o main cmd/main.go

# Stage 2
FROM quay.io/fedora/fedora

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
ENTRYPOINT [ "./main" ]