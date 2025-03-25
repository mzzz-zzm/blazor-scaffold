# FROM grpcweb/prereqs AS grpcweb_builder
# WORKDIR /app
# COPY ./protobuf-gen/*.proto .
# RUN mkdir compiledproto
# RUN protoc --js_out=import_style=commonjs:./compiledproto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./compiledproto *.proto


# FROM golang:latest AS grpcgo_builder
# WORKDIR /app
# RUN mkdir compiledproto
# COPY ./protobuf-gen/*.proto .
# RUN apt-get update && apt-get install -y protobuf-compiler
# RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# RUN protoc -I=. --go_out=./compiledproto --go_opt=paths=source_relative --go-grpc_out=./compiledproto --go-grpc_opt=paths=source_relative *.proto

FROM golang:latest AS grpcgo_builder
WORKDIR /app
RUN mkdir -p compiled/pb
RUN mkdir pb
COPY ./protobuf-gen/pb/ ./pb/
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Run protoc against all .proto files in the pb directory and its subdirectories.
RUN protoc --go_out=./compiled --go_opt=paths=source_relative \
    --go-grpc_out=./compiled --go-grpc_opt=paths=source_relative ./pb/**/*.proto

#TODO:MH test
# ENTRYPOINT [ "tail", "-f", "/dev/null" ]

FROM alpine:latest
WORKDIR /app
COPY ./protobuf-gen/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
RUN mkdir compiled
COPY --from=grpcgo_builder /app/compiled ./compiled
RUN mkdir proto
COPY ./protobuf-gen/pb/ ./proto/

ENTRYPOINT [ "sh","./entrypoint.sh" ]