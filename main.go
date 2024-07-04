//
// proxrm
//
// main.go
//

package main

import (
	"flag"
	"fmt"
	"os"

	"proxrm/pkg/proxrm"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Printf("%s: <vmid>\n", os.Args[0])
	fmt.Printf("  %s --id <vmid>\n", os.Args[0])
}

func main() {
	var err error
	var vmid int
	var proxrmx *proxrm.ProxRm

	// usage if no args
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	flag.IntVar(&vmid, "id", 0, "vm id")

	flag.Parse()

	// vmid
	// TODO: add more logic to verify
	if vmid <= 0 {
		fmt.Printf("ERROR: id is %d", vmid)
		os.Exit(1)
	}

	// proxrm
	proxrmx = new(proxrm.ProxRm)

	// Run
	err = proxrmx.Run(vmid)
	if err != nil {
		fmt.Printf("ERROR: Run: %s\n", err)
		os.Exit(1)
	}

	// Exit
	os.Exit(0)
}
