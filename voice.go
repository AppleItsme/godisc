package godisc

import "encoding/json"

func GetVoiceRegions(token string) ([]VoiceRegion, error) {
	resp, err := sendRequest("GET", "https://discord.com/api/v8/voice/regions", nil, token, "")
	if err != nil {return []VoiceRegion{}, err}

	defer resp.Body.Close()

	vc := make([]VoiceRegion, 0)
	err = json.NewDecoder(resp.Body).Decode(&vc)
	return vc, err
}