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
Download the [binary](https://github.com/informeai/genqrcode/releases/tag/v0.1.0) for platform in personal use.
Unzip directory received.
#### Linux or Macosx
converts text to qrcode image.
```
genqrcode create -t <text>
```
decodes qrcode image into text.
```
genqrcode decode -i <image>
```
#### Windows
converts text to qrcode image.
```
genqrcode.exe create -t <text>
```
decodes qrcode image into text.
```
genqrcode.exe decode -i <image>
```
### Alternative with makefile
Verify if make is instaled.

Execute:
```
make gen
```
will generate all binaries and zipped files.