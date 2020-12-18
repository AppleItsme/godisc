package godisc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var URL string = "https://discord.com/api/v8/channels/"



func GetChannel(channelID string, token string) (Channel, error) {
	var channel Channel
	
	resp, err := SendRequest("GET", fmt.Sprint(URL + channelID), nil, token, "application/json")

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

func EditChannel(channelID string, newParams *ChannelEdit, token string) (Channel, error) {
	var channel Channel

	var bNewParams []byte

	bNewParams, err := json.Marshal(newParams)

	if err != nil { return channel, err }

	resp, err := SendRequest("PATCH", fmt.Sprint(URL + channelID), bytes.NewReader(bNewParams), token, "application/json")

	DebugHttp(resp)
	
	if err != nil {
		return channel, err
	}

	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channel, err
	}
	err = json.Unmarshal(body, &channel)
	return channel, err
}

func DeleteChannel(channelID string, token string) error {
	_, err := SendRequest("DELETE", fmt.Sprint(URL + channelID), nil, token, "application/json")
	return err
}

func GetMessages(channelID string, messageID string, limit int, methodOfGettingMessages string, token string) ([]Message, error) {
	msgs := make([]Message, 0)

	var queryString string = strings.ToLower(methodOfGettingMessages)+ "=" + messageID

	if limit > 0 && limit <= 100{
		queryString += "&limit=" + strconv.Itoa(limit)
	}

	resp, err := SendRequest("GET", fmt.Sprint(URL + channelID + "/messages?" + queryString), nil, token, "application/json")

	if err != nil { return nil, err }

	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &msgs)
	return msgs, err
}

func GetMessage(channelID string, messageID string, token string) (Message, error){
	var message Message
	
	resp, err := SendRequest("GET", fmt.Sprint(URL + channelID + "/messages/" + messageID), nil, token, "application/json")
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return message, err
	}
	err = json.Unmarshal(body, &message)
	return message, err
}

func SendMessage(channelID string, message MessageSend, token string, path string) (Message, error){
	var msg Message

	file, err := os.Open(path)
	if err != nil {
		return msg, err
	}
	defer file.Close()
	
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	filePart, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return msg, err
	}
	_, err = io.Copy(filePart, file)


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
	fmt.Println(body)
	
	resp, err := SendRequest("POST",  fmt.Sprint(URL + channelID + "/messages"), bytes.NewReader(body.Bytes()), token, writer.FormDataContentType())

	defer resp.Body.Close()
	
	Respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msg, err
	}
	err = json.Unmarshal(Respbody, &msg)
	return msg, err
}