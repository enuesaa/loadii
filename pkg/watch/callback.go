package watch

func (ctl *Watchctl) AddCallback(callback func()) {
	ctl.callbacks = append(ctl.callbacks, callback)
}
