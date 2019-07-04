package mypackage

import (
	"fmt"
	"log"
	"os"
)

func ErrorCheckFatal(err error) {
	if nil != err {
		log.Fatal(err)
	}
}
func ErrorCheck(err error, err_information string) {
	if nil != err {
		fmt.Println(err_information)
		os.Exit(0)
	}
}
func AddrExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}
}
