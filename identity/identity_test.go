package identity

import (
	"testing"

	"github.com/braduf/fabric-chaincode-go-helpers/mocking"
)

func TestAssertClientMSPID(t *testing.T) {
	t.Run("same MSP", func(t *testing.T) {
		msp := "OrgMSP"
		mockStub := mocking.NewMockChaincodeStub("TestAssertClient", nil, nil)
		mockCID := mocking.NewMockClientIdentity(mockStub, msp, nil, nil)

		err := AssertClientMSPID(mockCID, msp)
		assertError(t, err, nil)
	})

	t.Run("different MSP", func(t *testing.T) {
		cidMSP := "Org1MSP"
		wantedMSP := "Org2MSP"
		mockStub := mocking.NewMockChaincodeStub("TestAssertClient", nil, nil)
		mockCID := mocking.NewMockClientIdentity(mockStub, cidMSP, nil, nil)

		err := AssertClientMSPID(mockCID, wantedMSP)
		assertError(t, err, ErrDifferentMSPID)
	})
}

func TestAssertClientOU(t *testing.T) {

}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatalf("expected to get error %q", want)
	}
}
