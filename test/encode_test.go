package test

import (
	"log"
	"testings"

	"github.com/informeai/genqrcode/qr"
)

func TestEncode(t *testings.T) {
	msg := "informeai"
	arqToSave := "qr"

	if err := qr.Encode(msg, arqToSave); err != nil {
		log.Fatalf("Not encode qrcode.\nerror: %v", err)
	}

}
