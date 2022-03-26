FROM golang:1.18 AS builder


WORKDIR /build/bin/
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

RUN make build

FROM scratch

WORKDIR /tmp/myapp/

WORKDIR /config
WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder ["/build/bin/myapp", "/"]

COPY --from=builder ["/build/config/config.yaml", "/config/"]


# Export necessary port.
EXPOSE 34567


# Command to run when starting the container.
CMD ["./myapp"]