// +build !windows

package checkadmin

import "os"

// environment for unixy systems

func checkAdmin() bool {
	return os.Getenv("USER") == "root"
}
