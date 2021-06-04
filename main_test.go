package main

import "testing"

func Test_run(t *testing.T) {
	exitCode := run()
	if exitCode != exitOK {
		t.Errorf("exitCode wrong, got: %d, want: %d.", exitCode, exitCode)
	}
}
