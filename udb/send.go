////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Wrapper for Send command
package udb

import (
	"gitlab.com/elixxir/client/api"
	"gitlab.com/elixxir/client/cmixproto"
	"gitlab.com/elixxir/client/parse"
	"gitlab.com/elixxir/primitives/id"
)

// Sender interface -- the api is broken here (does not return the error), so
// we reimplement a new interface...
type Sender interface {
	Send(recipientID *id.User, msg string) error
}

// ApiSender calls the api send function
type APISender struct{}

// Send calls the api send function
func (a APISender) Send(recipientID *id.User, msg string) error {
	return api.Send(api.APIMessage{
		Payload:     []byte(msg),
		SenderID:    UDB_USERID,
		RecipientID: recipientID,
	})
}

// UdbSender is the sender interface to use
var UdbSender Sender = APISender{}

// Wrap the API Send function (useful for mock tests)
func Send(recipientID *id.User, msg string, msgType cmixproto.Type) {
	// Create the message body and assign its type
	message := string(parse.Pack(&parse.TypedBody{
		InnerType: int32(msgType),
		Body:      []byte(msg),
	}))
	// Send the message
	sendErr := UdbSender.Send(recipientID, message)
	if sendErr != nil {
		Log.ERROR.Printf("Error responding to %d: %s", recipientID, sendErr)
	}
}
