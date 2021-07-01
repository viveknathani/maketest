package utils

import "fmt"

const (
	TYPE_INTEGER = uint8(0)
	TYPE_REAL    = uint8(1)
	TYPE_STRING  = uint8(2)
	TYPE_DATE    = uint8(3)
	TYPE_UUID    = uint8(4)
)

func PrintMenu() {
	fmt.Println("0. INTEGER")
	fmt.Println("1. FLOATING-POINT")
	fmt.Println("2. STRING")
	fmt.Println("3. DATE")
	fmt.Println("4. UUID")
}
