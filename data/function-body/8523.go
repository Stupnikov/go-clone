{
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Host) + 1
	return l
}