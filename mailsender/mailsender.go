package mailsender

import (
	"Infrastructure.MailerWorker/config"
	. "Infrastructure.MailerWorker/mailer"
	"bytes"
	"gopkg.in/gomail.v2"
	"io"
	"log"
	"os"
)

func SendMail(message *gomail.Message) error {
	mailerConfig := config.InitConfig()

	dialer := gomail.NewDialer(mailerConfig.SMTPHost, mailerConfig.SMTPPort, mailerConfig.SMTPUsername, mailerConfig.SMTPPassword)
	if err := dialer.DialAndSend(message); err != nil {
		return err
	} else {
		log.Println("Successfully sent message to:", message.GetHeader("To"))
		return nil
	}
}

func CreateMessage(input *MessageRequest) *gomail.Message {

	message := gomail.NewMessage()

	message.SetHeader("From", "noreply@example.com")
	message.SetHeader("To", input.SendTo...)
	message.SetHeader("Subject", input.Subject)
	message.SetBody("text/html", input.Body)

	return message
}

func CreateMessageWithDocument(input *DocumentMessageRequest) (*gomail.Message, error) {

	message := gomail.NewMessage()

	message.SetHeader("From", "noreply@example.com")
	message.SetHeader("To", input.SendTo...)
	message.SetHeader("Subject", input.Subject)
	message.SetBody("text/html", input.Body)

	for _, s := range input.Documents {
		err := AddFileToMessage(message, s)
		if err != nil {
			return nil, err
		}
	}
	return message, nil
}

func AddFileToMessage(message *gomail.Message, doc *Document) error {
	out, err := os.Create(doc.Filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(out)

	_, err = io.Copy(out, bytes.NewReader(doc.File))
	if err != nil {
		return err
	}
	message.Attach(doc.Filename)
	return nil
}

func DeleteFileFromOs(doc *Document) error {
	err := os.Remove("/" + doc.Filename)
	if err != nil {
		return err
	} else {
		return nil
	}
}
