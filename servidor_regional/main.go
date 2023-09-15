/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net" 
    "os"

	"google.golang.org/grpc"
	pb "github.com/seed4407/Tarea_Distribuidos/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedServidorRegionalServer
}

func (s *server) CuposDisponibles(ctx context.Context, in *pb.Cupo) (*pb.Recepcion, error) {
	log.Printf(in.GetCupos())
	return &pb.Recepcion{Ok:"ok "}, nil
}

func (s *server) CuposRechazados(ctx context.Context, in *pb.Rechazado) (*pb.Recepcion, error) {
	log.Printf(in.GetRechazados())
	return &pb.Recepcion{Ok:"ok"}, nil
}

func main() {
    // Abrir el archivo en modo lectura
    archivo, err := os.Open("./servidor_regional/parametros_de_inicio.txt")
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    defer archivo.Close()

    // Leer el contenido del archivo
    buffer := make([]byte, 1024)
    for {
        n, err := archivo.Read(buffer)
        if err != nil {
            break
        }
        fmt.Print(string(buffer[:n]))
    }


	lis, err := net.Listen("tcp", fmt.Sprintf(":%d",*port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServidorRegionalServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
