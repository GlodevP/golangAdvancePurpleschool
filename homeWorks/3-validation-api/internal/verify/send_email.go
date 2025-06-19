package verify

import (
	"fmt"
	"net/smtp"

	"github.com/google/uuid"
	e "github.com/jordan-wright/email"
)

func (handler *VerifyHandler) sendEmailVerify(emailAddr string) error {
	hash := uuid.NewString()
	mail := e.NewEmail()
	mail.From = "no-replay <" + handler.dependens.cfg.EmailConfig.EmailSendler + ">"
	mail.To = []string{emailAddr}
	mail.Subject = "Verify you email address"
	mail.HTML = []byte(fmt.Sprintf("<p>Click this <a href=\"%s/verify/%s\">link</a> to verify tou email.</p>", handler.dependens.cfg.Url, hash))
	err := mail.Send(handler.dependens.cfg.EmailConfig.EmailServer+":"+handler.dependens.cfg.EmailConfig.EmailServerPort, smtp.PlainAuth("", handler.dependens.cfg.EmailConfig.EmailSendler, handler.dependens.cfg.EmailConfig.EmailSecret, handler.dependens.cfg.EmailConfig.EmailServer))
	if err != nil {
		fmt.Println(err)
		return err
	}
	handler.dependens.DB.AddHash(emailAddr, hash)
	return nil
}
