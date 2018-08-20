{
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint16(rr.Priority, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint16(rr.Weight, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint16(rr.Port, msg, off)
	if err != nil {
		return off, err
	}
	off, err = PackDomainName(rr.Target, msg, off, compression, false)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}