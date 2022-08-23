package users

import (
	gomail "gopkg.in/gomail.v2"
)

type info struct {
	Name    string
	Message string
}

func (i info) sendMail(email string) error {

	// t := template.New("template.html")

	// var err error
	// t, err = t.ParseFiles("template.html")
	// if err != nil {
	// 	log.Println(err)
	// }

	// var tpl bytes.Buffer
	// if err := t.Execute(&tpl, i); err != nil {
	// 	log.Println(err)
	// }

	// result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "bedatuassefa@gmail.com")
	m.SetHeader("To", email)
	m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", "golang test")
	m.SetBody("text/html", "<h3>Name:</h3><span>{{.Name}}</span><br/><br/><h3>Email:</h3><span>{{.Message}}</span><br/>")
	// m.Attach("template.html") // attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, "mishassefa6@gmail.com", "mm")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
