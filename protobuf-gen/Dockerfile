FROM grpcweb/prereqs AS grpcweb_builder
WORKDIR /app
COPY ./protobuf-gen/*.proto .
RUN mkdir compiledproto
RUN protoc --js_out=import_style=commonjs:./compiledproto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./compiledproto *.proto


FROM golang:latest AS grpcgo_builder
WORKDIR /app
RUN mkdir compiledproto
COPY ./protobuf-gen/*.proto .
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN protoc -I=. --go_out=./compiledproto --go_opt=paths=source_relative --go-grpc_out=./compiledproto --go-grpc_opt=paths=source_relative *.proto


FROM alpine:latest
WORKDIR /app
COPY ./protobuf-gen/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
RUN mkdir artifacts

COPY --from=grpcweb_builder /app/compiledproto ./artifacts
COPY --from=grpcgo_builder /app/compiledproto ./artifacts

ENTRYPOINT [ "sh","./entrypoint.sh" ]