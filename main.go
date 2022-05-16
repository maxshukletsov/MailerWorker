package main

import (
	. "Infrastructure.MailerWorker/mailer"
	mailer "Infrastructure.MailerWorker/mailsender"
	. "context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
}

func (s *server) SendMessage(ctx Context, input *MessageRequest) (*MessageReply, error) {
	mail := mailer.CreateMessage(input)
	err := mailer.SendMail(mail)
	if err != nil {
		log.Fatalln(err)
		return &MessageReply{Sent: "Error"}, err
	} else {
		return &MessageReply{Sent: "Send successful"}, nil
	}

}

func (s *server) SendDocumentMessage(ctx Context, input *DocumentMessageRequest) (*MessageReply, error) {
	mail, err := mailer.CreateMessageWithDocument(input)
	err = mailer.SendMail(mail)
	for _, s := range input.Documents {
		err = mailer.DeleteFileFromOs(s)
	}

	if err != nil {
		log.Fatalln(err)
		return &MessageReply{Sent: "Error"}, err
	} else {
		return &MessageReply{Sent: "Send successful"}, nil
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	log.Print("[Start server]")
	listener, err := net.Listen("tcp", ":80")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcserver := grpc.NewServer()
	RegisterMailerServer(grpcserver, &server{})
	reflection.Register(grpcserver)

	err = grpcserver.Serve(listener)
	if err != nil {
		log.Fatal("fail to server", err)
	}
}
