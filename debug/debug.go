package debug

import (
	"fmt"
	"iter"
)

// DebugPrint prints a string to stdout.
//
// Parameters:
//   - title: the title of the string.
//   - fn: the function to print.
//
// Returns:
//   - error: any error that may have occurred.
func DebugPrint(title string, fn func() iter.Seq[string]) error {
	_, err := fmt.Println("[DEBUG]: " + title)
	if err != nil {
		return err
	}

	_, err = fmt.Println()
	if err != nil {
		return err
	}

	if fn == nil {
		return nil
	}

	for str := range fn() {
		_, err = fmt.Println(str)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Println()
	if err != nil {
		return err
	}

	return nil
}
