package vendors

import (
	"log"
	"olinker/internal/core"
)

type MockVendor struct{}

func NewMockVendor() *MockVendor {
	return &MockVendor{}
}

func (v *MockVendor) Init(config core.VendorConfig) error {
	log.Println("[Mock] Initialized with DLLPath:", config.DLLPath)
	return nil
}

func (v *MockVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
	log.Printf("[Mock] Encoding card for room %s", req.RoomName)
	return core.EncodeResult{Status: 0, CardNo: "MOCK-1234", RoomName: req.RoomName}, nil
}

func (v *MockVendor) CancelCard(cardID string) error {
	log.Printf("[Mock] Canceling card %s", cardID)
	return nil
}

func (v *MockVendor) ExtendCard(req core.ExtendRequest) error {
	log.Printf("[Mock] Extending card %s", req.CardID)
	return nil
}

func (v *MockVendor) ReadCard() (core.CardInfo, error) {
	log.Println("[Mock] Reading card")
	return core.CardInfo{Status: 0, CardNo: "MOCK-1234", RoomName: "101"}, nil
}
