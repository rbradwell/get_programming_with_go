package main

/*
func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	// go ensures that deferred actions take place before the contining function returns.
	// every return statement that follws the defer below will result in the f.close() method being called
	// NOTE: any method or function can be deferred and it isn't specifically for error handling.
	defer f.Close()
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}
	_, err = fmt.Fprintln(f, "Don't just check errors. handle them gracefully")
	return err
}
func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
*/
