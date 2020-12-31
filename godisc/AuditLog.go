package godisc

import "encoding/json"

func GetAuditLog(guildID string, token string) (AuditLog, error) {
	var audit AuditLog

	resp, err := sendRequest("GET", "https://discord.com/api/v8/guilds/" + guildID + "/audit-logs", nil, token, "") 
	if err != nil { return audit, err }

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&audit)
	return audit, err
}