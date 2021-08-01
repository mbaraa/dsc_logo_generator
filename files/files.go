package files

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

// BasicLogoProps defines basic logo props
//
type BasicLogoProps struct {
	Logo     string  `json:"logo"`
	TextSize float64 `json:"text_size"`
}

// DecodedB64File defines decoded base64 file
//
type DecodedB64File struct {
	Data     []byte
	TextSize float64 // just for the logo
}

// B64Files represents the file used in the program represented as base64 strings
// well this way is faster than loading an actual file (i've tested alot)
// and swapping files only requires restarting the server :0
//
type B64Files struct {
	Logos struct {
		VColor BasicLogoProps `json:"v-color"`
		VWhite BasicLogoProps `json:"v-white"`
		HColor BasicLogoProps `json:"h-color"`
		HWhite BasicLogoProps `json:"h-white"`
	} `json:"logos"`

	Fonts []string `json:"fonts"` // fonts[0] = ProductSans
}

// Files manages the base64 files pulled from the json file (file of files LOL)
//
type Files struct {
	b64Files *B64Files
	files    struct {
		VColor      DecodedB64File
		HColor      DecodedB64File
		VWhite      DecodedB64File
		HWhite      DecodedB64File
		ProductSans DecodedB64File
	}
}

var filesInstance *Files = nil

// GetFilesInstance returns a singlton Files instance (much faster)
//
func GetFilesInstance() *Files {
	if filesInstance == nil {
		filesInstance = new(Files).
			initB64Files().
			initFilesBytes()
	}
	return filesInstance
}

func (f *Files) initB64Files() *Files {
	b64json, err := os.Open("b64_files.json")
	if err != nil {
		panic("na ah ah, you can't run this w/o the files file, err: " + err.Error())
	}
	b64Files := new(B64Files)
	json.NewDecoder(b64json).Decode(&b64Files)
	f.b64Files = b64Files

	return f
}

func (f *Files) initFilesBytes() *Files {
	f.files.VColor.Data = DecodeB64IntoByteSlice(f.b64Files.Logos.VColor.Logo)
	f.files.VColor.TextSize = f.b64Files.Logos.VColor.TextSize
	
	f.files.HColor.Data = DecodeB64IntoByteSlice(f.b64Files.Logos.HColor.Logo)
	f.files.HColor.TextSize = f.b64Files.Logos.HColor.TextSize

	f.files.VWhite.Data = DecodeB64IntoByteSlice(f.b64Files.Logos.VWhite.Logo)
	f.files.VWhite.TextSize = f.b64Files.Logos.VWhite.TextSize
	
	f.files.HWhite.Data = DecodeB64IntoByteSlice(f.b64Files.Logos.HWhite.Logo)
	f.files.HWhite.TextSize = f.b64Files.Logos.HWhite.TextSize
	
	f.files.ProductSans.Data = DecodeB64IntoByteSlice(f.b64Files.Fonts[0]) // fonts[0] = ProductSans

	return f
}

// GetVerticalLogoColored returns basic logo props
//
func (f *Files) GetVerticalLogoColored() DecodedB64File {
	return f.files.VColor
}

// GetVerticalLogoColored returns basic logo props
//
func (f *Files) GetVerticalLogoWhite() DecodedB64File {
	return f.files.VWhite
}

// GetHorizontalLogoColored returns basic logo props
//
func (f *Files) GetHorizontalLogoColored() DecodedB64File {
	return f.files.HColor
}

// GetHorizontalLogoWhite returns basic logo props
//
func (f *Files) GetHorizontalLogoWhite() DecodedB64File {
	return f.files.HWhite
}

// GetProductSansFont returns a byte slice of the thing
//
func (f *Files) GetProductSansFont() []byte {
	return f.files.ProductSans.Data // the first font is Product Sans
}

// EncodeB64ByteSlice decodes the given b64 byte slice into a string
//
func EncodeB64ByteSlice(b64 []byte) string {
	return base64.StdEncoding.EncodeToString(b64)
}

// DecodeB64IntoByteSlice decodes the given b64 string into a byte slice
//
func DecodeB64IntoByteSlice(b64 string) []byte {
	bb, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil
	}
	return bb
}
