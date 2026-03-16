//go:build windows

package platform

import (
	"fmt"
	"golang.org/x/sys/windows"
)

type DLLLoader struct {
	dll *windows.DLL
}

func NewDLLLoader(path string) (*DLLLoader, error) {
	dll, err := windows.LoadDLL(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load DLL %s: %w", path, err)
	}
	return &DLLLoader{dll: dll}, nil
}

func (l *DLLLoader) GetProc(procName string) (*windows.Proc, error) {
	return l.dll.FindProc(procName)
}

func (l *DLLLoader) Release() error {
	return l.dll.Release()
}
