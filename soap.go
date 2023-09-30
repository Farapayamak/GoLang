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
	smartURL		string

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
		smartURL: "https://api.payamak-panel.com/post/Smartsms.asmx/%s?",
		username: username,
		password: password,
		debug: false,
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



// Receive web service methods

func (c *SoapClient) ChangeMessageIsRead(args *ChangeMessageIsReadSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "ChangeMessageIsRead", args)
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


func (c *SoapClient) GetInboxCount(args *GetInboxCountSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "GetInboxCount", args)
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


func (c *SoapClient) GetLatestReceiveMsg(args *GetLatestReceiveMsgSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "GetLatestReceiveMsg", args)
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


func (c *SoapClient) GetMessages(args *GetMessagesSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "GetMessages", args)
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


func (c *SoapClient) GetMessagesAfterID(args *GetMessagesAfterIDSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "GetMessagesAfterID", args)
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


func (c *SoapClient) GetMessagesFilterByDate(args *GetMessagesFilterByDateSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "GetMessagesFilterByDate", args)
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


func (c *SoapClient) GetMessagesReceptions(args *GetMessagesReceptionsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "GetMessagesReceptions", args)
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


func (c *SoapClient) GetMessagesWithChangeIsRead(args *GetMessagesWithChangeIsReadSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "GetMessagesWithChangeIsRead", args)
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


func (c *SoapClient) GetOutBoxCount() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.receiveURL, "GetOutBoxCount", args)
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


func (c *SoapClient) RemoveMessages(args *RemoveMessagesSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.receiveURL, "RemoveMessages", args)
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


// Users web service methods

func (c *SoapClient) AddUser(args *AddUserSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "AddUser", args)
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


func (c *SoapClient) AddUserWithLocation(args *AddUserWithLocationSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "AddUserWithLocation", args)
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


func (c *SoapClient) AddUserWithMobileNumber(args *AddUserWithMobileNumberSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "AddUserWithMobileNumber", args)
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


func (c *SoapClient) AddUserWithMobileNumber2(args *AddUserWithMobileNumber2SoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "AddUserWithMobileNumber2", args)
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


func (c *SoapClient) AddUserWithUserNameAndPass(args *AddUserWithUserNameAndPassSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "AddUserWithUserNameAndPass", args)
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


func (c *SoapClient) AuthenticateUser() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.usersURL, "AuthenticateUser", args)
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


func (c *SoapClient) ChangeUserCredit(args *ChangeUserCreditSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "ChangeUserCredit", args)
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


func (c *SoapClient) DeductUserCredit(args *DeductUserCreditSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "DeductUserCredit", args)
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


func (c *SoapClient) ForgotPassword(args *ForgotPasswordSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "ForgotPassword", args)
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


func (c *SoapClient) GetCities(args *GetCitiesSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "GetCities", args)
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


func (c *SoapClient) GetEnExpireDate() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.usersURL, "GetEnExpireDate", args)
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


func (c *SoapClient) GetExpireDate() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.usersURL, "GetExpireDate", args)
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


func (c *SoapClient) GetProvinces() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.usersURL, "GetProvinces", args)
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


func (c *SoapClient) GetUserBasePrice(args *GetUserBasePriceSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserBasePrice", args)
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


func (c *SoapClient) GetUserCredit(args *GetUserCreditSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserCredit", args)
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


func (c *SoapClient) GetUserCredit2() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserCredit2", args)
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


func (c *SoapClient) GetUserDetails(args *GetUserDetailsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserDetails", args)
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


func (c *SoapClient) GetUserIsExist(args *GetUserIsExistSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserIsExist", args)
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


func (c *SoapClient) GetUserNumbers() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserNumbers", args)
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


func (c *SoapClient) GetUserTransactions(args *GetUserTransactionsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserTransactions", args)
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


func (c *SoapClient) GetUserWallet() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserWallet", args)
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


func (c *SoapClient) GetUserWalletTransaction(args *GetUserWalletTransactionSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUserWalletTransaction", args)
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


func (c *SoapClient) GetUsers() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.usersURL, "GetUsers", args)
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


func (c *SoapClient) RemoveUser(args *RemoveUserSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.usersURL, "RemoveUser", args)
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


// Voice web service methods

func (c *SoapClient) SendBulkSpeechText(args *SendBulkSpeechTextSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.voiceURL, "SendBulkSpeechText", args)
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


func (c *SoapClient) SendBulkVoiceSMS(args *SendBulkVoiceSMSSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.voiceURL, "SendBulkVoiceSMS", args)
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


func (c *SoapClient) UploadVoiceFile(args *UploadVoiceFileSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.voiceURL, "UploadVoiceFile", args)
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


// Contacts web service methods

func (c *SoapClient) AddContact(args *AddContactSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "AddContact", args)
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


func (c *SoapClient) AddContactEvents(args *AddContactEventsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "AddContactEvents", args)
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


func (c *SoapClient) AddGroup(args *AddGroupSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "AddGroup", args)
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


func (c *SoapClient) ChangeContact(args *ChangeContactSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "ChangeContact", args)
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


func (c *SoapClient) ChangeGroup(args *ChangeGroupSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "ChangeGroup", args)
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


func (c *SoapClient) CheckMobileExistInContact(args *CheckMobileExistInContactSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "CheckMobileExistInContact", args)
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


func (c *SoapClient) GetContactEvents(args *GetContactEventsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "GetContactEvents", args)
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


func (c *SoapClient) GetContacts(args *GetContactsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "GetContacts", args)
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


func (c *SoapClient) GetContactsByID(args *GetContactsByIDSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "GetContactsByID", args)
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


func (c *SoapClient) GetGroups() (*string, error) {

	var args interface{}
	urlWithParams, err := c.setQueryParams(c.contactsURL, "GetGroups", args)
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


func (c *SoapClient) MergeGroups(args *MergeGroupsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "MergeGroups", args)
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


func (c *SoapClient) RemoveContact(args *RemoveContactSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "RemoveContact", args)
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


func (c *SoapClient) RemoveContactByContactID(args *RemoveContactByContactIDSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "RemoveContactByContactID", args)
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


func (c *SoapClient) RemoveGroup(args *RemoveGroupSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.contactsURL, "RemoveGroup", args)
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

// Schedule web service methods

func (c *SoapClient) AddNewMultipleSchedule(args *AddNewMultipleScheduleSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.scheduleURL, "AddNewMultipleSchedule", args)
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

func (c *SoapClient) AddNewUsance(args *AddNewUsanceSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.scheduleURL, "AddNewUsance", args)
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


func (c *SoapClient) AddSchedule(args *AddScheduleSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.scheduleURL, "AddSchedule", args)
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


func (c *SoapClient) GetScheduleDetails(args *GetScheduleDetailsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.scheduleURL, "GetScheduleDetails", args)
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


func (c *SoapClient) GetScheduleStatus(args *GetScheduleStatusSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.scheduleURL, "GetScheduleStatus", args)
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


func (c *SoapClient) RemoveSchedule(args *RemoveScheduleSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.scheduleURL, "RemoveSchedule", args)
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

// Bulks web service methods

func (c *SoapClient) AddNumberBulk(args *AddNumberBulkSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.bulksURL, "AddNumberBulk", args)
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


func (c *SoapClient) BulkReception(args *BulkReceptionSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.bulksURL, "BulkReception", args)
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


func (c *SoapClient) BulkReceptionCount(args *BulkReceptionCountSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.bulksURL, "BulkReceptionCount", args)
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


func (c *SoapClient) GetBulkDeliveries(args *GetBulkDeliveriesSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.bulksURL, "GetBulkDeliveries", args)
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


func (c *SoapClient) GetBulkDeliveries2(args *GetBulkDeliveries2SoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.bulksURL, "GetBulkDeliveries2", args)
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


func (c *SoapClient) GetBulkDetails(args *GetBulkDetailsSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.bulksURL, "GetBulkDetails", args)
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


func (c *SoapClient) SendSmartSMS(args *SendSmartSMSSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.smartURL, "SendSmartSMS", args)
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


func (c *SoapClient) SendMultipleSmartSMS(args *SendMultipleSmartSMSSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.smartURL, "SendMultipleSmartSMS", args)
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


func (c *SoapClient) GetSmartSMSDeliveries(args *GetSmartDeliveriesSoapModel) (*string, error) {

	urlWithParams, err := c.setQueryParams(c.smartURL, "GetSmartSMSDeliveries", args)
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