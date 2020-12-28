package cmix

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/elixxir/client/interfaces/message"
	"gitlab.com/elixxir/client/interfaces/params"
	"gitlab.com/elixxir/client/ud"
	"gitlab.com/xx_network/primitives/id"
)

func (m *Manager) LookupProcess() {
	for true {
		request := <-m.lookupChan

		lookupMsg := &ud.LookupSend{}
		if err := proto.Unmarshal(request.Payload, lookupMsg); err != nil {
			jww.ERROR.Printf("failed to unmarshal lookup "+
				"request from %s: %+v", request.Sender, err)
			continue
		}

		response := m.handleLookup(lookupMsg, request.Sender)

		marshaledResponse, err := proto.Marshal(response)
		if err != nil {
			jww.ERROR.Printf("failed to marshal responce "+
				"to request from %s: %+v", request.Sender, err)
			continue
		}

		responseMsg := message.Send{
			Recipient:   request.Sender,
			Payload:     marshaledResponse,
			MessageType: message.UdLookupResponse,
		}

		_, err = m.client.SendUnsafe(responseMsg, params.GetDefaultUnsafe())

		if err != nil {
			jww.ERROR.Printf("failed to send responce "+
				"to lookup request from %s: %+v", request.Sender, err)
		}
	}
}

func (m *Manager) handleLookup(msg *ud.LookupSend, requestor *id.ID) *ud.LookupResponse {

	response := &ud.LookupResponse{
		PubKey: nil,
		CommID: msg.CommID,
		Error:  "",
	}

	//decode the id to lookup
	lookupID, err := id.Unmarshal(msg.UserID)
	if err != nil {
		response.Error = fmt.Sprintf("failed to unmarshal lookup ID in "+
			"request from %s: %+v", requestor, err)
		jww.WARN.Println(response.Error)
		return response
	}

	//lookup the id
	usr, err := m.db.GetUser(lookupID.Bytes())
	if err != nil {
		response.Error = fmt.Sprintf("failed to lookup ID %s in "+
			"request from %s: %+v", lookupID, requestor, err)
		jww.WARN.Println(response.Error)
		return response
	}

	response.PubKey = usr.DhPub
	return response
}