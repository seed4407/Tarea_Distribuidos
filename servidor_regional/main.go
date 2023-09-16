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
	"math/rand"
	"time"
	"strconv"
	"google.golang.org/grpc"
	// "github.com/streadway/amqp"
	pb "github.com/seed4407/Tarea_Distribuidos/proto"
)

var (
	port = flag.Int("port", 80, "The server port")
)

var datos_cupos int 
var err error
var datos_rechazados int 
var valor_inicial int
var valor_modificado int
var numeroAleatorio int
var limite_inferior int
var limite_superior int
var aux string
// var ch *amqp.Channel
// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedServidorRegionalServer
}

func (s *server) CuposDisponibles(ctx context.Context, in *pb.Cupo) (*pb.Recepcion, error) {
	datos_cupos, err = strconv.Atoi(in.GetCupos())
	if err != nil {
        log.Printf("Error %v\n", err)
    }
	valor_modificado = valor_inicial
	limite_inferior = (valor_modificado/2) - (valor_modificado/5)
	limite_superior = (valor_modificado/2) + (valor_modificado/5)
	if valor_modificado == 0 {
		numeroAleatorio = 0
	} else if valor_modificado == 1{
		numeroAleatorio = 1
	} else {
		numeroAleatorio = rand.Intn(limite_superior-limite_inferior+1) + limite_inferior
	}
	log.Printf("Hay %d personas interesadas en acceder a la beta",numeroAleatorio)
	aux = strconv.Itoa(numeroAleatorio)

	return &pb.Recepcion{Ok:aux}, nil
}

func (s *server) CuposRechazados(ctx context.Context, in *pb.Rechazado) (*pb.Recepcion, error) {
	datos_rechazados, err = strconv.Atoi(in.GetRechazados())
	if err != nil {
        log.Printf("Error %v\n", err)
    }
	valor_inicial = valor_inicial - (numeroAleatorio -  datos_rechazados)

	log.Printf("Se inscribieron %d personas",numeroAleatorio -  datos_rechazados)
	log.Printf("Quedan %d personas en espera de cupo",valor_inicial)

	return &pb.Recepcion{Ok:"ok"}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
    // Abrir el archivo en modo lectura
	filePath := "/app/parametros_de_inicio.txt"

    // Lee el contenido del archivo
    contenido, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Error al leer el archivo: %v\n", err)
        return
    }

	valor_inicial,err = strconv.Atoi(string(contenido))

	if valor_inicial >= 0{
		log.Printf("Inicio exitoso")
	}

	// conn, err := amqp.Dial("amqp://guest:guest@10.6.46.109:8082/")
    // if err != nil {
    //     log.Fatalf("No se pudo conectar a RabbitMQ: %v", err)
    // }
    // defer conn.Close()

    // ch, err := conn.Channel()
    // if err != nil {
    //     log.Fatalf("No se pudo abrir un canal: %v", err)
    // }
    // defer ch.Close()

    // // Debugger
    // fmt.Println("Conexión exitosa a RabbitMQ")
	
	// serverID := "server-1"

	//     // Mensaje a enviar
    // message := "Mensaje guardado en cola"

    // err = ch.Publish(
    //     "",     // Exchange
    //     "centralQueue", // Nombre de la cola
    //     false,  // Mandatory
    //     false,  // Immediate
    //     amqp.Publishing{
    //         ContentType: "text/plain",
    //         Body:        []byte(message),
    //         Headers:     amqp.Table{"server_id": serverID},
    //     })
    // if err != nil {
    //     log.Fatalf("No se pudo publicar el mensaje: %v", err)
    // }

    // // Debugger
    // fmt.Println("Mensaje publicado exitosamente en la cola 'central-queue'")
	
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
