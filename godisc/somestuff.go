package godisc

import (
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

func SendRequest(method string, URL string, body io.Reader, Headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, body)

	if err != nil {
		return nil, err
	}

	for key, value := range Headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	return resp, err
}