package main

/*
blog.golang.org/errors-are-values
below is a technique to write a file without repeating error handling code
*/
/*
type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic")
	sw.writeln("make the zero value useful")
	return sw.err
}

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
*/
