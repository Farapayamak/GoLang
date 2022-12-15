package farapayamak

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"bytes"
)

type RestClient struct {
	baseURL		string
	HTTPClient	*http.Client
	username	string
	password	string
}


type RestResponse struct {
	Value			string
	RetStatus		int
	StrRetStatus	string
}



func InitRestClient(username string, password string) *RestClient {
	return &RestClient {
		HTTPClient: &http.Client {
			Timeout: 1 * time.Minute,
		},
		baseURL: "https://rest.payamak-panel.com/api/SendSMS",
		username: username,
		password: password,
	}
}


// http.MethodPost
func (c *RestClient) CallRestAPI(req *http.Request, v *RestResponse) error {
	
	req.Header.Set("Accept", "application/json; charset=utf-8")
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into RestResponse
	if res.StatusCode != http.StatusOK {
		var errRes RestResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.StrRetStatus)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall and populate v
	fullResponse := RestResponse {
		Value: 			v.Value,
		RetStatus:		v.RetStatus,
		StrRetStatus:	v.StrRetStatus,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}



type SendSMSRestModel struct {
	to		string
	from	string
	text	string
}

func (c *RestClient) SendSMS(args *SendSMSRestModel) (*RestResponse, error) {

	jsonBody := []byte(fmt.Sprintf("{"username": "%s", "password": "%s"}", c.username, c.password))
 	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/SendSMS", c.baseURL), bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res := RestResponse{}
	if err := c.CallRestAPI(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}