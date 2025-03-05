FROM golang:latest
WORKDIR /app
COPY ./svr/container/ .
RUN go get -u google.golang.org/grpc
RUN go mod tidy
COPY ./svr/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
# ENTRYPOINT ["go","run","main.go"]
ENTRYPOINT [ "sh","./entrypoint.sh" ]