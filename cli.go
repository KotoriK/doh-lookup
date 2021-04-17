package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/KotoriK/doh-lookup/doh"
)

func main() {
	result, err := doh.Query("1.1.1.1", "www.baidu.com", "28", true, false)
	if err != nil {
		fmt.Println(err)
	} else {
		bytes, _ := json.Marshal(result)

		fmt.Println(string(bytes))
	}

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("done")
}
