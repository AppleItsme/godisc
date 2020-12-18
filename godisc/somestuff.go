package godisc

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var Version string = "0.0.1"

func DebugHttp(resp *http.Response) {
	log.Printf("API RESPONSE  STATUS :: %s\n", resp.Status)
	for k, v := range resp.Header {
		log.Printf("API RESPONSE  HEADER :: [%s] = %+v\n", k, v)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("API RESPONSE    BODY :: *Couldn't fetch*", resp)
	} else {
		log.Printf("API RESPONSE    BODY :: [%s]\n\n\n", body)
	}
}


func SendRequest(method string, URL string, body io.Reader, token string, contentType string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, URL, body)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", token))
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "discord.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "__cfduid=d96c2d0da92ad05bf4370ef7a3065dedf1607773551")
	req.Header.Add("User-Agent", fmt.Sprintf("Applandia/%s", Version))
	req.Header.Add("Content-Type", contentType)

	resp, err := client.Do(req)
	return resp, err
}
