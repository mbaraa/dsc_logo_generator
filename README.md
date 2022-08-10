# GDSC Logo Generator 

[![GoDoc](https://godoc.org/github.com/GDSC-ASU/logo_generator?status.png)](https://godoc.org/github.com/GDSC-ASU/logo_generator) [![Go Report Card](https://goreportcard.com/badge/github.com/GDSC-ASU/logo_generator)](https://goreportcard.com/report/github.com/GDSC-ASU/logo_generator)

## dependencies:
- [go-cairo](https://github.com/ungerik/go-cairo) cairo the image manipulation library
- [go-ttf](https://godoc.org/github.com/golang/freetype/truetype) used to calculate the text's length using a specific ttf font
- [fixed](https://godoc.org/golang.org/x/image/math/fixed) the fixed numerical type is used for the ttf glyph's length
- [ToggleButton](https://github.com/webomnizz/vue-toggle-button) component from [WebOmnizz](https://github.com/webomnizz/) used in the dark mode switcher

## run locally:
1. install the dependencies one by one 
   - or just run `go mod tidy`
1. build the client(front-end)
   - `cd ./client`
   - `npm run build`
1. build run the server's executable
   - `go build -ldflags="-w -s"`
   - `./dsc_logo_generator`
1. open `127.0.0.1:1105` and generate some logos :)

## examples:
- ### horizontal
<p align="center">
<img src="https://github.com/GDSC-ASU/logo_generator/blob/main/res/example_horizontal.png" >
</p>

- ### vertical
<p align="center">
<img src="https://github.com/GDSC-ASU/logo_generator/blob/main/res/example.png" >
</p>
