package godisc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

const emojiURL string = "https://discord.com/api/v8/guilds/"

func GetGuildEmojis(guildID string, token string) ([]Emoji, error) {
	resp, err := sendRequest("GET", emojiURL + guildID + "/emojis", nil, token, "")
	if err != nil {return []Emoji{}, err}

	var emojis []Emoji

	err = json.NewDecoder(resp.Body).Decode(&emojis)

	return emojis, err
}

func GetGuildEmoji(guildID string, emojiID string, token string) (Emoji, error) {
	resp, err := sendRequest("GET", emojiURL + guildID + "/emojis/" + emojiID, nil, token, "")
	if err != nil {return Emoji{}, err}

	var emoji Emoji

	err = json.NewDecoder(resp.Body).Decode(&emoji)

	return emoji, err
}

func CreateGuildEmoji(guildID string, emojiName string, roles []Role, imageType string, image image.Image, token string) (Emoji, error) {
	var imgBuff bytes.Buffer

	switch imageType{
	case "png":
		png.Encode(&imgBuff, image)
	case "jpeg":
		jpeg.Encode(&imgBuff, image, nil)
	case "gif":
		gif.Encode(&imgBuff, image, nil)
	default:
		return Emoji{}, fmt.Errorf("expected jpeg/gif/png, got &v", imageType)
	}
	imageString := base64.StdEncoding.EncodeToString(imgBuff.Bytes())

	jsonParams := map[string]interface{}{
		"name": emojiName, 
		"image": "data:image/"+imageType+";base64,"+imageString,
		"roles": roles,
	}
	bjson, err := json.Marshal(jsonParams)
	if err != nil {
		return Emoji{}, err
	}

	resp, err := sendRequest("POST", emojiURL + guildID + "/emojis", bytes.NewReader(bjson),token, "application/json")
	if err != nil {return Emoji{}, err}
	defer resp.Body.Close()
	var emoji Emoji
	err = json.NewDecoder(resp.Body).Decode(&emoji)

	return emoji, err
}

func EditEmoji(guildID string, emojiID string, name string, roles []Role, token string) (Emoji, error) {
	jsonParams := map[string]interface{}{
		"name": name,
		"roles": roles,
	}
	bParams, err := json.Marshal(jsonParams)
	if err != nil { return Emoji{}, err }

	resp, err := sendRequest("PATCH", emojiURL + guildID + "/emojis/" + emojiID, bytes.NewReader(bParams), token, "application/json")
	if err != nil { return Emoji{}, err }

	defer resp.Body.Close()

	var emoji Emoji
	err = json.NewDecoder(resp.Body).Decode(&emoji)
	return emoji, err
}

func DeleteEmoji(guildID string, emojiID string, token string) error{
	_, err := sendRequest("DELETE", emojiURL + guildID + "/emojis/" + emojiID, nil, token, "application/json")
	return err
}