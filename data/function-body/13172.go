{
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	return ss.tornDown
}