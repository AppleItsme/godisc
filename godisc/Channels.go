package godisc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var URL string = "https://discord.com/api/channels/"

func GetChannel(ChannelID string, Token string) (Channel, error) {
	var channel Channel

	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bot %s", Token)
	headers["Accept"] = "*/*"
	headers["Host"] = "discord.com"
	headers["Connection"] = "keep-alive"
	headers["Cookie"] = "__cfduid=d96c2d0da92ad05bf4370ef7a3065dedf1607773551"
	headers["User-Agent"] = fmt.Sprintf("Applandia/%s", Version)
	
	resp, err := SendRequest("GET", fmt.Sprint(URL + ChannelID), nil, headers)

	if err != nil {
		return channel, err
	}


	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channel, err
	}
	err = json.Unmarshal(body, &channel)
	if err != nil {
		return channel, err
	}
	return channel, nil
}

func EditChannel(channelID string, newParams ChannelEdit) error {
	bNewParams, err := json.Marshal(newParams)
	if err != nil { return err }
	SendRequest("PATCH", "http")
}

