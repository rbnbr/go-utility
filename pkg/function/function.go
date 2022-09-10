package function

// MakeGetReturnElementAt
// Returns a function which, if provided multiple arguments, returns the argument at position index.
// May be used on functions that return multiple values but where we want to ignore all except one.
// Example: func do() (int64, string, error) { return 0, "", errors.New("error") }
// Example 1: MakeGetReturnElementAt(0)(do()).(int64) // yields the first element, i.e., just the int64 and ignores the string and the error (you will have to check if this can be done safely)
// Example 2: MakeGetReturnElementAt(-1)(do()).(error) // yields the last element, i.e., just the error and ignores the int64 and the string.
func MakeGetReturnElementAt(index int) func(...interface{}) interface{} {
	return func(i ...interface{}) interface{} {
		n := len(i)
		if index < 0 && -index <= n {
			index += n
		}
		return i[index]
	}
}

var GetLastReturnElement = MakeGetReturnElementAt(-1)
var GetFirstReturnElement = MakeGetReturnElementAt(0)
