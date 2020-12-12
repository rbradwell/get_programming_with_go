// In Go there are no classes or objects

package main

type celsius float64
type kelvin float64

/*
// declaring a function on the kelvin type
// below (k kelvin) is the receiver.  Methods must have exactly one receiver

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	var c celsius
	c = k.celsius() // call celsius on the kelvin type
	fmt.Print(k, " degrees kelvin is ", c, " degrees celcius")
}
*/
