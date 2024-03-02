package watch

func (ctl *Watchctl) AddCallback(callback func()) {
	ctl.callbacks = append(ctl.callbacks, callback)
}

func (ctl *Watchctl) triggerCallbacks() {
	for _, fnc := range ctl.callbacks {
		fnc()
	}
}
