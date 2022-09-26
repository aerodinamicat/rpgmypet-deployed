FROM golang:1.18

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-extldflags "-static"' -o server ./cmd/main.go

EXPOSE 8080

ENTRYPOINT [ "./server" ]