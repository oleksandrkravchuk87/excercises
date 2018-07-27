package main

import "fmt"

func main() {

	var i, n int
	vis := make([]bool, 1000)
	var res []int
	vis[0] = true
	vis[1] = true

	i = 2

	for {
		if !vis[i] {
			res = append(res, i)
			n++

			for k := range vis {
				if k%i == 0 {
					vis[k] = true
				}
			}
		}

		i++

		if n > 25 {
			break
		}
	}
	fmt.Println(res)

}
