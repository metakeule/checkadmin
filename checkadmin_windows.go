// +build windows

package checkforadmin

import "os"

// based on https://www.reddit.com/r/golang/comments/53dthc/way_to_detect_if_the_programs_running_with/d7tbz28/
// https://play.golang.org/p/bBtRZrk4_p
func checkAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}
