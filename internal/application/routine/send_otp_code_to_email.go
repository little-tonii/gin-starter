package routine

import (
	"fmt"
	"gin-starter/internal/shared/constant"

	"gopkg.in/gomail.v2"
)

func SendOtpCodeToEmail(from string, to string, subject string, code string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", fmt.Sprintf("%s <%s>", from, constant.Environment.SMTP_USERNAME))
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody(
		"text/html",
		fmt.Sprintf(
			"<p>Đây là mã xác thực OTP của bạn %s. Mã có hiệu lực trong vòng 5 phút, vui lòng không cung cấp cho bất kỳ ai.</p>",
			code,
		),
	)
	dialer := gomail.NewDialer(
		constant.Environment.SMTP_SERVER,
		constant.Environment.SMTP_PORT,
		constant.Environment.SMTP_USERNAME,
		constant.Environment.SMTP_PASSWORD,
	)
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}
