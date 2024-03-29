# Build the manager binary
FROM golang:1.20 as builder

WORKDIR /workspace
COPY . .
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o webinar main.go


FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/webinar .
USER 65532:65532
ENTRYPOINT ["/webinar"]
