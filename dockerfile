FROM golang:latest

WORKDIR /app

RUN rm  ~/.docker/config.json 
COPY go.mod .
# COPY server_central.go .
COPY central ./central
# COPY server_regional.go .
# COPY parametros_de_inicio.txt .
COPY region ./region
#COPY estructura.proto .
COPY proto ./proto

RUN apt-get update
RUN export PATH=$PATH:/usr/local/go/bin
RUN apt-get install -y protobuf-compiler
RUN go get google.golang.org/grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN protoc --go_out=./proto --go_opt=paths=import \ 
--go-grpc_out=./proto --go-grpc_opt=paths=import \
 ./proto/*.proto

WORKDIR /app/central
RUN go build -o bin .

WORKDIR /app/region
RUN go build -o bin .

WORKDIR /app
RUN go get github.com/rabbitmq/amqp091-go

# RUN protoc --go_out=./proto ./proto/*.proto

# RUN go run server_regional.go
# RUN go run server_central.go

ENTRYPOINT [ "/app/bin" ]

# option go_package = "/app/proto";

# package grpc;

# service sincrono {
#   rpc comunicacion_sincrona(stream MensajeEntrada) returns (stream MensajeSalida) {}
# }

# message MensajeEntrada {
#   string campo1 = 1;
#   int32 campo2 = 2;
# }

# message MensajeSalida {
#   string respuesta = 1;
# }