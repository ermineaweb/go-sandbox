FROM golang:1.17 as BUILD
WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -o main ./src 

FROM alpine:3.7
WORKDIR /app
COPY --from=BUILD /app/main ./main
COPY ./public ./public

ENTRYPOINT ["/app/main"]