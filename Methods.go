package godisc

import (
	"net/http"
)
	
func GetChannel(ID string) {
	URL := "https://discord.com/api/channels/" + ID
	resp, err := http.Get(URL)
	if err != nil {
		return 
	}
}