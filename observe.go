// Package main provides the observe application.
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

// context represents an observation context. It holds information
// for a single, specific observation.
type context struct {
	settings     *Settings // The parsed settings for an observation.
	interval     uint      // The lookup interval for checking an object.
	quitOnChange bool      // Indicates if observe should be quit after a change.
}

// observeWebsite runs an observation for a given URL. It tracks
// changes by comparing checksums and sends an e-mail accordingly
// to the provided settings.
func observeWebsite(ctx *context, url string, out io.Writer) error {
	var lastChecksum []byte

	for quit := false; !quit; {
		timer := time.NewTimer(time.Second * time.Duration(ctx.interval))
		<-timer.C

		checksum, err := getChecksum(url)
		if err != nil {
			return err
		}

		// If the last checksum has already been set and the new
		// checksum doesn't match the old one, the website changed.
		if lastChecksum != nil && bytes.Compare(checksum, lastChecksum) != 0 {
			err := sendNotificationMail(ctx, func() string {
				return fmt.Sprintf(`An observed website has changed: %s`, url)
			})
			if err != nil {
				return err
			}
			// Quit the observation if `--quit-on-change` has been set.
			quit = ctx.quitOnChange
		}

		lastChecksum = checksum
	}

	return nil
}

// getChecksum sends a GET request to the specified URL and calculates
// the checksum of the response body. Returns an error of the request
// or reading the body failed.
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

// sendNotificationMail sends an e-mail to the user indicating that an
// observed object has changed. The e-mail is sent via SendGrid, which
// has to be configured in the settings before.
func sendNotificationMail(ctx *context, mailBody func() string) error {
	from := &mail.Email{Address: ctx.settings.Mail.From}
	to := &mail.Email{Address: ctx.settings.Mail.To}

	subject := fmt.Sprintf(`observe: An observed object has changed`)
	body := mailBody()

	client := sendgrid.NewSendClient(ctx.settings.Sendgrid.Key)
	_, err := client.Send(mail.NewSingleEmail(from, subject, to, body, body))

	return err
}
