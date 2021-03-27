# genqrcode
CLI of encode or decode qrcode images.
### Get
```
go get -u github.com/informeai/genqrcode
```
### Build
```
go build main.go
```
### Usage
```
./main create -t <text>
```
converts text to qrcode image.
```
./main decode -i <image>
```
decodes qrcode image into text.
