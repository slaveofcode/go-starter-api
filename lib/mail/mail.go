package mail

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// Charset for email
	Charset = "UTF-8"
)

var sess, _ = session.NewSession(&aws.Config{
	Region: aws.String("us-east-1"),
})

var sesMail = ses.New(sess)

// Template is a wrapper template for sending email
type Template struct {
	From       string
	Recipients []*string
	CC         []*string
	BCC        []*string
	Subject    string
	HTML       string
	Text       string
}

// Send used for sending a single email
func Send(tpl *Template) (*ses.SendEmailOutput, error) {
	mailData := compile(tpl)
	res, err := sesMail.SendEmail(mailData)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return nil, err
	}

	return res, nil
}

func compile(tpl *Template) *ses.SendEmailInput {
	var email ses.SendEmailInput
	email.SetDestination(compileDestination(tpl.Recipients, tpl.CC, tpl.BCC))
	email.SetMessage(compileMessage(tpl.Subject, tpl.HTML, tpl.Text))
	email.SetSource(*aws.String(tpl.From))
	return &email
}

func compileDestination(recipients, ccs, bccs []*string) *ses.Destination {
	var dest ses.Destination
	if len(recipients) > 0 {
		dest.SetToAddresses(recipients)
	}

	if len(ccs) > 0 {
		dest.SetCcAddresses(ccs)
	}

	if len(bccs) > 0 {
		dest.SetBccAddresses(bccs)
	}

	return &dest
}

func compileMessage(subject, msgHTML, msgText string) *ses.Message {
	var msg ses.Message

	if subject == "" {
		msg.SetSubject(&ses.Content{
			Charset: aws.String(Charset),
			Data:    aws.String("[No Subject]"),
		})
	} else {
		msg.SetSubject(&ses.Content{
			Charset: aws.String(Charset),
			Data:    aws.String(subject),
		})
	}

	var body ses.Body

	if msgHTML != "" {
		body.SetHtml(&ses.Content{
			Charset: aws.String(Charset),
			Data:    aws.String(msgHTML),
		})
	}

	if msgText != "" {
		body.SetText(&ses.Content{
			Charset: aws.String(Charset),
			Data:    aws.String(msgText),
		})
	}

	msg.SetBody(&body)

	return &msg
}
