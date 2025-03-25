FROM golang:latest
WORKDIR /app
EXPOSE 50051
COPY ./svr/container/ .
RUN go mod tidy
COPY ./svr/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
# ENTRYPOINT ["go","run","main.go"]
ENTRYPOINT [ "sh","./entrypoint.sh" ]