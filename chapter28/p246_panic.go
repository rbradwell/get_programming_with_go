package main

/*
func main() {

	// deferred function are executed before a function returns, even in the case of a panic
	// if a deferred function calls recover, the panic will stop, and the program will continue running
	// recovery is similar to catch etc in other languages
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	// GO has panic which is a similar mechanism to exceptions
	// there are some situations where Go will panic instead of providing an error value, such as divide by zero
	panic("I forgot my towel")
}
*/
