package exec_v12

import (
	"fmt"
)

func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}

}

func GetValue() int {
	return 1
}
