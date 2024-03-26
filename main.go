//
// proxrm
//
// main.go
//

package main

import (
	"fmt"
	"os"

	"proxrm/pkg/proxrm"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Printf("%s: <vmid>\n", os.Args[0])
}

func main() {
	var err error
	var proxrmx *proxrm.ProxRm

	// usage if no args
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	// proxrm
	proxrmx = new(proxrm.ProxRm)

	// Run
	err = proxrmx.Run()
	if err != nil {
		fmt.Printf("ERROR: Run: %s\n", err)
		os.Exit(1)
	}

	// Exit
	os.Exit(0)
}
