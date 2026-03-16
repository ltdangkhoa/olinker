//go:build !windows

package vendors

import (
	"olinker/internal/core"
)

type dummyVendor struct{ name string }

func (v *dummyVendor) Init(config core.VendorConfig) error { return nil }
func (v *dummyVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
	return core.EncodeResult{Status: 0, CardNo: "STUB-MODE", RoomName: req.RoomName}, nil
}
func (v *dummyVendor) CancelCard(cardID string) error { return nil }
func (v *dummyVendor) ExtendCard(req core.ExtendRequest) error { return nil }
func (v *dummyVendor) ReadCard() (core.CardInfo, error) {
	return core.CardInfo{Status: 0, CardNo: "STUB", RoomName: "101"}, nil
}

func NewOrbitaVendor() LockVendor { return &dummyVendor{name: "Orbita"} }
func NewBeTechVendor() LockVendor { return &dummyVendor{name: "BeTech"} }
func NewAdelVendor() LockVendor   { return &dummyVendor{name: "Adel"} }
func NewHuneVendor() LockVendor   { return &dummyVendor{name: "Hune"} }
func NewProUSBVendor() LockVendor { return &dummyVendor{name: "ProUSB"} }
func NewDLockVendor() LockVendor  { return &dummyVendor{name: "DLock"} }
