package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	qr "github.com/informeai/genqrcode/qr"
)

//go:embed doc/version
var version string

//go:embed doc/usage
var usage string

var description string = usage + "\n\nVersion: " + version

func main() {
	start()
}

func start() {
	fmt.Println(os.Args[0])
	if len(os.Args) > 3 && (os.Args[1] == "create" || os.Args[1] == "decode") {
		create := flag.NewFlagSet("create", flag.ExitOnError)
		text := create.String("t", "", "Text for convert in qrcode image.")
		arq := create.String("o", "", "name of arquivo for save.")
		decode := flag.NewFlagSet("decode", flag.ExitOnError)
		img := decode.String("i", "", "image of decoder.")
		flag.Parse()
		switch os.Args[1] {
		case "create":
			create.Parse(os.Args[2:])
			if err := qr.Encode(*text, *arq); err != nil {
				log.Fatalf("Not encode qrcode.\nerror: %v", err)
			}

		case "decode":
			decode.Parse(os.Args[2:])
			image, err := qr.Decode(*img)
			if err != nil {
				log.Fatalf("Not decoded image.\nerror:%v", err)
			}
			fmt.Printf("Decode of %s is: \n%v", *img, image)

		default:
			fmt.Println(description)
		}
	} else {
		fmt.Println(description)
	}
}
