package watch

func (ctl *Watchctl) triggerCallbacks() {
	for _, fnc := range ctl.options.Callbacks {
		fnc()
	}
}
