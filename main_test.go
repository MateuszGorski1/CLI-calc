package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	clihandler "gorski.mateusz/calc/cli"
)

// Stopping logs from logging while testing
func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	os.Exit(m.Run())
}
func TestRunApp(t *testing.T) {

	testCases := map[string]struct {
		args   []string
		output string
	}{
		"test add": {
			args:   []string{"calc", "-op", "add", "3", "4"},
			output: "7\n",
		},
		"test sub": {
			args:   []string{"calc", "-op", "sub", "3", "4"},
			output: "-1\n",
		},
		"test mul": {
			args:   []string{"calc", "-op", "mul", "3", "4"},
			output: "12\n",
		},
		"test div": {
			args:   []string{"calc", "-op", "div", "6", "4"},
			output: "1.5\n",
		},
	}

	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			os.Args = test.args
			clihandler.Out = bytes.NewBuffer(nil)
			main()

			if actual := clihandler.Out.(*bytes.Buffer).String(); actual != test.output {
				require.Equal(t, test.output, actual, "unexpected result")
			}
		})
	}
}
