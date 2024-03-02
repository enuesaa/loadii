package watch

func (ctl *Watchctl) Close() error {
	if ctl.watcher == nil {
		return nil
	}

	return ctl.watcher.Close()
}
