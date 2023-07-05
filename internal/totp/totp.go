package totp

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func GetCode(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now().UTC())
}
