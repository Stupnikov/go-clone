{
	d := initxattrdest(dest, 0)
	destsize := len(dest)

	nsid, a, err := xattrnamespace(attr)
	if err != nil {
		return -1, err
	}

	return ExtattrGetLink(link, nsid, a, uintptr(d), destsize)
}