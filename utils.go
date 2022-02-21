package comiteame

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Error: ", err)
	os.Exit(1)
}

func Info(str string) {
	fmt.Println("INFO: ", str)
}
