package farapayamak


import (
	// "encoding/xml"
	"fmt"
	"net/http"
	"time"
	"bytes"
	"io"
)



type SoapClient struct {
	sendURL		string
	receiveURL	string

	httpClient	*http.Client
	username	string
	password	string
	debug		bool
}


func InitSoapClient(username string, password string) *SoapClient {
	return &SoapClient {
		httpClient: &http.Client {
			Timeout: 1 * time.Minute,
		},
		sendURL: "http://api.payamak-panel.com/post/send.asmx/%s?username=%s&password=%s",
		username: username,
		password: password,
		debug: true,
	}
}




func (c *SoapClient) callSoapAPI(req *http.Request) (*string, error) {
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// var response interface {}
	// if err = xml.NewDecoder(res.Body).Decode(&response); err != nil {
	// 	return nil, err
	// }

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := string(response);
	if c.debug {
		fmt.Printf("%+v\n", result)
	}

	return &result, nil
}


// Send web service methods

func (c *SoapClient) GetCredit() (*string, error) {

	args := ""
	// body, err := c.addCredentials(args)
	// if err != nil {
	// 	return nil, err
	// }

	// "GET" or http.MethodGet
	req, err := http.NewRequest("GET", fmt.Sprintf(c.sendURL, "GetCredit", c.username, c.password), bytes.NewReader([]byte(args)))
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil

}


func (c *SoapClient) GetDeliveries() (*string, error) {

	args := ""
	// body, err := c.addCredentials(args)
	// if err != nil {
	// 	return nil, err
	// }

	// "GET" or http.MethodGet
	req, err := http.NewRequest("GET", fmt.Sprintf(c.sendURL, "GetCredit", c.username, c.password), bytes.NewReader([]byte(args)))
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil

}