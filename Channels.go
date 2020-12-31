package godisc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const channelURL string = "https://discord.com/api/v8/channels/"



func GetChannel(channelID string, token string) (Channel, error) {
	var channel Channel
	
	resp, err := sendRequest("GET", fmt.Sprint(channelURL + channelID), nil, token, "application/json")

	if err != nil {
		return channel, err
	}


	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&channel)
	return channel, nil
}

func EditChannel(channelID string, newParams *ChannelEdit, token string) (Channel, error) {
	var channel Channel

	var bNewParams []byte

	bNewParams, err := json.Marshal(newParams)

	if err != nil { return channel, err }

	resp, err := sendRequest("PATCH", fmt.Sprint(channelURL + channelID), bytes.NewReader(bNewParams), token, "application/json")

	debugHttp(resp)
	
	if err != nil {
		return channel, err
	}

	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&channel)
	return channel, err
}

func DeleteChannel(channelID string, token string) error {
	_, err := sendRequest("DELETE", fmt.Sprint(channelURL + channelID), nil, token, "application/json")
	return err
}

func GetMessages(channelID string, messageID string, limit int, methodOfGettingMessages string, token string) ([]Message, error) {
	msgs := make([]Message, 0)

	var queryString string = strings.ToLower(methodOfGettingMessages)+ "=" + messageID

	if limit > 0 && limit <= 100{
		queryString += "&limit=" + strconv.Itoa(limit)
	}

	resp, err := sendRequest("GET", fmt.Sprint(channelURL + channelID + "/messages?" + queryString), nil, token, "application/json")

	if err != nil { return nil, err }

	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&msgs)
	return msgs, err
}

func GetMessage(channelID string, messageID string, token string) (Message, error){
	var message Message
	
	resp, err := sendRequest("GET", fmt.Sprint(channelURL + channelID + "/messages/" + messageID), nil, token, "application/json")
	
	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&message)
	return message, err
}

func SendMessage(channelID string, message MessageSend, token string, path string) (Message, error){
	var msg Message
	
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return msg, err
		}
		defer file.Close()
		filePart, err := writer.CreateFormFile("file", filepath.Base(path))
		if err != nil {
			return msg, err
		}
		_, err = io.Copy(filePart, file)
	}

	bMessage, err := json.Marshal(message)
	if err != nil { return msg, err}

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="payload_json"`)
	h.Set("Content-Type", "application/json")

	p, err := writer.CreatePart(h)

	_, err = p.Write(bMessage)
	if err != nil { return msg, err }

	err = writer.Close()
	if err != nil {
		return msg, err
	}
	
	resp, err := sendRequest("POST",  fmt.Sprint(channelURL + channelID + "/messages"), bytes.NewReader(body.Bytes()), token, writer.FormDataContentType())

	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&msg)
	return msg, err
}

func React(channelID string, messageID string, emoji string, token string) error{
	_, err:= sendRequest("PUT", fmt.Sprint(channelURL + channelID + "/messages/" + messageID + "/reactions/" + url.QueryEscape(emoji) + "/@me"), nil, token, "")
	return err
}

func RemoveReaction(channelID string, messageID string, emoji string, userID string, rmAllWithThisEmoji bool, token string) error {
	urlToSend := fmt.Sprint(channelURL + channelID + "/messages/" + messageID + "/reactions")
	
	if emoji != "" {
		urlToSend += "/" + url.QueryEscape(emoji)
		if userID != "" {
			urlToSend += "/" + userID
		} else if !rmAllWithThisEmoji {
			urlToSend += "/@me"
		}
	}
	_, err:= sendRequest("DELETE", urlToSend, nil, token, "")
	return err
}

func GetReactioners(channelID string, messageID string, emoji string, method string, userID string, limit int, token string) ([]User, error) {
	var usr []User
	
	safeEmoji := url.QueryEscape(emoji)

	urlToSend := fmt.Sprint(channelURL + channelID + "/messages/" + messageID + "/reactions/" + safeEmoji)

	usrExists := (method == "after" || method == "before") && userID != ""
	limitPresent := limit > 0 && limit <= 100

	if usrExists || limitPresent {
		urlToSend += "?"
		if usrExists {
			urlToSend += method + "=" + userID
			if limitPresent {
				urlToSend += "&limit=" + strconv.Itoa(limit)
			}
		} else if limitPresent {
			urlToSend += "limit=" + strconv.Itoa(limit)
		}
	}
	resp, err:= sendRequest("GET", urlToSend, nil, token, "")

	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&usr)
	return usr, err
}

func EditMessage(channelID string, messageID string, params MessageEdit, token string) (Message, error) {
	var msg Message

	bParams, err := json.Marshal(params)
	if err != nil { return msg, err }

	resp, err := sendRequest("PATCH", channelURL + channelID + "/messages/" + messageID, bytes.NewReader(bParams), token, "application/json")
	if err != nil { return msg, err }

	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&msg)
	return msg, err
}

func DeleteMessage(channelID string, messageID string, token string) error {
	_, err := sendRequest("DELETE", channelURL + channelID + "/messages/" + messageID, nil, token, "")
	return err
}

func BulkDeleteMessages(channelID string, messageIDs []string, token string) error {
	type messageStruct struct {
		Messages []string `json:"messages"`
	}

	msgStruct := messageStruct{
		Messages: messageIDs,
	}
	fmt.Println(msgStruct)
	bMessageIDs, err := json.Marshal(msgStruct)
	if err != nil { return err }
	_, err = sendRequest("POST", channelURL + channelID + "/messages/bulk-delete", bytes.NewReader(bMessageIDs), token, "application/json")
	return err
}

func EditChannelPermissions(channelID string, overwrite Overwrite, token string) error {
	url := channelURL + channelID + "/permissions/" + overwrite.ID
	type tmp struct {
		Allow string `json:"allow"`
		Deny string `json:"deny"`
		Type int `json:"type"`
	}

	typeForPerm, err := strconv.Atoi(overwrite.Type)
	if err != nil {return err}

	perm := tmp{
		Allow: strconv.Itoa(overwrite.Allow),
		Deny: strconv.Itoa(overwrite.Deny),
		Type: typeForPerm,
	}

	bPerm, err := json.Marshal(perm)
	if err != nil {return err}
	
	_, err = sendRequest("PUT", url, bytes.NewReader(bPerm), token, "application/json")
	return err
}

func DeleteChannelPermission(channelID string, overwriteID string, token string) error {
	_, err := sendRequest("DELETE", channelURL + channelID + "/permissions/" + overwriteID, nil, token, "")
	return err
}

func GetChannelInvites(channelID string, token string) ([]Invite, error) {
	var invObjects []Invite
	resp, err := sendRequest("GET", channelURL + channelID + "/invites", nil, token, "")
	if err != nil {return invObjects, err}
	
	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&invObjects)
	return invObjects, err
}

func CreateInvite(channelID string, params CreateInviteObject, token string) (Invite, error) {
	var inv Invite
	
	bParams, err := json.Marshal(params)
	if err != nil {return inv, err}

	resp, err := sendRequest("POST", channelURL + channelID + "/invites", bytes.NewReader(bParams), token, "application/json")

	debugHttp(resp)

	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&inv)
	return inv, err
}

func TypingIndicator(channelID string, token string) error {
	_, err := sendRequest("POST", channelURL + channelID + "/typing", nil, token, "")
	return err
}

func GetPinnedMessages(channelID string, token string) ([]Message, error) {
	var msgs []Message
	resp, err := sendRequest("GET", channelURL + channelID + "/pins", nil, token, "")
	if err != nil {return msgs, err}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&msgs)
	return msgs, err
}

func PinMessage(channelID string, messageID string, token string) error {
	_, err := sendRequest("PUT", channelURL + channelID + "/pins/" + messageID, nil, token, "")
	return err
}

func DeletePinnedMessage(channelID string, messageID string, token string) error {
	_, err := sendRequest("DELETE", channelURL + channelID + "/pins/" + messageID, nil, token, "")
	return err
}

func AddToGDM(dmID string, userID string, AccessToken string, nickname string, token string) error {
	params := map[string]string{
		"nick": nickname,
		"access_token": AccessToken,
	} 
	bParams, err := json.Marshal(params)
	if err != nil { return err }

	_, err = sendRequest("PUT", channelURL + dmID + "/recipients/" + userID, bytes.NewReader(bParams), token, "application/json")
	return err
}