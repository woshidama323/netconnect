package netconnect

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// ws offer such pool provider
type WebsocketInfra struct {
	ConEndpoints  string
	ReconnSignal  chan bool
	Wscon         *websocket.Conn
	WsLog         *zap.SugaredLogger
	WsHubChans    chan []byte
	WsFuncEnables map[string]string
	// StoreIT       store.StoreInterface
}

// WebsocketInfra will create a ws obj for management the ws connect
func NewWebsocketInfra(endpoint string, datachanclose bool, fEnalble ...string) (*WebsocketInfra, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	// u := url.URL{Scheme: "wss", Host: host, Path: path}
	// log.Printf("connecting to %s", u.String())
	sc, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		// log.Fatal("dial:", err)
		return nil, err
	}
	// needstop := make(chan bool)

	wsi := &WebsocketInfra{
		ConEndpoints: endpoint,
		ReconnSignal: make(chan bool),
		Wscon:        sc,
		WsLog:        logger.Sugar(),
		WsHubChans:   make(chan []byte, 10000),
		// StoreIT:      si,
	}

	if datachanclose {
		close(wsi.WsHubChans)
	}
	wsi.WsFuncEnables = map[string]string{}
	for _, e := range fEnalble {
		wsi.WsFuncEnables[e] = "true"
	}

	go wsi.ReadWsData()

	// go wsi.CreatePoolChan(wsi.ReconnSignal, needstop)

	return wsi, nil
}

func (Wsi *WebsocketInfra) GetConnect() *websocket.Conn {

	return Wsi.Wscon
}

func (Wsi *WebsocketInfra) Reconn() error {
	sc, _, err := websocket.DefaultDialer.Dial(Wsi.ConEndpoints, nil)
	if err != nil {
		return err
	}

	Wsi.Wscon = sc

	//reconn successful and send notification signal

	// notification need to be handle here , will notify if no reconn successful
	Wsi.ReconnSignal <- true
	return nil
}

func (Wsi *WebsocketInfra) ReadWsData() {

	// need to make sure Wsi.Wscon is not nil
	for {

		if Wsi.Wscon != nil {
			_, message, err := Wsi.Wscon.ReadMessage()
			// if err what to do ?
			if err != nil {
				//will recon
				Wsi.WsLog.Infof("read err: %+v", err)
				//close first and then sending reconnect signal
				err := Wsi.Wscon.Close()
				fmt.Println("close error:", err)
				err = Wsi.Reconn()
				if err != nil {
					fmt.Println("Reconn error:", err)
					// break
				}
			}
			//here need to recieve always ,so i need to send the message by
			Wsi.WsLog.Infof("recv: %+v\n", len(message))

			// d, err := Wsi.HandlesData(message)
			// if err != nil {
			// 	Wsi.WsLog.Errorf("HandlesData err: %+v", err)
			// 	d = ""
			// }
			select {
			case <-Wsi.WsHubChans: //不被chan 阻塞
				Wsi.WsLog.Info("Wsi.WsHubChans is closed")
			default:
				Wsi.WsLog.Info("Wsi.WsHubChans is not closed")
				Wsi.WsHubChans <- message
			}

		} else {
			Wsi.WsLog.Info("Wsi.Wscon nil,reconn now")
			err := Wsi.Reconn()
			if err != nil {
				fmt.Println("Wsi.Wscon nil,reconn,Reconn error:", err)
			}
			time.Sleep(1 * time.Second)
		}

	}
}

func (Wsi *WebsocketInfra) SubWsReadData() (chan []byte, error) {

	//外部会调用该接口，每调用一次就增加一个channel

	return Wsi.WsHubChans, nil
}

func (Wsi *WebsocketInfra) SendJsonRequest(requestBody map[string]interface{}) error {
	// if !Wsi.EnableChecks(GetFuncName(WebsocketInfra.SendTxPoolRequest)) {
	// 	return nil
	// }

	postBody, _ := json.Marshal(requestBody)
	Wsi.WsLog.Infof("start SendJsonRequest....")
	err := Wsi.Wscon.WriteMessage(websocket.TextMessage, postBody)
	if err != nil {
		Wsi.WsLog.Errorf("WriteMessage error:%+v", err)
		// Ticker.Stop()
		// time.Sleep(5 * time.Second)
		// Ticker = time.NewTicker(IntervalTime * time.Millisecond)
		if strings.EqualFold(err.Error(), "use of closed network connection") {
			sc, _, err := websocket.DefaultDialer.Dial(Wsi.ConEndpoints, nil)
			if err != nil {
				return err
			}
			Wsi.Wscon = sc
		}

	}
	return err
}

func (Wsi WebsocketInfra) EnableChecks(enable string) bool {

	if _, exists := Wsi.WsFuncEnables[enable]; exists {

		return true
	}

	return false
}

// GetNewHeads(string) (string, error)
func GetFuncName(f interface{}) string {
	fname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	strlist := strings.Split(fname, ".")
	name := strlist[len(strlist)-1]
	fmt.Println("function name is :", name)
	return name
}
