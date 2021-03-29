v = 0.1.0
gen:
	rm -rf bin dist
	mkdir bin dist
	env GOOS=darwin GOARCH=amd64 go build -o bin/macosx/genqrcode main.go && zip -r dist/genqrcode-darwin-$(v).zip bin/macosx/genqrcode
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/genqrcode main.go && zip -r dist/genqrcode-linux-$(v).zip bin/linux/genqrcode
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/genqrcode.exe main.go && zip -r dist/genqrcode-windows-$(v).zip bin/windows/genqrcode.exe