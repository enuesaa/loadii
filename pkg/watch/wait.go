package watch

func (ctl *Watchctl) Wait() {
	<-make(chan struct{})
}
