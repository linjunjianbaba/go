package main

import (
	"fmt"
	"net/http"
)

func main01() {
	url := "https://www.manhuatai.com/all.html"
	rest, err := http.Get(url)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer rest.Body.Close()

	buf := make([]byte, 4*1024)

	var resp string
	for {
		n, _ := rest.Body.Read(buf)
		if n == 0 {
			break
		}

		resp += string(buf[n:])

	}

	fmt.Println(resp)
}

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d = %d ", i, j, i*j)

		}
		fmt.Println()

	}

}
