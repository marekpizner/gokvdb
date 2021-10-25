package main

import (
	"github.com/khan745/gokvdb/internal/pkg/command"
)

type apiResponse struct {
	reply string
	item  string
	items []string
}

func getApiResponseByResult(res command.Response) apiResponse {
	apiRes := apiResponse{}

	switch t := res.(type) {
	case command.OkResult:
		apiRes.reply = "OK"
	case command.StringReply:
		apiRes.reply = "OKString"
		apiRes.item = t.Value
	case command.ErrResult:
		apiRes.reply = "Error"
	default:
		apiRes.reply = "No api response"
	}
	return apiRes
}
