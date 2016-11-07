package main

func f() (int, int) {
	return 5, 6
}
func main() {
	int x; int y;
	x, y := f()
}
