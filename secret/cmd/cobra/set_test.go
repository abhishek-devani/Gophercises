package cobra

import "testing"

func TestSet(t *testing.T) {
	a := []string{"twitter_api_key", "nothing"}
	setCmd.Run(setCmd, a)
	MockSet = true
	setCmd.Run(setCmd, a)
}
