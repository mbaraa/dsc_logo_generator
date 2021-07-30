package imagefile

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func ConvertReaderToByteSlice(imageFile *os.File) ([]byte, error) {
	fileData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return nil, err
	}

	// happily ever after
	return []byte(EncodeB64ByteSlice(fileData)), nil
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