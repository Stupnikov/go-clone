{
	defer goRecover()
	return C.bool(matchaGoGet(v).Bool())
}