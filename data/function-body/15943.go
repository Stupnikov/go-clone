{
	_, _, e1 := Syscall(SYS_FSTAT, uintptr(fd), uintptr(unsafe.Pointer(st)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}