package api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/GDSC-ASU/logo_generator/logogen"
)

// setupResponse sets required response headers.
func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// GetLogo generates a dsc logo based on the given request body it uses
// GetLogoWithTextWithPadding from the LogoGenerator package to append university
// name with padding. it works on the very basic 4 steps: 1. get logo properties
// from the request body, 2. get a proper raw-logo and font color based on image
// color type 3. pass data to LogoGenerator, and generate a logo :) 4. send the
// generated b64 image to the response. no error handling is available yet!,
// since this api is ONLY called from the provided front end.
func GetLogo(res http.ResponseWriter, req *http.Request) {
	setupResponse(&res)

	propsJuicer := NewLogoPropsJuicer(req)
	generator := logogen.NewLogoGenerator(
		propsJuicer.GetRawLogo(),
		propsJuicer.GetText(),
		propsJuicer.GetImageBackground())

	newLogoBytes := generator.GenerateLogoWithPadding(
		propsJuicer.GetPadding())

	responseJSON := make(map[string]string)
	responseJSON["image"] = base64.StdEncoding.EncodeToString(newLogoBytes)
	_ = json.NewEncoder(res).Encode(responseJSON)
}
