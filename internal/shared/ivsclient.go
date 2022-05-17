package shared

import (
	"net/http"
	"net/url"
)

type AuthFunc func(string) url.Values

// IvsClient Contructor
type IvsClientConstructor struct {
	url          string
	clientString string
	headers      AuthFunc
	httpCli      *http.Client
}

type IvsClient struct {
	Url                         string
	httpClient                  *http.Client
	clientString                string
	headers                     AuthFunc
	sendParametersInRequestBody bool
}

// Create new IvsClient
func NewIvsClient(clientString string, httpCli *http.Client, headers AuthFunc) *IvsClientConstructor {
	ivsObj := &IvsClientConstructor{}

	return ivsObj
}

// Create a new test IvsClient
func NewTestIvsClient(clientString string, url string, httpCli *http.Client, headers AuthFunc) *IvsClientConstructor {
	ivsObj := &IvsClientConstructor{}

	return ivsObj
}

//Build the new IvsClient
func (ivs *IvsClientConstructor) Build() *IvsClient {
	if ivs.httpCli == nil {
		ivs.httpCli = &http.Client{}
	}

	newClient := &IvsClient{
		httpClient:   ivs.httpCli,
		clientString: ivs.clientString,
		headers:      ivs.headers,
	}

	if newClient.headers == nil {
		newClient.headers = newClient.getDefaultHeaders
	}

	if ivs.url == "" {
		if ivs.clientString != "" {
			newClient.Url = GetIvsUrl(ivs.clientString)
		} else {
			newClient.Url = GetIvsUrlFromAuthFunc(newClient.headers)
		}
	} else {
		newClient.Url = ivs.url
	}
	return newClient
}

func (ivs *IvsClient) SendParametersInRequestBody() {
	ivs.sendParametersInRequestBody = true
}

func (ivs *IvsClientConstructor) IvsWithURL(url string) {
	ivs.url = url
}

func (ivs *IvsClientConstructor) IvsWithClientCode(clientCode string) {
	ivs.clientString = clientCode
}

func (ivs *IvsClientConstructor) IvsWithHttpClient(httpCli *http.Client) {
	ivs.httpCli = httpCli
}

func (ivs *IvsClientConstructor) IvsWithHeaderFunc(headers AuthFunc) {
	ivs.headers = headers
}
