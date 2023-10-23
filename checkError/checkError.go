package checkError

import (
	"fmt"
	"log"
)

func ErrCheck(err error) {
	if err != nil {
		fmt.Printf("theres an errror: ")
		log.Fatal(err)
	}
}
