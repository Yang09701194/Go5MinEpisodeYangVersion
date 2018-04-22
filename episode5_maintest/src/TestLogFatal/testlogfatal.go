package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	defer deferTest()

	fmt.Println("start printing error message")

	err := errors.New("fatal error message")
	log.Fatal(err)
	//不會被執行
	os.Exit(1)

	fmt.Println("finish printing error message")
	//time.Sleep
	time.Sleep(5 * time.Second)

}

func deferTest() {
	fmt.Println("test defer")
}

//err := errors.New("fatal error message")
// if err != nil {
// 		fmt.Print(err)
// }
