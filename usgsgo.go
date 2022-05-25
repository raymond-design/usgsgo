package usgsgo

import "net/http"

// client is the default client for all requests.
var client = &http.Client{}

// SetClient the http client to use for all requests.
func SetClient(c *http.Client) {
	client = c
}

// Call makes a request to the Ivs API.
func Call(queryParams interface{}) err {
	if queryParams != nil {
		q, err := query.Encode(queryParams)
		if err != nil {
			return err
		}
	}
}
