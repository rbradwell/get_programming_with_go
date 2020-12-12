package main

/*
var (
	errBounds = errors.New("out of bounds")
	errDigit  = errors.New("invalid digit")
)

const rows, columns = 9, 9

// The error type is a built in interface.  Any type that implements Error() method that returns a string satisfies the interface.
type sudokuError []error

type grid [rows][columns]int8

func (se sudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *grid) Set(row, column int, digit int8) error {
	var errs sudokuError
	if !inBounds(row, column) {
		errs = append(errs, errBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, errDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}

func main() {
	var g grid
	err := g.Set(10, 0, 5)
	if err != nil {
		if errs, ok := err.(sudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}
	os.Exit(1)
}
*/
