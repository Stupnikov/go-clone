{
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint16(rr.Type, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint16(rr.KeyTag, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint8(rr.Algorithm, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packStringBase64(rr.Certificate, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}