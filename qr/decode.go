package qr

import (
	"os"

	decoder "github.com/tuotoo/qrcode"
)

func Decode(img string) (string, error) {
	fi, err := os.Open(img)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	qrmatrix, err := decoder.Decode(fi)
	if err != nil {
		return "", err
	}
	return qrmatrix.Content, nil
}
