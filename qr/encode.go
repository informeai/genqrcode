package qr

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func Encode(text string, name string) error {
	if name != "" {

		if err := qrcode.WriteFile(text, 3, 256, fmt.Sprintf("%s.png", name)); err != nil {
			return err
		}
	} else {
		if err := qrcode.WriteFile(text, 3, 256, fmt.Sprintf("%s.png", text)); err != nil {
			return err
		}
	}
	return nil
}
