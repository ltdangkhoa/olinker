//go:build windows

package vendors

import (
	"fmt"
	"log"
	"olinker/internal/core"
	"olinker/internal/platform"
)

type ProUSBVendor struct {
	dll *platform.DLLLoader
}

func NewProUSBVendor() *ProUSBVendor {
	return &ProUSBVendor{}
}

func (v *ProUSBVendor) Init(config core.VendorConfig) error {
	var err error
	v.dll, err = platform.NewDLLLoader(config.DLLPath)
	if err != nil {
		return err
	}
	log.Printf("ProUSB Vendor DLL loaded successfully from %s", config.DLLPath)
	return nil
}

func (v *ProUSBVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
	log.Printf("[ProUSB] Encoding card for room %s", req.RoomName)
	return core.EncodeResult{Status: 0, CardNo: "PROUSB-1234", RoomName: req.RoomName}, nil
}

func (v *ProUSBVendor) CancelCard(cardID string) error {
	return fmt.Errorf("CancelCard not implemented for ProUSB")
}

func (v *ProUSBVendor) ExtendCard(req core.ExtendRequest) error {
	return fmt.Errorf("ExtendCard not implemented for ProUSB")
}

func (v *ProUSBVendor) ReadCard() (core.CardInfo, error) {
	return core.CardInfo{}, fmt.Errorf("ReadCard not implemented for ProUSB")
}
