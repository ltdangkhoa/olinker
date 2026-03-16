//go:build windows

package vendors

import (
	"fmt"
	"log"
	"olinker/internal/core"
	"olinker/internal/platform"
)

type OrbitaVendor struct {
	dll *platform.DLLLoader
}

func NewOrbitaVendor() *OrbitaVendor {
	return &OrbitaVendor{}
}

func (v *OrbitaVendor) Init(config core.VendorConfig) error {
	var err error
	v.dll, err = platform.NewDLLLoader(config.DLLPath)
	if err != nil {
		return err
	}
	log.Printf("Orbita Vendor DLL loaded successfully from %s", config.DLLPath)
	return nil
}

func (v *OrbitaVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
	log.Printf("[Orbita] Encoding card for room %s", req.RoomName)
	return core.EncodeResult{Status: 0, CardNo: "ORBITA-1234", RoomName: req.RoomName}, nil
}

func (v *OrbitaVendor) CancelCard(cardID string) error {
	return fmt.Errorf("CancelCard not implemented for Orbita")
}

func (v *OrbitaVendor) ExtendCard(req core.ExtendRequest) error {
	return fmt.Errorf("ExtendCard not implemented for Orbita")
}

func (v *OrbitaVendor) ReadCard() (core.CardInfo, error) {
	return core.CardInfo{}, fmt.Errorf("ReadCard not implemented for Orbita")
}
