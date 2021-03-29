package test

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/informeai/genqrcode/qr"
)

func TestDecodeImg(t *testing.T) {
	imageQr := filepath.Join("../images", "qr.png")
	text, err := qr.Decode(imageQr)
	if err != nil {
		log.Fatalf("Not decoded image.\nerror:%v", err)
	}
	fmt.Println(text)
}
