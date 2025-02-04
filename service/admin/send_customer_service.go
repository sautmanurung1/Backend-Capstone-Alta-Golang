package admin

import (
	"fmt"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/lib"
	"net/smtp"
)

type svcSendCustomer struct {
	c    database.Config
	repo domains.SendCustomerRepository
}

func NewSendCustomerService(repo domains.SendCustomerRepository, c database.Config) domains.SendCustomerService {
	return &svcSendCustomer{
		c:    c,
		repo: repo,
	}
}

func (s *svcSendCustomer) SendEmailService(message entities.SendCustomer) error {
	server := "smtp-mail.outlook.com"
	port := 587
	user := "rolandbrilianto@outlook.com"
	from := user
	pass := "Kipasangin_1"
	dest := message.To

	auth := lib.LoginAuth(user, pass)

	to := []string{dest}

	messages := []byte("From: " + from + "\n" +
		"To: " + dest + "\n" +
		"Subject: " + message.Subject + "\n\n" +
		message.Body)

	endpoint := fmt.Sprintf("%v:%v", server, port)

	err := smtp.SendMail(endpoint, auth, from, to, messages)
	if err != nil {
		fmt.Println("this is error guys :", err)
	}
	return s.repo.SendEmail(message)
}
