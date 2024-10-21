package main

import (
	"bytes"
	"os"
	"testing"
)

// Helper function to capture the output of the main function
func captureOutput(f func()) string {
	old := os.Stdout

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function f that writes to stdout
	f()

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output from the reader
	var buf bytes.Buffer
	buf.ReadFrom(r)

	return buf.String()
}

func TestMainWithVariousStrings(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "student$ go run . 'hello' standard",
			args: []string{"cmd", "hello", "standard"},
		},
		{
			name: "student$ go run . 'Hello There!' shadow.",
			args: []string{"cmd", "Hello There!", "shadow"},
		},
		{
			name: "student$ go run . 'Hello There!' thinkertoy",
			args: []string{"cmd", "Hello There!", "thinkertoy"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Temporarily set os.Args to the test case arguments
			os.Args = tt.args

			// Capture the output of main()
			output := captureOutput(main)

			// Print the output to see the result
			t.Logf("Output for %s:\n%s", tt.name, output)
		})
	}
}
