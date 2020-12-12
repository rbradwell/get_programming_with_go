package main

/*
type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func caliberate(s sensor, offset kelvin) sensor {
	// declare and return an anonymous function with a return type of kelvin
	// this is forming a closure so the function s passed in is called and the value it returns is offset
	return func() kelvin {
		return s() + offset
	}
}
func main() {
	func() {
		sensor := caliberate(realSensor, 5)
		fmt.Println(sensor())
	}()
}
*/
