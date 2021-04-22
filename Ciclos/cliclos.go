package main

import "fmt"

func main() {
	//Ciclo for
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//Ciclo while
	// j := 0
	// for j < 10 {
	// 	fmt.Println(j)
	// 	j++
	// }

	k := 0
	for {

		if k == 2 {
			k++
			continue //Vuelve al inicio del ciclo
		}

		fmt.Println(k)
		k++

		if k > 10 {
			break
		}
	}
}

//break corta el ciclo
