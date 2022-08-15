package controller_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"process-monitoring/controller"
	"testing"
)

func TestProcessIsRunning(t *testing.T) {
	res, err := controller.ProcessIsRunning("aaa")
	require.NoError(t, err)
	spew.Dump(res)
}
