package main

//how to import many things
//packages that are unused are deleted when saved
import (
	"fmt"
	"math"

	"github.com/artytheparty/tutorial/third/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}
