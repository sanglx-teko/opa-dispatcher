FROM golang:1.12 as builder 
LABEL maintainer="Sang Li <sang.lx@teko.vn>"
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/opa-dispatcher .

######## Start a new stage from scratch #######
FROM alpine:latest
WORKDIR /opa-dispatcher-app
COPY --from=builder /usr/local/bin/opa-dispatcher .
RUN chmod +x opa-dispatcher
COPY configs.json .

EXPOSE 1323

CMD [ "./opa-dispatcher" ]