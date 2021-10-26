package main

import (
	"github.com/khan745/gokvdb/internal/pkg/command"
)

type apiResponse struct {
	Reply string
	Item  string
	Items []string
}

func getApiResponseByResult(res command.Response) apiResponse {
	apiRes := apiResponse{}

	switch t := res.(type) {
	case command.OkResult:
		apiRes.Reply = "OK"
	case command.StringReply:
		apiRes.Reply = "OKString"
		apiRes.Item = t.Value
	case command.ErrResult:
		apiRes.Reply = "Error"
	default:
		apiRes.Reply = "No api response"
	}
	return apiRes
}
