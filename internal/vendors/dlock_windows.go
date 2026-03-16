//go:build windows

package vendors

import (
	"fmt"
	"log"
	"olinker/internal/core"
	"olinker/internal/platform"
)

type DLockVendor struct {
	dll *platform.DLLLoader
}

func NewDLockVendor() *DLockVendor {
	return &DLockVendor{}
}

func (v *DLockVendor) Init(config core.VendorConfig) error {
	var err error
	v.dll, err = platform.NewDLLLoader(config.DLLPath)
	if err != nil {
		return err
	}
	log.Printf("DLock Vendor DLL loaded successfully from %s", config.DLLPath)
	return nil
}

func (v *DLockVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
	log.Printf("[DLock] Encoding card for room %s", req.RoomName)
	return core.EncodeResult{Status: 0, CardNo: "DLOCK-1234", RoomName: req.RoomName}, nil
}

func (v *DLockVendor) CancelCard(cardID string) error {
	return fmt.Errorf("CancelCard not implemented for DLock")
}

func (v *DLockVendor) ExtendCard(req core.ExtendRequest) error {
	return fmt.Errorf("ExtendCard not implemented for DLock")
}

func (v *DLockVendor) ReadCard() (core.CardInfo, error) {
	return core.CardInfo{}, fmt.Errorf("ReadCard not implemented for DLock")
}
