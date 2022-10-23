#!/bin/bash

BUILD_IMAGE="golang:1.17"
RUN_IMAGE="alpine:3.15.0"
printf "Build image: %s\n" "$BUILD_IMAGE"
printf "Run image: %s\n" "$RUN_IMAGE"

cat > Dockerfile <<- EOF
FROM ${BUILD_IMAGE} as build
WORKDIR /build
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -o ./bin/consumer-sarama ./consumer-sarama
RUN CGO_ENABLED=0 go build -o ./bin/consumer-sarama-console ./consumer-sarama-console
RUN CGO_ENABLED=0 go build -o ./bin/consumer-segmentio ./consumer-segmentio
RUN CGO_ENABLED=0 go build -o ./bin/producer-http ./producer-http
RUN CGO_ENABLED=0 go build -o ./bin/producer-segmentio ./producer-segmentio

FROM ${RUN_IMAGE}
RUN apk add --update --no-cache \
	bash \
	nano \
	jq \
	tcpdump \
	curl \
	kafkacat

# bash with color and user@host:path/$ 
RUN echo "PS1='\[\e]0;${debian_chroot:+($debian_chroot)}\u@\H: \w\a\]${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\H\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '" >> ~/.bashrc

WORKDIR /app
COPY --from=build /build/bin/ ./
COPY ./scripts/ ./
RUN chmod +x *.sh

ENTRYPOINT ["bash"]
EOF

printf "Build image...\n"
IMAGE=$(docker build -q .)

SERVICE="fms_kafka"
CONTAINER=$(docker ps -f name="$SERVICE" --format "{{.Names}}")
printf "Run with namespace: %s\n" "$CONTAINER"
docker run -it --rm --net container:"$CONTAINER" "$IMAGE"