package udb

import (
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/crypto/hash"
	"gitlab.com/elixxir/user-discovery-bot/storage"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/primitives/id"
	"testing"
	"time"
)

// Check that our nil check works
func TestDeleteFact_NilCheck(t *testing.T) {
	_, err := DeleteFact(nil, nil, nil, nil)
	if err == nil {
		t.Error("DeleteFact receiving a nil msg didn't error")
	}

	badmsg := pb.FactRemovalRequest{
		UID:         nil,
		RemovalData: nil,
	}
	_, err = DeleteFact(&badmsg, nil, nil, nil)
	if err == nil {
		t.Error("DeleteFact receiving a nil msg didn't error")
	}
}

// Check that our auth check works
func TestDeleteFact_AuthCheck(t *testing.T) {
	// Make a FactRemovalRequest to put into the Delete function
	badmsg := pb.FactRemovalRequest{
		UID: id.DummyUser.Marshal(),
		RemovalData: &pb.Fact{
			Fact:     "Testing",
			FactType: 0,
		},
	}

	// Make a new host for auth
	h, err := connect.NewHost(&id.DummyUser, "0.0.0.0:0", nil,
		connect.HostParams{MaxRetries: 0, AuthEnabled: false})
	if err != nil {
		t.Error(err)
	}
	input_auth := connect.Auth{
		IsAuthenticated: false,
		Sender:          h,
		Reason:          "",
	}

	_, err = DeleteFact(&badmsg, nil, nil, &input_auth)
	if err == nil {
		t.Error("DeleteFact receiving a nil msg didn't error")
	}
}

// Checks test for no user having registered the fact
func TestDeleteFact_UsersCheck(t *testing.T) {
	// Make a FactRemovalRequest to put into the Delete function
	badmsg := pb.FactRemovalRequest{
		UID: id.DummyUser.Marshal(),
		RemovalData: &pb.Fact{
			Fact:     "Testing",
			FactType: 0,
		},
	}

	// Make a new host for auth
	h, err := connect.NewHost(&id.DummyUser, "0.0.0.0:0", nil,
		connect.HostParams{MaxRetries: 0, AuthEnabled: true})
	if err != nil {
		t.Error(err)
	}
	input_auth := connect.Auth{
		IsAuthenticated: true,
		Sender:          h,
		Reason:          "",
	}

	// Setup a Storage object
	store, _, err := storage.NewDatabase("", "", "", "", "")
	if err != nil {
		t.Fatal("connect.NewHost returned an error: ", err)
	}

	_, err = DeleteFact(&badmsg, nil, store, &input_auth)
	if err == nil {
		t.Error("DeleteFact receiving a nil msg didn't error")
	}
}

// Test that the function doesn't work when a different user to the Fact owner tries to delete the Fact
func TestDeleteFact_WrongOwner(t *testing.T) {
	// Create an input message
	input_msg := pb.FactRemovalRequest{
		UID: []byte{0, 1, 2, 3},
		RemovalData: &pb.Fact{
			Fact:     "Testing",
			FactType: 0,
		},
	}

	// Create a new host and an auth object for it
	h, err := connect.NewHost(&id.DummyUser, "0.0.0.0:0", nil,
		connect.HostParams{MaxRetries: 0, AuthEnabled: true})
	if err != nil {
		t.Fatal("connect.NewHost returned an error: ", err)
	}
	input_auth := connect.Auth{
		IsAuthenticated: true,
		Sender:          h,
		Reason:          "",
	}

	// Setup a Storage object
	store, _, err := storage.NewDatabase("", "", "", "", "")
	if err != nil {
		t.Fatal("connect.NewHost returned an error: ", err)
	}

	// Insert a user to assign the Fact to
	suser_factstorage := new([]storage.Fact)
	suser := storage.User{
		Id:        id.NewIdFromUInt(0, id.User, t).Marshal(),
		RsaPub:    "",
		DhPub:     nil,
		Salt:      nil,
		Signature: nil,
		Facts:     *suser_factstorage,
	}
	err = store.InsertUser(&suser)
	if err != nil {
		t.Fatal(err)
	}

	// Create a Fact object to put into our Storage object
	// Generate the hash function and hash the fact
	sfhash, err := hash.NewCMixHash()
	if err != nil {
		t.Fatal(err)
	}
	sfhash.Write(input_msg.RemovalData.Digest())
	hashFact := sfhash.Sum(nil)
	sfact := storage.Fact{
		Hash:         hashFact,
		UserId:       id.NewIdFromUInt(0, id.User, t).Marshal(),
		Fact:         "Testing",
		Type:         0,
		Signature:    nil,
		Verified:     false,
		Timestamp:    time.Time{},
		Verification: storage.TwilioVerification{},
	}
	err = store.InsertFact(&sfact)
	if err != nil {
		t.Fatal(err)
	}

	// Attempt to delete our Fact
	_, err = DeleteFact(&input_msg, nil, store, &input_auth)
	if err == nil {
		t.Error("DeleteFact did not return an error deleting a fact the input user doesn't own")
	}
}

// Test that the function does work given the right inputs and DB entries
func TestDeleteFact_Happy(t *testing.T) {
	// Create an input message
	input_msg := pb.FactRemovalRequest{
		UID: []byte{0, 1, 2, 3},
		RemovalData: &pb.Fact{
			Fact:     "Testing",
			FactType: 0,
		},
	}

	// Create a new host and an auth object for it
	h, err := connect.NewHost(&id.DummyUser, "0.0.0.0:0", nil,
		connect.HostParams{MaxRetries: 0, AuthEnabled: true})
	if err != nil {
		t.Fatal("connect.NewHost returned an error: ", err)
	}
	input_auth := connect.Auth{
		IsAuthenticated: true,
		Sender:          h,
		Reason:          "",
	}

	// Setup a Storage object
	store, _, err := storage.NewDatabase("", "", "", "", "")
	if err != nil {
		t.Fatal("connect.NewHost returned an error: ", err)
	}

	// Insert a user to assign the Fact to
	suser_factstorage := new([]storage.Fact)
	suser := storage.User{
		Id:        id.DummyUser.Marshal(),
		RsaPub:    "",
		DhPub:     nil,
		Salt:      nil,
		Signature: nil,
		Facts:     *suser_factstorage,
	}
	err = store.InsertUser(&suser)
	if err != nil {
		t.Fatal(err)
	}

	// Create a Fact object to put into our Storage object
	// Generate the hash function and hash the fact
	sfhash, err := hash.NewCMixHash()
	if err != nil {
		t.Fatal(err)
	}
	sfhash.Write(input_msg.RemovalData.Digest())
	hashFact := sfhash.Sum(nil)
	sfact := storage.Fact{
		Hash:         hashFact,
		UserId:       id.DummyUser.Marshal(),
		Fact:         "Testing",
		Type:         0,
		Signature:    nil,
		Verified:     false,
		Timestamp:    time.Time{},
		Verification: storage.TwilioVerification{},
	}
	err = store.InsertFact(&sfact)
	if err != nil {
		t.Fatal(err)
	}

	// Attempt to delete our Fact
	_, err = DeleteFact(&input_msg, nil, store, &input_auth)
	if err != nil {
		t.Error("DeleteFact returned an error: ", err)
	}
}
