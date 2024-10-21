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
			name: "Try passing 'hello' as an argument.",
			args: []string{"cmd", "hello"},
		},
		{
			name: "Try passing 'HELLO' as an argument.",
			args: []string{"cmd", "HELLO"},
		},
		{
			name: "Try passing 'HeLlo HuMaN' as an argument.",
			args: []string{"cmd", "HeLlo HuMaN"},
		},
		{
			name: "Try passing '1Hello 2There' as an argument.",
			args: []string{"cmd", "1Hello 2There"},
		},
		{
			name: "Try passing 'Hello\\nThere' as an argument.",
			args: []string{"cmd", "Hello\\nThere"},
		},
		{
			name: "Try passing 'Hello\\n\\nThere' as an argument.",
			args: []string{"cmd", "Hello\\n\\nThere"},
		},
		{
			name: "Try passing '{Hello & There #}' as an argument.",
			args: []string{"cmd", "{Hello & There #}"},
		},
		{
			name: "Try passing 'hello There 1 to 2!' as an argument.",
			args: []string{"cmd", "hello There 1 to 2!"},
		},
		{
			name: "Try passing 'MaD3IrA&LiSboN' as an argument.",
			args: []string{"cmd", "MaD3IrA&LiSboN"},
		},
		{
			name: "Try passing '1a\"#FdwHywR&/()=' as an argument.",
			args: []string{"cmd", "1a\"#FdwHywR&/()="},
		},
		{
			name: "Try passing '{|}~' as an argument.",
			args: []string{"cmd", "{|}~"},
		},
		{
			name: "Try passing '[\\]^_ 'a' as an argument.",
			args: []string{"cmd", "[\\]^_ 'a"},
		},
		{
			name: "Try passing 'RGB' as an argument.",
			args: []string{"cmd", "RGB"},
		},
		{
			name: "Try passing ':;<=>?@' as an argument.",
			args: []string{"cmd", ":;<=>?@"},
		},
		{
			name: "Try passing <a random string> () with at least four lower case letters and three upper case letters.",
			args: []string{"cmd", "abcdABC"},
		},
		{
			name: "Try passing <a random string> with at least five lower case letters, a space and two numbers.",
			args: []string{"cmd", "abcde 12"},
		},
		{
			name: "Try passing <a random string> with at least one upper case letters and 3 special characters.",
			args: []string{"cmd", "A!@#"},
		},
		{
			name: "Try passing <a random string> with at least two lower case letters, two spaces, one number, two special characters and three upper case letters.",
			args: []string{"cmd", "ab  1!@ABC"},
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
