// Code generated by 'go generate'; DO NOT EDIT.

package dpapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modcrypt32 = windows.NewLazySystemDLL("crypt32.dll")

	procCryptProtectData   = modcrypt32.NewProc("CryptProtectData")
	procCryptUnprotectData = modcrypt32.NewProc("CryptUnprotectData")
)

func cryptProtectData(dataIn *dpBlob, name *uint16, optionalEntropy *dpBlob, reserved uintptr, promptStruct uintptr, flags uint32, dataOut *dpBlob) (err error) {
	r1, _, e1 := syscall.Syscall9(procCryptProtectData.Addr(), 7, uintptr(unsafe.Pointer(dataIn)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(optionalEntropy)), uintptr(reserved), uintptr(promptStruct), uintptr(flags), uintptr(unsafe.Pointer(dataOut)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func cryptUnprotectData(dataIn *dpBlob, name **uint16, optionalEntropy *dpBlob, reserved uintptr, promptStruct uintptr, flags uint32, dataOut *dpBlob) (err error) {
	r1, _, e1 := syscall.Syscall9(procCryptUnprotectData.Addr(), 7, uintptr(unsafe.Pointer(dataIn)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(optionalEntropy)), uintptr(reserved), uintptr(promptStruct), uintptr(flags), uintptr(unsafe.Pointer(dataOut)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
