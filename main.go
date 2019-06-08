package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max - min)
}

func generatePassMap(i int) func() (map[int]string, bool) {
	passAmount := 300
	passLen := i

	return func() (passwordsMap map[int]string, exist bool) {
		minPass := 1
		maxPass := 10

		for i := 1; i < passLen; i++ {
			minPass *= 10
			maxPass *= 10
		}

		maxPass -= 1

		passwordsMap = make(map[int]string)
		for i := 0; i < passAmount; i++ {
			pass := int(randInt(minPass, maxPass))

			_, passExist := passwordsMap[pass]

			if passExist == true {
				fmt.Println("Pass exist", pass)
				exist = true
				break
			} else {
				passwordsMap[pass] = ""
			}
		}

		return
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 1; ; i++ {
		a := generatePassMap(i)

		passwordsMap, exist := a()

		fmt.Println(passwordsMap)

		if exist == false {
			break
		}
	}
}
