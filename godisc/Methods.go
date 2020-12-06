package godisc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetChannel(ID string) (Channel, error) {
	var channel Channel
	resp, err := http.Get(fmt.Sprintf("https://discord.com/api/channels/%s", ID))
	if err != nil {
		return channel, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channel, err
	}
	err = json.Unmarshal(body, channel)
	if err != nil {
		return channel, err
	}
	return channel, nil
}

func ChangeChannel(ID string, newSettings *ChannelEdit)