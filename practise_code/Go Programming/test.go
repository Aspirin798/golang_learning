/* package main

import (
	"fmt"
)

func main() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}

}
*/

/* package main

import "errors"

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("should be non-negative numbers!")
		return
	}
	return a + b, nil
}

func main() {
	c, er := Add(-1, 2)
	println(c, er)
}
*/

package main

import (
	"fmt"
)

func main() {
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}

	}()
	a()
	j *= 2
	a()
}
