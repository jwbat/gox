package inspect

import "fmt"


func Describe(v any) string {
	return fmt.Sprintf("Type: %T, Value: %v", v, v)
}
