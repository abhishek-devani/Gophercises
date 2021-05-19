package cobra

import "testing"

func TestGet(t *testing.T) {
	a := []string{"twitter_api_key"}
	getCmd.Run(getCmd, a)
	MockGet = true
	getCmd.Run(getCmd, a)
}
