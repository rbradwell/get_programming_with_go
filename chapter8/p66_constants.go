package main

/*
Go doesn't infer a type for constants unlike vars.
Calculations on constants and literals are performed at compilation time rather than runtime.
untyped numerical constants are backed by the 'big' package.
Constant values can be assigned to variables so long as they fit.
*/
/*
func main() {
	fmt.Println("andromeda galaxy is", 2400000000000000000/299792/86400, "light years away")
	const distance = 2400000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const days = distance / lightSpeed / secondsPerDay
	fmt.Println("andromeda galaxy is", days, "light years away")
}
*/
