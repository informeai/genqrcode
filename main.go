package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
	decoder "github.com/tuotoo/qrcode"
)

//go:embed version
var version string

func main() {
	usage := fmt.Sprintf("Info:\nGenerate e decode qrcode of many entries.\nversion: %v\n\nUsage: genqrcode <command> <flags> <text|bytes>\n\nCommands:\n\tgenqrcode create <flags>	create qrcode with option flag.\n\tgenqrcode decode <flags>	decode qrcode of image.\n\nFlags:\n\t-t	Text for convert in qrcode image.\n\t-o   	name of arquivo for save.\n\t-i	image of decoder.", version)
	create := flag.NewFlagSet("create", flag.ExitOnError)
	text := create.String("t", "", "Text for convert in qrcode image.")
	arq := create.String("o", "", "name of arquivo for save.")
	decode := flag.NewFlagSet("decode", flag.ExitOnError)
	img := decode.String("i", "", "image of decoder.")
	flag.Parse()

	if len(os.Args) > 3 {

		switch os.Args[1] {
		case "create":
			create.Parse(os.Args[2:])
			convertText(*text, *arq)
		case "decode":
			decode.Parse(os.Args[2:])
			image, err := decodeImage(*img)
			if err != nil {
				log.Fatalf("Not decoded image.\nerror:%v", err)
			}
			fmt.Printf("Decode of %s is: \n%v", *img, image)
		default:
			fmt.Println(usage)
		}
	} else {
		fmt.Println(usage)
	}

}

func convertText(text string, name string) {
	if name != "" {

		if err := qrcode.WriteFile(text, 3, 256, fmt.Sprintf("%s.png", name)); err != nil {
			log.Fatalf("Not created qrcode image.\nerror:%v", err)
		}
	} else {
		if err := qrcode.WriteFile(text, 3, 256, fmt.Sprintf("%s.png", text)); err != nil {
			log.Fatalf("Not created qrcode image.\nerror:%v", err)
		}
	}
}
func decodeImage(img string) (string, error) {
	fi, err := os.Open(img)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer fi.Close()
	qrmatrix, err := decoder.Decode(fi)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return qrmatrix.Content, nil
}
