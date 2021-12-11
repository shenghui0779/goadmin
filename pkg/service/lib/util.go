package lib

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/mojocn/base64Captcha"
)

var CaptchaDriver base64Captcha.Driver = base64Captcha.NewDriverString(
	39,
	120,
	0,
	base64Captcha.OptionShowHollowLine,
	4,
	base64Captcha.TxtNumbers+base64Captcha.TxtAlphabet,
	nil,
	nil,
	nil,
)

// GenSalt 生成随机加密盐
func GenSalt() string {
	nonce := make([]byte, 8)
	io.ReadFull(rand.Reader, nonce)

	return hex.EncodeToString(nonce)
}
