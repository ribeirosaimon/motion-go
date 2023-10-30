package emailSender

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/config/util"
	"gopkg.in/gomail.v2"
)

func SendEmail(code string) error {
	propertiesFile := "config.properties"
	dir, _ := util.FindRootDir()
	file := properties.MustLoadFile(fmt.Sprintf("%s/%s", dir, propertiesFile), properties.UTF8)
	user := file.GetString("email.sender.user", "")
	pass := file.GetString("email.sender.password", "")
	host := file.GetString("email.sender.host", "")
	port := 587
	msg := gomail.NewMessage()

	msg.SetHeader("From", "saimonribeiros@hotmail.com")
	msg.SetHeader("To", "saimonribeiros@hotmail.com")
	msg.SetHeader("Subject", "Your code")
	msg.SetBody("text/html", emailValue(code))

	dialer := gomail.NewDialer(host, port, user, pass)
	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func GenerateEmailCode() string {
	rand.Seed(time.Now().UnixNano())
	codigo := rand.Intn(900000) + 100000
	return fmt.Sprintf("%06d", codigo)
}

// https://mailtrap.io/
