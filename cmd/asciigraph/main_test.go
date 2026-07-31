package main

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/guptarohit/asciigraph"
)

const cliTestProcess = "ASCIIGRAPH_CLI_TEST_PROCESS"

func TestCLIProcess(t *testing.T) {
	if os.Getenv(cliTestProcess) != "1" {
		return
	}

	separator := -1
	for i, arg := range os.Args {
		if arg == "--" {
			separator = i
			break
		}
	}
	if separator == -1 {
		os.Exit(2)
	}

	os.Args = append([]string{"asciigraph"}, os.Args[separator+1:]...)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
	os.Exit(0)
}

func TestCLIStdinStdout(t *testing.T) {
	got := runCLI(t, "1\n2\n3\n")
	want := asciigraph.PlotMany([][]float64{{1, 2, 3}}) + "\n"

	if got != want {
		t.Fatalf("unexpected output:\n%s\nwant:\n%s", got, want)
	}
}

func TestCLIFileInputOutput(t *testing.T) {
	dir := t.TempDir()
	inputPath := filepath.Join(dir, "input.txt")
	outputPath := filepath.Join(dir, "output.txt")
	if err := os.WriteFile(inputPath, []byte("3\n2\n1\n"), 0600); err != nil {
		t.Fatal(err)
	}

	input, err := os.Open(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	output, err := os.Create(outputPath)
	if err != nil {
		t.Fatal(err)
	}

	cmd := cliCommand()
	cmd.Stdin = input
	cmd.Stdout = output
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		output.Close()
		t.Fatalf("command failed: %v\n%s", err, stderr.String())
	}
	if err := output.Close(); err != nil {
		t.Fatal(err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("unexpected stderr:\n%s", stderr.String())
	}

	got, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	want := asciigraph.PlotMany([][]float64{{3, 2, 1}}) + "\n"
	if string(got) != want {
		t.Fatalf("unexpected output:\n%s\nwant:\n%s", got, want)
	}
}

func TestCLIFlags(t *testing.T) {
	got := runCLI(
		t,
		"1;3\n2;2\n3;1\n",
		"-d", ";",
		"-sn", "2",
		"-h", "2",
		"-w", "4",
		"-o", "6",
		"-p", "0",
		"-c", "two series",
		"-x", "*,#",
	)
	want := asciigraph.PlotMany(
		[][]float64{{1, 2, 3}, {3, 2, 1}},
		asciigraph.Height(2),
		asciigraph.Width(4),
		asciigraph.Offset(6),
		asciigraph.Precision(0),
		asciigraph.Caption("two series"),
		asciigraph.SeriesChars(
			asciigraph.CreateCharSet("*"),
			asciigraph.CreateCharSet("#"),
		),
	) + "\n"

	if got != want {
		t.Fatalf("unexpected output:\n%s\nwant:\n%s", got, want)
	}
}

func runCLI(t *testing.T, input string, args ...string) string {
	t.Helper()
	cmd := cliCommand(args...)
	cmd.Stdin = strings.NewReader(input)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("command failed: %v\n%s", err, stderr.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("unexpected stderr:\n%s", stderr.String())
	}
	return stdout.String()
}

func cliCommand(args ...string) *exec.Cmd {
	commandArgs := append([]string{"-test.run=^TestCLIProcess$", "--"}, args...)
	cmd := exec.Command(os.Args[0], commandArgs...)
	cmd.Env = append(os.Environ(), cliTestProcess+"=1")
	return cmd
}
