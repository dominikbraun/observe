package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"io"
	"net/http"
	"time"
)

type context struct {
	settings     *Settings
	interval     uint
	quitOnChange bool
}

func observeWebsite(ctx *context, url string, out io.Writer) error {
	var lastChecksum []byte

	for quit := false; !quit; {
		timer := time.NewTimer(time.Second * time.Duration(ctx.interval))
		<-timer.C

		checksum, err := getChecksum(url)
		if err != nil {
			return err
		}

		if lastChecksum != nil && bytes.Compare(checksum, lastChecksum) != 0 {
			err := sendNotificationMail(ctx, func() string {
				return fmt.Sprintf("An observed website has changed: %s", url)
			})
			if err != nil {
				return err
			}
			quit = ctx.quitOnChange
		}

		lastChecksum = checksum
	}

	return nil
}

func getChecksum(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	hash := sha256.New()

	if _, err := io.Copy(hash, resp.Body); err != nil {
		return nil, err
	}
	_ = resp.Body.Close()

	return hash.Sum(nil), nil
}

func sendNotificationMail(ctx *context, mailBody func() string) error {
	from := &mail.Email{Address: ctx.settings.Mail.From}
	to := &mail.Email{Address: ctx.settings.Mail.To}

	subject := fmt.Sprintf("observe: An observed object has changed")
	body := mailBody()

	client := sendgrid.NewSendClient(ctx.settings.Sendgrid.Key)
	_, err := client.Send(mail.NewSingleEmail(from, subject, to, body, body))

	return err
}
