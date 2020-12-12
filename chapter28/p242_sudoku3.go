package main

/*
// many GO packages declare and export variables from the errors that they could return.
// note: normally capital would be used to export however it caused the way we are working here
var (
	errBounds = errors.New("out of bounds")
	errDigit  = errors.New("invalid digit")
)

const rows, columns = 9, 9

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
	g[row][column] = digit
	return nil
}

func main() {
	var g grid
	err := g.Set(10, 0, 5)
	if err != nil {
		switch err {
		case errBounds, errDigit:
			fmt.Println("Parameters out of bounds")
		default:
			fmt.Println(err)
		}
	}
	os.Exit(1)
}
*/
