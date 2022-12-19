FROM golang:1.19

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN go build -o myapp

#ENV PORT 8080

EXPOSE 8080

ENTRYPOINT ["./myapp"]