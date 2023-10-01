FROM golang:1.20 as base

WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build cmd/api-server/main.go

CMD [ "./main" ]