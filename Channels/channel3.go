package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "Deneme d1 d2 d3 d4"
	datasplit := strings.Split(data, " ")

	mychannel := make(chan string, len(datasplit))

	for _, splits := range datasplit {
		mychannel <- splits
	}

	close(mychannel)

	for msg := range mychannel {
		fmt.Print(msg + " ")
	}

	// for {
	// 	if msg, ok := <-mychannel; ok {
	// 		fmt.Print(msg + " ")
	// 	} else {
	// 		break
	// 	}
	// }

	// for i := 0; i < len(datasplit); i++ {
	// 	//fmt.Println(<-mychannel + " ")
	// 	fmt.Print(<-mychannel + " ")
	// }
}

// Deneme d1 d2 d3 d4 %
