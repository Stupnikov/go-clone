{
	if len(p) != 2 {
		return EINVAL
	}
	var pp [2]_C_int
	err = pipe2(&pp, 0) // pipe2 is the same as pipe when flags are set to 0.
	p[0] = int(pp[0])
	p[1] = int(pp[1])
	return
}