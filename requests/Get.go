package requests

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func Get(url string) *resty.Response {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeader("User-Agent", "BoofRip/v1.0.0").
		SetHeader("x-api-version", "10").
		Get(url)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}
