package test

import (
	"log"
	"testing"

	"github.com/informeai/genqrcode/qr"
)

func TestEncodeImg(t *testing.T) {
	msg := "informeai"
	arqToSave := "qr"

	if err := qr.Encode(msg, arqToSave); err != nil {
		log.Fatalf("Not encode qrcode.\nerror: %v", err)
	}

}
