//
// proxrm
//
// ClientConfig.go
//

package proxrm

import (
	"fmt"
	"os"
	"strconv"
)

// ClientConfig holds Proxmox credentials
type ClientConfig struct {
	url      string
	username string
	password string
	token    string
	timeout  int
}

// Initialize client configuration
func (cc *ClientConfig) Initialize() error {
	var err error

	// load proxmox url from environment variable
	cc.url = os.Getenv("PROXMOX_URL")

	// load credentials from environment variables
	cc.username = os.Getenv("PROXMOX_USERNAME")
	cc.password = os.Getenv("PROXMOX_PASSWORD")
	cc.token = os.Getenv("PROXMOX_TOKEN")

	// task timeout
	cc.timeout, err = strconv.Atoi(os.Getenv("PROXMOX_TIMEOUT"))
	if err != nil {
		fmt.Printf("ERROR: PROXMOX_TIMEOUT: %s\n", err)
		return err
	}

	// need url
	if cc.url == "" {
		return fmt.Errorf(" 💲 $PROXMOX_URL is empty")
	}

	// need username
	if cc.username == "" {
		return fmt.Errorf(" 💲 $PROXMOX_USERNAME is empty")
	}

	// need either password or token, if both are empty, error
	if cc.password == "" && cc.token == "" {
		return fmt.Errorf(" 💲 $PROXMOX_PASSWORD and PROXMOX_TOKEN are empty")
	}

	if cc.timeout == 0 {
		return fmt.Errorf(" 💲 $PROXMOX_TIMEOUT is zero")
	}

	return err
}
