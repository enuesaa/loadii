package usecase

func Sleep() {
	<-make(chan struct{})
}
