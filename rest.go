package farapayamak

import (
	"encoding/json"
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
	debug		bool
}



func InitRestClient(username string, password string) *RestClient {
	return &RestClient {
		HTTPClient: &http.Client {
			Timeout: 1 * time.Minute,
		},
		baseURL: "https://rest.payamak-panel.com/api/SendSMS",
		username: username,
		password: password,
		debug: true,
	}
}


// http.MethodPost
func (c *RestClient) callRestAPI(req *http.Request, v *RestResponse) error {
	
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall and populate v
	fullResponse := RestResponse {
		// Value: 			v.Value,
		// RetStatus:		v.RetStatus,
		// StrRetStatus:	v.StrRetStatus,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	if c.debug {
		// fmt.Printf("%+v\n", *v)
		fmt.Printf("%+v\n", fullResponse)
	}

	return nil
}


func (c *RestClient) addCredentials(data interface{}) (*string, error) {
   
	var mapped map[string]interface{}
    jsonStr, err := json.Marshal(data)
	if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(jsonStr, &mapped); err != nil {
		return nil, err
	}

	mapped["username"] = c.username
	mapped["password"] = c.password

	if c.debug {
		// iterate through
		for field, val := range mapped {
			fmt.Println("KV Pair: ", field, val)
		}
	}
    

	jsonStr2, err := json.Marshal(mapped)
	if err != nil {
		return nil, err
	}

	result := string(jsonStr2)
	return &result, nil
}



func (c *RestClient) SendSMS(args *SendSMSRestModel) (*RestResponse, error) {

	body, err := c.addCredentials(args)
	if err != nil {
		return nil, err
	}
	jsonBody := []byte(*body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/SendSMS", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res := RestResponse{}
	if err := c.callRestAPI(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}