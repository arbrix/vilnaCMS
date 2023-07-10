# Source code + build dependencies base image
FROM golang:1.20-alpine AS build

# set timezone to avoid postgres issues
RUN apk add -U tzdata
ENV TZ=Europe/Kiev
RUN cp /usr/share/zoneinfo/Europe/Kiev /etc/localtime

# Set the working directory inside the container
WORKDIR /go/src/vilnaCMS


# Copy the Go application files to the container
COPY . .

# Install dependencies (if any)
RUN go mod download && go mod verify

# build binary
RUN go env -w CGO_ENABLED="0"
RUN go build -ldflags="-s -w" -o /bin/app

# Dev image with live reloading + debugger
FROM build AS dev

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]

## Serve image
FROM alpine:latest AS serve

# Set timezone to avoid postgres issues
RUN apk add -U tzdata
ENV TZ=Europe/Kiev
RUN cp /usr/share/zoneinfo/Europe/Kiev /etc/localtime

# Copy binary
COPY --from=build /bin/app /app

# ----------------------
# Add specific db config if needed
# ----------------------
COPY --from=build /go/src/vilnaCMS/local/conf.json /local/conf.json

ENTRYPOINT ["/app"]
