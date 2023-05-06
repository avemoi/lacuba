package main

import (
	"bytes"
	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
	"html/template"
	"sync"
	"time"
)

type Mail struct {
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
	ToAddress  string
	FromName   string
	Wait       *sync.WaitGroup
	MailerChan chan LacubaMessage
	ErrorChan  chan error
	DoneChan   chan bool
}

type LacubaMessage struct {
	FromName   string
	Subject    string
	Data       any
	DataMap    map[string]any
	LacubaId   int64
	LacubaLat  float64
	LacubaLng  float64
	LacubaAuth string
}

func (m *Mail) sendMail(lacubaMsg LacubaMessage, errorChan chan error) {
	defer m.Wait.Done()

	if lacubaMsg.FromName == "" {
		lacubaMsg.FromName = m.FromName
	}

	data := map[string]any{
		"message":    lacubaMsg.Data,
		"lacubaId":   lacubaMsg.LacubaId,
		"lacubaLat":  lacubaMsg.LacubaLat,
		"lacubaLng":  lacubaMsg.LacubaLng,
		"lacubaAuth": lacubaMsg.LacubaAuth,
	}

	lacubaMsg.DataMap = data

	// build html mail
	formattedMessage, err := m.buildHTMLMessage(lacubaMsg)

	if err != nil {
		errorChan <- err
	}

	// build plain text mail
	plainMessage, err := m.buildPlainTextMessage(lacubaMsg)
	if err != nil {
		errorChan <- err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		errorChan <- err
	}
	email := mail.NewMSG()
	email.SetFrom(m.Username).AddTo(m.ToAddress).SetSubject(lacubaMsg.Subject)
	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	err = email.Send(smtpClient)

	if err != nil {
		errorChan <- err
	}

}

func (m *Mail) buildPlainTextMessage(msg LacubaMessage) (string, error) {
	templateToRender := "./cmd/templates/plain-email.gohtml"
	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", nil
	}
	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", nil
	}
	plainMessage := tpl.String()

	return plainMessage, nil
}

func (m *Mail) buildHTMLMessage(lacubaMsg LacubaMessage) (string, error) {
	templateToRender := "./cmd/templates/html-email.gohtml"
	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", nil
	}
	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", lacubaMsg.DataMap); err != nil {
		return "", nil
	}
	formattedMessage := tpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", nil
	}
	return formattedMessage, nil
}

func (m *Mail) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		return "", nil
	}
	html, err := prem.Transform()
	if err != nil {
		return "", nil
	}
	return html, nil

}

func (app *Config) listenForMail() {
	for {
		select {
		case lacubaMsg := <-app.Mailer.MailerChan:
			go app.Mailer.sendMail(lacubaMsg, app.Mailer.ErrorChan)
		case err := <-app.Mailer.ErrorChan:
			// We can notify someone
			app.ErrorLog.Println(err)
		case <-app.Mailer.DoneChan:
			return
		}
	}
}

func (app *Config) createMail() Mail {
	// Create channels
	errorChan := make(chan error)
	mailerChan := make(chan LacubaMessage, 100)
	mailerDoneChan := make(chan bool)

	m := Mail{
		Domain:     "localhost",
		Host:       "smtp.gmail.com",
		Port:       587,
		Username:   "newlacuba@gmail.com",
		Password:   "",
		Encryption: "none",
		ToAddress:  "",
		FromName:   "LacubaInfo",
		Wait:       app.Wait,
		MailerChan: mailerChan,
		ErrorChan:  errorChan,
		DoneChan:   mailerDoneChan,
	}

	return m

}
