package main

func (app *Config) sendEmail(lacubaMsg LacubaMessage) {
	app.Wait.Add(1)
	app.Mailer.MailerChan <- lacubaMsg
}
