package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(t.Date())
	fmt.Println(t.Clock())
}
