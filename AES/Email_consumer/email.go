package main

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/go-gomail/gomail"
)

func SendEm1(emailAddress string) {
	// Create a new message
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "teamfaketixito@gmail.com")
	mailer.SetHeader("To", emailAddress)
	mailer.SetHeader("Subject", "Welcome to our Ticketing Platform!")

	// Read the infographic image file
	infographicPath := "./emailwlc.jpg"
	_, infographicFileName := filepath.Split(infographicPath)
	infographicContent, err := ioutil.ReadFile(infographicPath)
	if err != nil {
		log.Fatal(err)
	}

	// Embed the infographic image
	mailer.Embed(infographicFileName, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(infographicContent)
		return err
	}))

	// Set the HTML body of the email with responsive design
	mailer.SetBody("text/html", `
		<html>
			<head>
				<style>
					@media screen and (max-width: 600px) {
						.container {
							width: 100% !important;
						}
					}
				</style>
			</head>
			<body style="margin: 0; padding: 0;">
				<table align="center" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
					<tr>
						<td align="center" bgcolor="#ffffff" style="padding: 40px 30px 20px 30px;">
							<h1 style="margin: 0; font-size: 30px; line-height: 36px;">Hi Cutie</h1>
							<h1 style="margin: 0; font-size: 30px; line-height: 36px;">Welcome to our Ticketing Platform!</h1>
							<p style="margin-top: 20px; font-size: 18px; line-height: 24px;">Thank you for joining us. We are excited to have you on board!</p>
						</td>
					</tr>
					<tr>
						<td align="center" bgcolor="#ffffff">
							<img src="cid:`+infographicFileName+`" alt="Infographic" style="display: block; max-width: 100%; height: auto;">
						</td>
					</tr>
				</table>
			</body>
		</html>
	`)

	//PDF attachment
	pdfPath := "./emailpdf.pdf"
	_, pdfFileName := filepath.Split(pdfPath)
	mailer.Attach(pdfPath, gomail.Rename(pdfFileName))

	//new mailer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "teamfaketixito@gmail.com", "pohkimrgoawdhacz")

	//TLS connection
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// end the email
	if err := dialer.DialAndSend(mailer); err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}

	log.Println("Email sent successfully!")
}
func SendEm2(emailAddress string) {
	// Create a new message
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "teamfaketixito@gmail.com")
	mailer.SetHeader("To", emailAddress)
	mailer.SetHeader("Subject", "Hello, World!")

	// Set the plain text body of the email
	mailer.SetBody("text/plain", "Hello, World!")

	// Create a new mailer and set the authentication details
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "teamfaketixito@gmail.com", "pohkimrgoawdhacz")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := dialer.DialAndSend(mailer); err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}

	log.Println("Email sent successfully!")
}

func FatherEmail(mstype string, ead string) {
	if mstype == "EM1" {
		SendEm1(ead)
	} else if mstype == "EM2" {
		SendEm2(ead)
	}
}
