package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_run(t *testing.T) {
	exitCode := run()
	assert.Equal(t, exitOK, exitCode, "should be zero")
}
