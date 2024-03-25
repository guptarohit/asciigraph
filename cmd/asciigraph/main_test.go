package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCLI(t *testing.T) {
	// Capture the output of os.Stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Provide sample data as input
	sampleData := "1 2 3 4 5 6 7 8 9 10"
	stdin := os.Stdin
	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdin = pipeReader
	pipeWriter.Write([]byte(sampleData))
	pipeWriter.Close()

	// Call the main function
	main()

	// Restore the original os.Stdout and os.Stdin
	w.Close()
	os.Stdout = oldStdout
	os.Stdin = stdin

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Define the expected output
	expectedOutput := ` 10.00 ┤        ╭
  9.00 ┤       ╭╯
  8.00 ┤      ╭╯
  7.00 ┤     ╭╯
  6.00 ┤    ╭╯
  5.00 ┤   ╭╯
  4.00 ┤  ╭╯
  3.00 ┤ ╭╯
  2.00 ┤╭╯
  1.00 ┼╯
`

	// Compare the captured output with the expected output
	if output != expectedOutput {
		t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", output, expectedOutput)
	}
}

func TestCLIWithArgs(t *testing.T) {
	// Capture the output of os.Stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Provide sample data as input
	sampleData := "1 2 3 4 5 6 7 8 9 10"
	stdin := os.Stdin
	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdin = pipeReader
	pipeWriter.Write([]byte(sampleData))
	pipeWriter.Close()

	// Temporarily replace os.Args with the test arguments
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = append([]string{os.Args[0]}, []string{"-h", "5", "-w", "10"}...)

	// Call the main function
	main()

	// Restore the original os.Stdout and os.Stdin
	w.Close()
	os.Stdout = oldStdout
	os.Stdin = stdin

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Define the expected output
	expectedOutput := ` 10.00 ┤        ╭
  8.20 ┤       ╭╯
  6.40 ┤     ╭─╯
  4.60 ┤   ╭─╯
  2.80 ┤ ╭─╯
  1.00 ┼─╯
`

	// Compare the captured output with the expected output
	if output != expectedOutput {
		t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", output, expectedOutput)
	}
}
