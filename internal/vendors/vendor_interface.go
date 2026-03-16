package vendors

import "olinker/internal/core"

type LockVendor interface {
	Init(config core.VendorConfig) error
	EncodeCard(req core.EncodeRequest) (core.EncodeResult, error)
	CancelCard(cardID string) error
	ExtendCard(req core.ExtendRequest) error
	ReadCard() (core.CardInfo, error)
}
