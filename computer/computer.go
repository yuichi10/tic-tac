package computer

var neural *Neural

func initComputer() {
	neural := NewNeural()
	neural.LoadTheta()
}
