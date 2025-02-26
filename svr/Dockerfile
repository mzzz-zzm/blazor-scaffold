FROM golang:latest AS grpcgo_builder
WORKDIR /app
RUN mkdir compiledproto
COPY ./protobuf-gen/*.proto .
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN protoc -I=. --go_out=./compiledproto --go_opt=paths=source_relative --go-grpc_out=./compiledproto --go-grpc_opt=paths=source_relative *.proto

FROM golang:latest
WORKDIR /app
COPY ./svr/go.mod .
# COPY ./svr/go.sum .
RUN go mod tidy
COPY ./svr/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
RUN mkdir artifacts
COPY --from=grpcgo_builder /app/compiledproto ./artifacts

#copy compiled proto to mounted volume
ENTRYPOINT [ "sh","./entrypoint.sh" ]