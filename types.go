package nprotoo

import (
	"encoding/json"
	"time"
)

// AcceptFunc .
type AcceptFunc func(data json.RawMessage)

// RejectFunc .
type RejectFunc func(errorCode int, errorReason string)

// RequestFunc .
type RequestFunc func(request map[string]interface{}, accept AcceptFunc, reject RejectFunc)

// BroadCastFunc .
type BroadCastFunc func(data map[string]interface{}, subj string)

type PeerMsg struct {
	Request      bool `json:"request"`
	Response     bool `json:"response"`
	Ok           bool `json:"ok"`
	Notification bool `json:"notification"`
}

/*
* Request
{
  request : true,
  id      : 12345678,
  method  : 'chatmessage',
  data    :
  {
    type  : 'text',
    value : 'Hi there!'
  }
}
*/
type Request struct {
	Request   bool            `json:"request"`
	ID        int             `json:"id"`
	ReplySubj string          `json:"reply"`
	Method    string          `json:"method"`
	Data      json.RawMessage `json:"data"`
}

/*
* Success response
{
	response : true,
	id       : 12345678,
	ok       : true,
	data     :
	{
	  foo : 'lalala'
	}
}
*/
type Response struct {
	Response    bool            `json:"response"`
	ID          int             `json:"id"`
	Ok          bool            `json:"ok"`
	Data        json.RawMessage `json:"data"`
	ErrorCode   int             `json:"errorCode"`
	ErrorReason string          `json:"errorReason"`
}

/*
* Error response
{
  response    : true,
  id          : 12345678,
  ok          : false,
  errorCode   : 123,
  errorReason : 'Something failed'
}
*/
type ResponseError struct {
	Response    bool   `json:"response"`
	ID          int    `json:"id"`
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"errorCode"`
	ErrorReason string `json:"errorReason"`
}

/*
* Notification
{
  notification : true,
  method       : 'chatmessage',
  data         :
  {
    foo : 'bar'
  }
}
*/
type Notification struct {
	Notification bool            `json:"notification"`
	Method       string          `json:"method"`
	Data         json.RawMessage `json:"data"`
}

// Transcation .
type Transcation struct {
	id     int
	accept AcceptFunc
	reject RejectFunc
	close  func()
	timer  *time.Timer
}
