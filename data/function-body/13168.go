{
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	return ss.state.Update(state, hosting)
}