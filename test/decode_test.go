package test

import (
	"fmt"
	"log"
	"testings"

	"github.com/informeai/genqrcode/qr"
)

func TestDecode(t *testings.T) {
	imageQr := "qr.png"
	text, err := qr.Decode(imageQr)
	if err != nil {
		log.Fatalf("Not decoded image.\nerror:%v", err)
	}
	fmt.Println(text)
}
