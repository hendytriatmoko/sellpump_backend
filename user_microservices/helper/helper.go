package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"mime/multipart"
	"net/smtp"
	"os"
	"time"
	"user_microservices/common"
)

const charset = "abcdefghijklmnopqrstuvwxyz1234567890"
const length = 36

var screatkey = []byte("Si kepo hahaha")

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type Helper struct {
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

var auth smtp.Auth

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte("From: SellPump <no-reply@gmail.com>\r\n" + subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	//go smtp.SendMail(addr, auth, common.Config.EMAIL, r.to, msg);
	if err := smtp.SendMail(addr, auth, common.Config.EMAIL, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}


func (u *Helper) GetTimeNow() string {
	t := time.Now()
	return string(t.Format("2006-01-02 15:04:05.999999"))
}

func (u *Helper) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (u *Helper) StringWithCharset() string {
	b := make([]byte, length)
	for i := range b {
		if i < 8 {
			b[i] = charset[seededRand.Intn(len(charset))]
		} else if i > 8 && i < 13 {
			b[i] = charset[seededRand.Intn(len(charset))]
		} else if i > 13 && i < 18 {
			b[i] = charset[seededRand.Intn(len(charset))]
		} else if i > 18 && i < 23 {
			b[i] = charset[seededRand.Intn(len(charset))]
		} else if i > 23 {
			b[i] = charset[seededRand.Intn(len(charset))]
		} else {
			b[i] = '-'
		}
	}
	return string(b)
}

func (u *Helper) SendEmailVerifikasi(email string, id_user string, id_verifikasi string) error {

	templateData := struct {
		Email   		string
		IdUser			string
		IdVerifikasi 	string
	}{
		Email:     email,
		IdUser: id_user,
		IdVerifikasi: id_verifikasi,
	}
	r := NewRequest([]string{email}, "Verifikasi Email", "Hello, World!")
	//err := r.ParseTemplate("helper/email_invitation.html", templateData)
	if err := r.ParseTemplate("config/verifikasi.html", templateData); err == nil {
		ok, err := r.SendEmail()
		fmt.Println(ok)

		if err != nil {
			return err
		}
	}

	return error(nil)
}