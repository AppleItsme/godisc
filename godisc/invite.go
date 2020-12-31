package godisc

import (
	"encoding/json"
	"strconv"
)

const inviteURL string = "https://discord.com/api/v8/invites/"

func GetInvite(code string, containApproxMemberCounts bool, token string) (Invite, error) {
	resp, err := sendRequest("GET", inviteURL + code + "?with_counts=" + strconv.FormatBool(containApproxMemberCounts), nil, token, "")
	if err != nil {return Invite{}, err}

	defer resp.Body.Close()

	var inv Invite
	err = json.NewDecoder(resp.Body).Decode(&inv)
	return inv, err
}

func DeleteInvite(code string, token string) (Invite, error) {
	resp, err := sendRequest("DELETE", inviteURL + code, nil, token, "")
	if err != nil {return Invite{}, err}

	defer resp.Body.Close()

	var inv Invite
	err = json.NewDecoder(resp.Body).Decode(&inv)
	return inv, err
}

