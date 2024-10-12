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

	"github.com/CTFfactory/proxrm/pkg/proxrm"
)

func usage() {
	fmt.Println("Usage:")

	fmt.Printf("  %s --id <vmid>\n", os.Args[0])
	fmt.Printf("  %s --name <vmname>\n", os.Args[0])
}

func main() {
	var err error
	var vmid int
	var vmname string
	var proxrmx *proxrm.ProxRm

	// usage if no args
	if len(os.Args) < 1 {
		usage()
		os.Exit(1)
	}

	// command line arguments
	flag.IntVar(&vmid, "id", 0, "vm id")
	flag.StringVar(&vmname, "name", "empty", "vm name")

	if len(os.Args) > 2 {
		flag.Parse()
	} else { // usage if no args
		usage()
		os.Exit(1)
	}

	// vmid/vmname
	// TODO: add more logic to verify
	if (vmid > 100 && vmid < 999999999) || vmname != "empty" {
	} else {
		fmt.Printf("ERROR: id is %d, and vm name is \"%s\", id must be between 100 - 999,999,999", vmid, vmname)
		os.Exit(1)
	}

	// proxrm
	proxrmx = new(proxrm.ProxRm)

	// Run
	err = proxrmx.Run(vmid, vmname)
	if err != nil {
		fmt.Printf("ERROR: Run: %s\n", err)
		os.Exit(1)
	}

	// Exit
	os.Exit(0)
}
