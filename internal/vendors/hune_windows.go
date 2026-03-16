//go:build windows

package vendors

import (
	"fmt"
	"log"
	"olinker/internal/core"
	"olinker/internal/platform"
)

type HuneVendor struct {
	dll *platform.DLLLoader
}

func NewHuneVendor() *HuneVendor {
	return &HuneVendor{}
}

func (v *HuneVendor) Init(config core.VendorConfig) error {
	var err error
	v.dll, err = platform.NewDLLLoader(config.DLLPath)
	if err != nil {
		return err
	}
	log.Printf("Hune Vendor DLL loaded successfully from %s", config.DLLPath)
	return nil
}

func (v *HuneVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
	log.Printf("[Hune] Encoding card for room %s", req.RoomName)
	return core.EncodeResult{Status: 0, CardNo: "HUNE-1234", RoomName: req.RoomName}, nil
}

func (v *HuneVendor) CancelCard(cardID string) error {
	return fmt.Errorf("CancelCard not implemented for Hune")
}

func (v *HuneVendor) ExtendCard(req core.ExtendRequest) error {
	return fmt.Errorf("ExtendCard not implemented for Hune")
}

func (v *HuneVendor) ReadCard() (core.CardInfo, error) {
	return core.CardInfo{}, fmt.Errorf("ReadCard not implemented for Hune")
}
