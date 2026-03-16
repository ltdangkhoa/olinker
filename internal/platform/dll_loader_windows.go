//go:build windows

package platform

import (
	"fmt"
	"path/filepath"
	"golang.org/x/sys/windows"
)

type DLLLoader struct {
	dll *windows.DLL
}

func NewDLLLoader(path string) (*DLLLoader, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("could not resolve absolute path for %s: %w", path, err)
	}

	dll, err := windows.LoadDLL(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load DLL %s (resolved: %s): %w", path, absPath, err)
	}
	return &DLLLoader{dll: dll}, nil
}

func (l *DLLLoader) GetProc(procName string) (*windows.Proc, error) {
	return l.dll.FindProc(procName)
}

func (l *DLLLoader) Release() error {
	return l.dll.Release()
}
