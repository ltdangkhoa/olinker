package vendors

import (
	"fmt"
	"olinker/internal/core"
)

func LoadVendor(config core.VendorConfig) (LockVendor, error) {
	var v LockVendor

	switch config.Vendor {
	case "orbita":
		v = NewOrbitaVendor()
	case "betech":
		v = NewBeTechVendor()
	case "adel":
		v = NewAdelVendor()
	case "hune":
		v = NewHuneVendor()
	case "prousb":
		v = NewProUSBVendor()
	case "dlock":
		v = NewDLockVendor()
	case "mock":
		v = NewMockVendor()
	default:
		return nil, fmt.Errorf("unsupported vendor: %s", config.Vendor)
	}

	err := v.Init(config)
	if err != nil {
		return nil, err
	}

	return v, nil
}
