//go:build !windows

package platform

import (
	"errors"
	"log"
)

type DLLLoader struct{}

func NewDLLLoader(path string) (*DLLLoader, error) {
	log.Printf("[STUB] DLL Loading is not supported on this platform. Path: %s", path)
	return &DLLLoader{}, nil
}

func (l *DLLLoader) GetProc(procName string) (interface{}, error) {
	return nil, errors.New("DLL procedures are not supported on this platform")
}

func (l *DLLLoader) Release() error {
	return nil
}
