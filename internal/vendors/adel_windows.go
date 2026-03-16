//go:build windows

package vendors

import (
	"fmt"
	"log"
	"olinker/internal/core"
	"olinker/internal/platform"
)

type AdelVendor struct {
	dll *platform.DLLLoader
}

func NewAdelVendor() *AdelVendor {
	return &AdelVendor{}
}

func (v *AdelVendor) Init(config core.VendorConfig) error {
	var err error
	v.dll, err = platform.NewDLLLoader(config.DLLPath)
	if err != nil {
		return err
	}
	log.Printf("Adel Vendor DLL loaded successfully from %s", config.DLLPath)
	return nil
}

func (v *AdelVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
	proc, err := v.dll.GetProc("NewKey")
	if err != nil {
		log.Printf("[Adel] NewKey function not found in DLL: %v", err)
		return core.EncodeResult{}, err
	}
	
	log.Printf("[Adel] Calling Adel_NewKey for room %s (stub implementation)", req.RoomName)
	// TODO: map arguments using windows.BytePtrFromString and proc.Call
	_ = proc

	return core.EncodeResult{Status: 0, CardNo: "ADEL-1234", RoomName: req.RoomName}, nil
}

func (v *AdelVendor) CancelCard(cardID string) error {
	return fmt.Errorf("CancelCard not implemented for Adel")
}

func (v *AdelVendor) ExtendCard(req core.ExtendRequest) error {
	return fmt.Errorf("ExtendCard not implemented for Adel")
}

func (v *AdelVendor) ReadCard() (core.CardInfo, error) {
	return core.CardInfo{}, fmt.Errorf("ReadCard not implemented for Adel")
}
