package pkg1

import (
	"fmt"
	_ "go101/init/pkg2"
)

func init() {
	fmt.Println("pkg1 init")
}
