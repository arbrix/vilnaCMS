# Use the official Go image as the base image
FROM golang:alpine

# set timezone to avoid postgres issues
RUN apk add -U tzdata
ENV TZ=America/Santiago
RUN cp /usr/share/zoneinfo/America/Santiago /etc/localtime

# Set the working directory inside the container
WORKDIR /go/src/github.com/arbrix/vilnaCMS

# Copy the Go application files to the container
COPY .. .

# Install dependencies (if any)
RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go", "--pathToDBConf=local/conf.json"]
# RUN go build -o /goapp
#CMD ["/goapp", "--pathToDBConf=/go/src/github.com/arbrix/vilnaCMS/local/conf.json"]
