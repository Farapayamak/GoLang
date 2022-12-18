package farapayamak


import (
	"encoding/json"
	// "encoding/xml"
	"net/url"
	"fmt"
	"net/http"
	"time"
	"io"
	"reflect"
)



type SoapClient struct {
	sendURL			string
	receiveURL		string
	usersURL		string
	voiceURL		string
	contactsURL 	string
	scheduleURL		string
	bulksURL		string

	httpClient		*http.Client
	username		string
	password		string
	debug			bool
}


func InitSoapClient(username string, password string) *SoapClient {
	return &SoapClient {
		httpClient: &http.Client {
			Timeout: 1 * time.Minute,
		},
		sendURL: "http://api.payamak-panel.com/post/send.asmx/%s?",
		receiveURL: "http://api.payamak-panel.com/post/receive.asmx/%s?",
		usersURL: "http://api.payamak-panel.com/post/Users.asmx/%s?",
		voiceURL: "http://api.payamak-panel.com/post/Voice.asmx/%s?",
		contactsURL: "http://api.payamak-panel.com/post/contacts.asmx/%s?",
		scheduleURL: "http://api.payamak-panel.com/post/Schedule.asmx/%s?",
		bulksURL: "http://api.payamak-panel.com/post/newbulks.asmx/%s?",
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


func (c *SoapClient) setQueryParams(endpoint string, method string, data interface{}) (*string, error) {
	
	baseUrl := fmt.Sprintf(endpoint, method)

	var mapped map[string]interface{}
    jsonStr, err := json.Marshal(data)
	if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(jsonStr, &mapped); err != nil {
		return nil, err
	}
	// Avoid > panic: assignment to entry in nil map
	if len(mapped) == 0 {
		mapped = make(map[string]interface{})
	}

	params := url.Values{}
	params.Add("username", c.username)
	params.Add("password", c.password)

	for k, v := range mapped {
		
		if reflect.TypeOf(v).String() == "[]interface {}" {
			vv := v.([]interface{})
			for _, val := range vv {
				params.Add(k, fmt.Sprint(val))
			}
		} else {
			params.Add(k, fmt.Sprint(v))
		}
	}
	finalUrl := baseUrl + params.Encode()

	if c.debug {
		fmt.Println(finalUrl)
	}

	return &finalUrl , nil
}


// Send web service methods

func (c *SoapClient) GetCredit() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.sendURL, "GetCredit", args)
	if err != nil {
		return nil, err
	}

	// "GET" or http.MethodGet
	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) GetDeliveries(args *GetDeliveriesSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "GetDeliveries", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) GetDeliveries3(args *GetDeliveries3SoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "GetDeliveries3", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) GetSmsPrice(args *GetSmsPriceSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "GetSmsPrice", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendByBaseNumber(args *SendByBaseNumberSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendByBaseNumber", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendByBaseNumber2(args *SendByBaseNumber2SoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendByBaseNumber2", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendByBaseNumber3(args *SendByBaseNumber3SoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendByBaseNumber3", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendSimpleSMS(args *SendSimpleSMSSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendSimpleSMS", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendSimpleSMS2(args *SendSimpleSMS2SoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendSimpleSMS2", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendSms(args *SendSmsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendSms", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendSms2(args *SendSms2SoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendSms2", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendMultipleSMS(args *SendMultipleSMSSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendMultipleSMS", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *SoapClient) SendMultipleSMS2(args *SendMultipleSMS2SoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.sendURL, "SendMultipleSMS2", args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", *urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.callSoapAPI(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
