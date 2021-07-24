# DSC Logo Generator 

[![GoDoc](https://godoc.org/github.com/baraa-almasri/dsc_logo_generator?status.png)](https://godoc.org/github.com/baraa-almasri/dsc_logo_generator) [![Go Report Card](https://goreportcard.com/badge/github.com/baraa-almasri/dsc_logo_generator)](https://goreportcard.com/report/github.com/baraa-almasri/dsc_logo_generator)

## dependencies:
- [fixed](https://godoc.org/golang.org/x/image/math/fixed)
- [go-ttf](https://godoc.org/github.com/golang/freetype/truetype)
- [ToggleButton](https://github.com/webomnizz/vue-toggle-button) component from [WebOmnizz](https://github.com/webomnizz/)

## run locally:
1. install the dependencies one by one 
   - or just run `go build -ldflags="-w -s"`
1. build the client
   - `cd ./client`
   - `npm run build`
1. run the server's executable
   - `./dsc_logo_generator`
1. open `127.0.0.1:1105` and generate some logos :)

## examples:
- ### vertical
<p align="center">
<img src="https://github.com/baraa-almasri/dsc_logo_generator/blob/main/res/example.png" >
</p>

- ### horizontal
<p align="center">
<img src="https://github.com/baraa-almasri/dsc_logo_generator/blob/main/res/example_horizontal.png" >
</p>