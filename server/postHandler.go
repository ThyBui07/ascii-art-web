package server

import (
	"ascii/drawing"
	"errors"
	"net/http"
	"strings"
	"unicode"
)

var FinalText string

type AsciiData struct {
	DrawingText string
	Submit      bool
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	FinalText = ""
	if r.Method != http.MethodPost {
		errorBadRequest(w, errors.New("only POST method allowed"))
		return
	}

	r.ParseForm()
	input := r.FormValue("text")
	banner := r.FormValue("fonts")
	// check the validate banner
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		errorBadRequest(w, errors.New("wrong banner type"))
		return
	}

	if banner == "" || input == "" {
		errorBadRequest(w, errors.New("body is wrong, should contain text and banner fields"))
		return
	}

	input = strings.ReplaceAll(input, "\r", "")
	// 32 - 126
	inputValidate := isASCII(input)
	if inputValidate != true {
		errorBadRequest(w, errors.New("only ASCII character allowed"))
		return
	}
	sourceTextArr := strings.Split(input, "\n")
	for _, v := range sourceTextArr {
		if len(v) != 0 {
			res := drawing.Display(v, banner)
			FinalText += res
		} else {
			FinalText += "\n"
		}

	}
	//take the last \n out
	FinalText = strings.TrimSuffix(FinalText, "\n")
	if FinalText == "" {
		errorInternalServer(w)
		return
	} else {
		data := AsciiData{
			DrawingText: FinalText,
			Submit:      true,
		}
		Tpl.Execute(w, data)
	}
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}
