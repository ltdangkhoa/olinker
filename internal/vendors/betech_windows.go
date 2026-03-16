//go:build windows

package vendors

import (
	"fmt"
	"log"
	"olinker/internal/core"
	"olinker/internal/platform"
)

type BeTechVendor struct {
	dll *platform.DLLLoader
}

func NewBeTechVendor() *BeTechVendor {
	return &BeTechVendor{}
}

func (v *BeTechVendor) Init(config core.VendorConfig) error {
	var err error
	v.dll, err = platform.NewDLLLoader(config.DLLPath)
	if err != nil {
		return err
	}
	log.Printf("BeTech Vendor DLL loaded successfully from %s", config.DLLPath)
	return nil
}

func (v *BeTechVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
	log.Printf("[BeTech] Encoding card for room %s", req.RoomName)
	return core.EncodeResult{Status: 0, CardNo: "BETECH-1234", RoomName: req.RoomName}, nil
}

func (v *BeTechVendor) CancelCard(cardID string) error {
	return fmt.Errorf("CancelCard not implemented for BeTech")
}

func (v *BeTechVendor) ExtendCard(req core.ExtendRequest) error {
	return fmt.Errorf("ExtendCard not implemented for BeTech")
}

func (v *BeTechVendor) ReadCard() (core.CardInfo, error) {
	return core.CardInfo{}, fmt.Errorf("ReadCard not implemented for BeTech")
}
