////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Wrapper for Send command
package udb

import (
	jww "github.com/spf13/jwalterweatherman"
	client "gitlab.com/privategrity/client/api"
	"gitlab.com/privategrity/crypto/format" // <-- FIXME: this is annoying, WHY?
)

// Wrap the API Send function (useful for mock tests)
func Send(userId uint64, msg string) {
	myId := uint64(UDB_USERID)
	messages, err := format.NewMessage(myId, userId, msg)
	if err != nil {
		jww.FATAL.Panicf("Error creating message: %d, %d, %s",
			myId, userId, msg)
	}

	for i := range messages {
		sendErr := client.Send(messages[i])
		if sendErr != nil {
			jww.ERROR.Printf("Error responding to %d", userId)
		}
	}
}
