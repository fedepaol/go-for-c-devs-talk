func getAdder(toSum int) func(int) int {
	return func(arg int) int  {
		return arg + toSum;
	}
}
