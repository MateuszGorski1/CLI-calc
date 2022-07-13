FROM golang:latest AS builder
ENV GOPROXY="https://repo.cci.nokia.net/proxy-golang-org"
ENV GOSUMDB=off
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o calc -ldflags "-w -extldflags '-static'" .

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/calc .
USER 65535:65536
ENTRYPOINT ["/calc"]