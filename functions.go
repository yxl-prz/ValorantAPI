package ValorantAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SERVER_TYPE string

var SERVER_TYPE_EU SERVER_TYPE = "eu"
var SERVER_TYPE_NA SERVER_TYPE = "na"

func GetUserMatches(server_type SERVER_TYPE, username, tag string) (*API_Response, error) {
	var res *API_Response
	resp, err := http.Get(fmt.Sprintf("https://api.henrikdev.xyz/valorant/v3/matches/%v/%v/%v", string(server_type), url.QueryEscape(username), url.QueryEscape(tag)))
	if err != nil {
		fmt.Println("Error Fetching:", err.Error())
		return res, err
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	return res, err
}
