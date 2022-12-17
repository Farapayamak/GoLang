package farapayamak


import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
	// "bytes"
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




func (c *SoapClient) callSoapAPI(req *http.Request) (*interface{}, error) {
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	var response interface {}
	if err = xml.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("%+v\n", response)
	}

	return &response, nil
}



func (c *SoapClient) GetCredit() (*interface{}, error) {

	// body, err := c.addCredentials(args)
	// if err != nil {
	// 	return nil, err
	// }

	// "POST" or http.MethodPost
	req, err := http.NewRequest("POST", fmt.Sprintf(c.sendURL, "GetCredit"), nil) //bytes.NewReader([]byte(*body))
	if err != nil {
		return nil, err
	}

	// req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil

}