package files

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

// B64Files represents the file used in the program represented as base64 strings
// well this way is faster than loading an actual file (i've tested alot)
// and swapping files only requires restarting the server :0
//
type B64Files struct {
	Logos struct {
		VColor string `json:"v-color"`
		VGray  string `json:"v-gray"`
		VWhite string `json:"v-white"`
		HColor string `json:"h-color"`
		HGray  string `json:"h-gray"`
		HWhite string `json:"h-white"`
	} `json:"logos"`

	Fonts []string `json:"fonts"` // fonts[0] = ProductSans
}

// Files manages the base64 files pulled from the json file (file of files LOL)
//
type Files struct {
	b64Files *B64Files
}

var filesInstance *Files = nil

// GetFilesInstance returns a singlton Files instance (much faster)
//
func GetFilesInstance() *Files {
	if filesInstance == nil {
		b64json, err := os.Open("b64_files.json")
		if err != nil {
			panic("na ah ah, you can't run this w/o the files file, err: " + err.Error())
		}

		b64Files := new(B64Files)
		json.NewDecoder(b64json).Decode(&b64Files)
		filesInstance = &Files{b64Files}
	}
	return filesInstance
}

func (f *Files) GetVerticalLogoColored() []byte {
	return DecodeB64IntoByteSlice(f.b64Files.Logos.VColor)
}

func (f *Files) GetVerticalLogoGray() []byte {
	return DecodeB64IntoByteSlice(f.b64Files.Logos.VGray)
}

func (f *Files) GetVerticalLogoWhite() []byte {
	return DecodeB64IntoByteSlice(f.b64Files.Logos.VWhite)
}

func (f *Files) GetHorizontalLogoColored() []byte {
	return DecodeB64IntoByteSlice(f.b64Files.Logos.HColor)
}

func (f *Files) GetHorizontalLogoGray() []byte {
	return DecodeB64IntoByteSlice(f.b64Files.Logos.HGray)
}

func (f *Files) GetHorizontalLogoWhite() []byte {
	return DecodeB64IntoByteSlice(f.b64Files.Logos.HWhite)
}

func (f *Files) GetProductSansFont() []byte {
	return DecodeB64IntoByteSlice(f.b64Files.Fonts[0]) // the first font is Product Sans
}

func EncodeB64ByteSlice(b64 []byte) string {
	return base64.StdEncoding.EncodeToString(b64)
}

func DecodeB64IntoByteSlice(b64 string) []byte {
	bb, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil
	}
	return bb
}
