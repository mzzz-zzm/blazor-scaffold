FROM golang:latest
WORKDIR /app
COPY ./svr/container/ .
RUN go install google.golang.org/grpc
RUN go mod tidy
COPY ./svr/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

# ENTRYPOINT [ "sh","./entrypoint.sh" ]
ENTRYPOINT [ "tail", "-f", "/dev/null" ]