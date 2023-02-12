package netconnect

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpProtocolObj struct {
	Hclient *http.Client
	Hlogger *Logger
}

func NewHttpProtocol(urlstr string) (*HttpProtocolObj, error) {
	// HTTPEndpointConnect
	hlog, err := NewLogger("HttpProtocol")
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	Hpclient := HttpProtocolObj{
		Hclient: client,
		Hlogger: hlog,
	}
	return &Hpclient, nil
}

//http1.0 支持keep alive 可以重用

func (Hpo *HttpProtocolObj) SendRequest(endpoint, method string, requestBody map[string]interface{}) ([]byte, error) {

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		// log.Fatalf("Error Occurred. %+v", err)
		return nil, err
	}

	response, err := Hpo.Hclient.Do(req)
	if err != nil {
		// log.Fatalf("Error sending request to API endpoint. %+v", err)
		return nil, err
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// log.Fatalf("Couldn't parse response body. %+v", err)
		Hpo.Hlogger.Log.Infof("Couldn't parse response body. %+v", err)
		return nil, err
	}

	return body, nil
}
