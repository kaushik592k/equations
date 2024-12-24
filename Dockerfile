FROM golang:1.23 AS builder
# install protoc compiler with go plugin
RUN apt-get update && apt-get install -y \ 
    protobuf-compiler \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
COPY protobuf /protobuf
RUN protoc -I=/protobuf --go_out=. --go-grpc_out=. /protobuf/equation.proto
RUN ls -l /app

RUN go build -o equation . 

# run the image
FROM alpine:latest 
WORKDIR /root/
COPY --from=builder /app/equation .
EXPOSE 50051
CMD ["./equation"]
