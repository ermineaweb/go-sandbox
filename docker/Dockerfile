FROM golang:1.21.5 as builder
WORKDIR /build
COPY ./go.* ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -o ./bin/main ./src

FROM scratch
WORKDIR /app
COPY --from=builder /build/bin/main ./main
ENTRYPOINT ["./main"]
