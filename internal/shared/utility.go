package shared

import "net/url"

const (
	clientString = "clientString"
)

func (cli *IvsClient) getDefaultHeaders(request string) url.Values {
	params := url.Values{}
	params.Add("request", request)
	params.Add("setContentType", "1")
	params.Add(clientString, cli.clientString)

	return params
}
