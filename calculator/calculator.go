package calculator

func Add(a int, b int) int {
	var c = a + b
	return c
}
func Subtract(a, b int) int {
	return (a - b)
}
func Mul(a, b int) int {
	return (a * b)

}
func Divison(a int, b int) (int, int) {
	var c, d = a / b, a % b
	return c, d
}
