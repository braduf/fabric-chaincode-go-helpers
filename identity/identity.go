package identity

import (
	"errors"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// Technical identity errors
// TODO: make errors constant
var (
	ErrDifferentMSPID = errors.New("Client identity MSP is different")
)

func AssertClientMSPID(cid cid.ClientIdentity, mspID string) (err error) {
	cidMSP, err := cid.GetMSPID()
	if err != nil {
		return
	}
	if cidMSP != mspID {
		err = ErrDifferentMSPID
	}
	return
}

func AssertClientOU(cid cid.ClientIdentity, ou string) bool {
	return true
}
