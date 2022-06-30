FROM golang:latest

LABEL maintainer="Simant Thapa Magar <tmsimant96@gmail.com>"

WORKDIR /app

COPY go.mod ./
COPY *.* ./
RUN go build -o /go-cli-video
ENTRYPOINT [ "/go-cli-video"]
CMD ["get","--all"]
