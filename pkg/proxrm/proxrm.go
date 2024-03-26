//
// proxrm
//
// proxrm.go
//
// Uses the proxmoxa-api-go library:
// https://pkg.go.dev/github.com/Telmate/proxmox-api-go
//

package proxrm

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

// ProxRm app struct for shared storage
type ProxRm struct {
	vmid   int
	client *proxmox.Client
	vmr    *proxmox.VmRef
}

// Run does all the coordinating
func (proxrm *ProxRm) Run() error {
	var err error

	// vmid
	// TODO: add more logic to verify
	vmid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return err
	} else if vmid <= 0 {
		return fmt.Errorf("ERROR: vmid is %d", vmid)
	}

	// Initialize
	err = proxrm.Initialize(vmid)
	if err != nil {
		fmt.Printf("ERROR: Initialization: %s\n", err)
		return err
	}

	// Ping vm
	err = proxrm.ping()
	if err != nil {
		fmt.Printf("ERROR: Ping: %s\n", err)
		return err
	}

	// Stop vm
	err = proxrm.stop()
	if err != nil {
		fmt.Printf("ERROR: Stop: %s\n", err)
		return err
	}

	// Delete
	err = proxrm.delete()
	if err != nil {
		fmt.Printf("ERROR: Delete: %s\n", err)
		return err
	}

	return err
}

// Initialize it all
func (proxrm *ProxRm) Initialize(vmid int) error {
	var err error

	// vmid
	proxrm.vmid = vmid

	// client
	err = proxrm.newClient()
	if err != nil {
		return err
	}

	// make sure client is not nil
	if proxrm.client == nil {
		return fmt.Errorf(" 🔐 client is nil")
	}

	// get vm reference
	vmr := proxmox.NewVmRef(vmid)

	// make sure vm reference is not nil
	if vmr == nil {
		return fmt.Errorf(" 🔐 vm reference is nil")
	}

	// vmref
	proxrm.vmr = vmr

	return err
}

// newClient establishes communication with Proxmox
func (proxrm *ProxRm) newClient() error {
	var err error

	clientConfig := new(ClientConfig)
	err = clientConfig.Initialize()
	if err != nil {
		return err
	}

	// use insecure client
	tlsConfig := tls.Config{
		InsecureSkipVerify: true, // #nosec G402
	}

	// new client
	proxrm.client, err = proxmox.NewClient(clientConfig.url, nil, "", &tlsConfig, "", clientConfig.timeout)
	if err != nil {
		return err
	}

	// authenticate with username/password or username/token
	if clientConfig.password != "" {
		err = proxrm.client.Login(clientConfig.username, clientConfig.password, "")
		if err != nil {
			return err
		}
	} else if clientConfig.token != "" {
		proxrm.client.SetAPIToken(clientConfig.username, clientConfig.token)
	}

	// make sure client is not empty
	if proxrm.client == nil {
		return fmt.Errorf(" 🔐 proxmox client is nil: %s", err)
	}

	fmt.Println(" 🔓 proxmox client obtained")

	return err
}

// Ping verifies communication with Proxmox server
func (proxrm *ProxRm) ping() error {
	// As test, get the version of the server
	_, err := proxrm.client.GetVersion()
	if err != nil {
		return err
	}

	fmt.Printf(" 🏓 pinged vmid: %d\n", proxrm.vmid)

	return err
}

// Stop the Proxmox vm
func (proxrm *ProxRm) stop() error {
	_, err := proxrm.client.StopVm(proxrm.vmr)
	if err != nil {
		return err
	}

	fmt.Printf(" 🛑 stopped vmid: %d\n", proxrm.vmid)

	return err
}

// Delete the Proxmox vm
func (proxrm *ProxRm) delete() error {
	_, err := proxrm.client.DeleteVm(proxrm.vmr)
	if err != nil {
		return err
	}

	fmt.Printf(" 🗑️ deleted vmid: %d\n", proxrm.vmid)

	return err
}
