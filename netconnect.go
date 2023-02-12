package netconnect

import "github.com/gorilla/websocket"

type WsConnect interface {
	Reconn() error
	ReadWsData()
	SubWsReadData() (chan []byte, error)
	SendJsonRequest(requestBody map[string]interface{}) error
	GetConnect() *websocket.Conn
}

type RPCEndPointConnect interface{}

type HTTPEndpointConnect interface {
	SendRequest(endpoint, method string, requestBody map[string]interface{}) ([]byte, error)
}
