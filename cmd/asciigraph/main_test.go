package main

import (
	"bytes"
	"flag"
	"math"
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
	flag.CommandLine.Usage = func() {
		if flag.Usage != nil {
			flag.Usage()
		}
	}
	main()
	os.Exit(0)
}

func cliCommand(args ...string) *exec.Cmd {
	commandArgs := append([]string{"-test.run=^TestCLIProcess$", "--"}, args...)
	cmd := exec.Command(os.Args[0], commandArgs...)
	cmd.Env = append(os.Environ(), cliTestProcess+"=1")
	return cmd
}

func runCLIWithExitCode(t *testing.T, input string, args ...string) (string, string, int) {
	t.Helper()
	cmd := cliCommand(args...)
	cmd.Stdin = strings.NewReader(input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			t.Fatalf("failed to execute cli command: %v", err)
		}
	}
	return stdout.String(), stderr.String(), exitCode
}

func runCLI(t *testing.T, input string, args ...string) (string, string) {
	t.Helper()
	stdout, stderr, exitCode := runCLIWithExitCode(t, input, args...)
	if exitCode != 0 {
		t.Fatalf("command failed with exit code %d\nstderr:\n%s", exitCode, stderr)
	}
	return stdout, stderr
}

func resetGlobals() {
	height = 0
	width = 0
	offset = 3
	precision = 2
	caption = ""
	enableRealTime = false
	realTimeDataBuffer = 0
	fps = 24
	seriesColors = nil
	gradientColors = nil
	colorAbove = nil
	colorBelow = nil
	seriesLegends = nil
	captionColor = 0
	axisColor = 0
	labelColor = 0
	lowerBound = math.Inf(1)
	upperBound = math.Inf(-1)
	delimiter = ","
	seriesNum = 1
	customChar = ""
	xAxisMin = math.NaN()
	xAxisMax = math.NaN()
	xAxisTicks = 0
	lastGraphLines = 0
}

func runMainInProcess(t *testing.T, input string, args ...string) (string, string) {
	t.Helper()
	resetGlobals()

	oldArgs := os.Args
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	oldCommandLine := flag.CommandLine

	defer func() {
		os.Args = oldArgs
		os.Stdin = oldStdin
		os.Stdout = oldStdout
		os.Stderr = oldStderr
		flag.CommandLine = oldCommandLine
		resetGlobals()
	}()

	os.Args = append([]string{"asciigraph"}, args...)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	rIn, wIn, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdin pipe: %v", err)
	}
	os.Stdin = rIn
	go func() {
		_, _ = wIn.Write([]byte(input))
		_ = wIn.Close()
	}()

	rOut, wOut, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}
	os.Stdout = wOut

	rErr, wErr, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stderr pipe: %v", err)
	}
	os.Stderr = wErr

	main()

	_ = wOut.Close()
	_ = wErr.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	_, _ = stdoutBuf.ReadFrom(rOut)
	_, _ = stderrBuf.ReadFrom(rErr)

	return stdoutBuf.String(), stderrBuf.String()
}

func TestCLIStdinStdout(t *testing.T) {
	input := "1\n2\n3\n"
	stdout, stderr := runCLI(t, input)
	want := asciigraph.PlotMany([][]float64{{1, 2, 3}}) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
	if stderr != "" {
		t.Fatalf("unexpected stderr:\n%s", stderr)
	}
}

func TestCLIFileRedirection(t *testing.T) {
	dir := t.TempDir()
	inputPath := filepath.Join(dir, "input.txt")
	outputPath := filepath.Join(dir, "output.txt")
	if err := os.WriteFile(inputPath, []byte("3\n2\n1\n"), 0600); err != nil {
		t.Fatalf("failed to write input file: %v", err)
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		t.Fatalf("failed to open input file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		t.Fatalf("failed to create output file: %v", err)
	}

	cmd := cliCommand()
	cmd.Stdin = inputFile
	cmd.Stdout = outputFile
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		_ = outputFile.Close()
		t.Fatalf("command failed: %v\nstderr: %s", err, stderr.String())
	}
	if err := outputFile.Close(); err != nil {
		t.Fatalf("failed to close output file: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("unexpected stderr:\n%s", stderr.String())
	}

	got, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	want := asciigraph.PlotMany([][]float64{{3, 2, 1}}) + "\n"
	if string(got) != want {
		t.Fatalf("unexpected output:\ngot:\n%s\nwant:\n%s", string(got), want)
	}
}

func TestCLIDimensionsAndLayout(t *testing.T) {
	input := "1\n5\n2\n8\n"
	stdout, _ := runCLI(t, input, "-h", "5", "-w", "12", "-o", "4")
	want := asciigraph.PlotMany(
		[][]float64{{1, 5, 2, 8}},
		asciigraph.Height(5),
		asciigraph.Width(12),
		asciigraph.Offset(4),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
}

func TestCLICaption(t *testing.T) {
	input := "10\n20\n30\n"
	stdout, _ := runCLI(t, input, "-c", "CPU Utilization")
	want := asciigraph.PlotMany(
		[][]float64{{10, 20, 30}},
		asciigraph.Caption("CPU Utilization"),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
}

func TestCLIPrecision(t *testing.T) {
	tests := []struct {
		name      string
		precision string
		precUint  uint
	}{
		{name: "precision 0", precision: "0", precUint: 0},
		{name: "precision 4", precision: "4", precUint: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := "1.23456\n2.34567\n"
			stdout, _ := runCLI(t, input, "-p", tt.precision)
			want := asciigraph.PlotMany(
				[][]float64{{1.23456, 2.34567}},
				asciigraph.Precision(tt.precUint),
			) + "\n"
			if stdout != want {
				t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
			}
		})
	}
}

func TestCLIBounds(t *testing.T) {
	input := "10\n20\n30\n"
	stdout, _ := runCLI(t, input, "-lb", "0", "-ub", "50")
	want := asciigraph.PlotMany(
		[][]float64{{10, 20, 30}},
		asciigraph.LowerBound(0),
		asciigraph.UpperBound(50),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
}

func TestCLIMultiSeriesAndDelimiter(t *testing.T) {
	t.Run("custom semicolon delimiter", func(t *testing.T) {
		input := "1;10\n2;20\n3;30\n"
		stdout, _ := runCLI(t, input, "-sn", "2", "-d", ";")
		want := asciigraph.PlotMany(
			[][]float64{{1, 2, 3}, {10, 20, 30}},
		) + "\n"
		if stdout != want {
			t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
		}
	})

	t.Run("default comma delimiter", func(t *testing.T) {
		input := "5,15\n10,25\n15,35\n"
		stdout, _ := runCLI(t, input, "-sn", "2")
		want := asciigraph.PlotMany(
			[][]float64{{5, 10, 15}, {15, 25, 35}},
		) + "\n"
		if stdout != want {
			t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
		}
	})
}

func TestCLICustomCharacters(t *testing.T) {
	t.Run("single series custom char", func(t *testing.T) {
		input := "1\n2\n3\n"
		stdout, _ := runCLI(t, input, "-x", "*")
		want := asciigraph.PlotMany(
			[][]float64{{1, 2, 3}},
			asciigraph.SeriesChars(asciigraph.CreateCharSet("*")),
		) + "\n"
		if stdout != want {
			t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
		}
	})

	t.Run("multi series custom chars", func(t *testing.T) {
		input := "1,4\n2,5\n3,6\n"
		stdout, _ := runCLI(t, input, "-sn", "2", "-x", "*,#")
		want := asciigraph.PlotMany(
			[][]float64{{1, 2, 3}, {4, 5, 6}},
			asciigraph.SeriesChars(
				asciigraph.CreateCharSet("*"),
				asciigraph.CreateCharSet("#"),
			),
		) + "\n"
		if stdout != want {
			t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
		}
	})
}

func TestCLISeriesColorsAndLegends(t *testing.T) {
	input := "1,4\n2,5\n3,6\n"
	stdout, _ := runCLI(t, input, "-sn", "2", "-sc", "red,blue", "-sl", "Alpha,Beta")
	want := asciigraph.PlotMany(
		[][]float64{{1, 2, 3}, {4, 5, 6}},
		asciigraph.SeriesColors(asciigraph.Red, asciigraph.Blue),
		asciigraph.SeriesLegends("Alpha", "Beta"),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
}

func TestCLIGradientColors(t *testing.T) {
	t.Run("spectrum gradient", func(t *testing.T) {
		input := "1\n2\n3\n"
		stdout, _ := runCLI(t, input, "-g", "spectrum")
		want := asciigraph.PlotMany(
			[][]float64{{1, 2, 3}},
			asciigraph.SeriesColorGradient(asciigraph.HeatmapSpectrum...),
		) + "\n"
		if stdout != want {
			t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
		}
	})

	t.Run("custom gradient", func(t *testing.T) {
		input := "1\n2\n3\n"
		stdout, _ := runCLI(t, input, "-g", "blue,red")
		want := asciigraph.PlotMany(
			[][]float64{{1, 2, 3}},
			asciigraph.SeriesColorGradient(asciigraph.Blue, asciigraph.Red),
		) + "\n"
		if stdout != want {
			t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
		}
	})
}

func TestCLIColorThresholds(t *testing.T) {
	input := "1\n3\n7\n"
	stdout, _ := runCLI(t, input, "-ca", "red,5", "-cb", "blue,2")
	want := asciigraph.PlotMany(
		[][]float64{{1, 3, 7}},
		asciigraph.ColorAbove(asciigraph.Red, 5),
		asciigraph.ColorBelow(asciigraph.Blue, 2),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
}

func TestCLIPlotColors(t *testing.T) {
	input := "1\n2\n3\n"
	stdout, _ := runCLI(t, input, "-c", "Title", "-cc", "green", "-ac", "yellow", "-lc", "cyan")
	want := asciigraph.PlotMany(
		[][]float64{{1, 2, 3}},
		asciigraph.Caption("Title"),
		asciigraph.CaptionColor(asciigraph.Green),
		asciigraph.AxisColor(asciigraph.Yellow),
		asciigraph.LabelColor(asciigraph.Cyan),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
}

func TestCLIXAxisOptions(t *testing.T) {
	input := "1\n2\n3\n"
	stdout, _ := runCLI(t, input, "-xmin", "0", "-xmax", "10", "-xt", "5")
	want := asciigraph.PlotMany(
		[][]float64{{1, 2, 3}},
		asciigraph.XAxisRange(0, 10),
		asciigraph.XAxisTickCount(5),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
}

func TestCLIComprehensiveFlagCombination(t *testing.T) {
	input := "10|40\n20|50\n30|60\n"
	stdout, _ := runCLI(
		t,
		input,
		"-h", "4",
		"-w", "10",
		"-o", "5",
		"-p", "1",
		"-c", "Combined Test",
		"-sn", "2",
		"-d", "|",
		"-lb", "0",
		"-ub", "100",
		"-x", "^,v",
	)
	want := asciigraph.PlotMany(
		[][]float64{{10, 20, 30}, {40, 50, 60}},
		asciigraph.Height(4),
		asciigraph.Width(10),
		asciigraph.Offset(5),
		asciigraph.Precision(1),
		asciigraph.Caption("Combined Test"),
		asciigraph.LowerBound(0),
		asciigraph.UpperBound(100),
		asciigraph.SeriesChars(
			asciigraph.CreateCharSet("^"),
			asciigraph.CreateCharSet("v"),
		),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
}

func TestCLIRealTime(t *testing.T) {
	t.Run("basic realtime stream", func(t *testing.T) {
		input := "1\n2\n3\n"
		stdout, stderr, exitCode := runCLIWithExitCode(t, input, "-r", "-f", "100", "-b", "5")
		if exitCode != 0 {
			t.Fatalf("expected exit code 0, got %d, stderr: %s", exitCode, stderr)
		}
		if stdout == "" {
			t.Fatal("expected non-empty stdout in realtime mode")
		}
	})

	t.Run("realtime buffer truncation and options", func(t *testing.T) {
		input := "1,10\n2,20\n3,30\n4,40\n5,50\n"
		stdout, stderr, exitCode := runCLIWithExitCode(
			t,
			input,
			"-r",
			"-w", "3",
			"-f", "500",
			"-sn", "2",
			"-x", "*,#",
			"-ca", "red,25",
			"-cb", "blue,15",
			"-xmin", "0",
			"-xmax", "10",
			"-xt", "3",
			"-sc", "red,blue",
			"-sl", "S1,S2",
		)
		if exitCode != 0 {
			t.Fatalf("expected exit code 0, got %d, stderr: %s", exitCode, stderr)
		}
		if stdout == "" {
			t.Fatal("expected non-empty stdout in realtime mode with options")
		}
	})
}

func TestCLIInvalidInputHandling(t *testing.T) {
	t.Run("non-numeric input value logs error to stderr and continues", func(t *testing.T) {
		input := "1\ninvalid\n3\n"
		stdout, stderr, exitCode := runCLIWithExitCode(t, input)
		if exitCode != 0 {
			t.Fatalf("expected exit code 0, got %d; stderr: %s", exitCode, stderr)
		}
		if !strings.Contains(stderr, "cannot parse value") {
			t.Fatalf("expected stderr to report unparseable value, got: %s", stderr)
		}
		if stdout == "" {
			t.Fatal("expected graph output despite invalid data point")
		}
	})

	t.Run("fewer series points than specified causes fatal exit status 1", func(t *testing.T) {
		input := "1\n"
		_, stderr, exitCode := runCLIWithExitCode(t, input, "-sn", "2")
		if exitCode != 1 {
			t.Fatalf("expected exit code 1 on series count mismatch, got %d", exitCode)
		}
		if !strings.Contains(stderr, "number of series in the input data stream is less than the specified series number") {
			t.Fatalf("unexpected stderr output: %s", stderr)
		}
	})

	t.Run("more series points than specified are truncated", func(t *testing.T) {
		input := "1,2,3\n4,5,6\n"
		stdout, stderr, exitCode := runCLIWithExitCode(t, input, "-sn", "2")
		if exitCode != 0 {
			t.Fatalf("expected exit code 0 on truncated series, got %d; stderr: %s", exitCode, stderr)
		}
		want := asciigraph.PlotMany([][]float64{{1, 4}, {2, 5}}) + "\n"
		if stdout != want {
			t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
		}
	})

	t.Run("zero series specified causes fatal exit status 1", func(t *testing.T) {
		input := ""
		_, stderr, exitCode := runCLIWithExitCode(t, input, "-sn", "0")
		if exitCode != 1 {
			t.Fatalf("expected exit code 1 on zero series count, got %d", exitCode)
		}
		if !strings.Contains(stderr, "no data") {
			t.Fatalf("unexpected stderr output: %s", stderr)
		}
	})
}

func TestCLIExitStatusAndFlagErrors(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantExitCode int
		wantStderr   string
	}{
		{
			name:         "unknown flag",
			args:         []string{"-unsupported-flag"},
			wantExitCode: 2,
			wantStderr:   "flag provided but not defined",
		},
		{
			name:         "help flag",
			args:         []string{"-help"},
			wantExitCode: 0,
			wantStderr:   "Usage of",
		},
		{
			name:         "invalid series color",
			args:         []string{"-sc", "not-a-valid-color"},
			wantExitCode: 2,
			wantStderr:   "unrecognized color",
		},
		{
			name:         "invalid gradient color",
			args:         []string{"-g", "not-a-valid-color"},
			wantExitCode: 2,
			wantStderr:   "unrecognized color",
		},
		{
			name:         "invalid color above format",
			args:         []string{"-ca", "red"},
			wantExitCode: 2,
			wantStderr:   "expected `color,value`",
		},
		{
			name:         "invalid color above color name",
			args:         []string{"-ca", "invalidcolor,5"},
			wantExitCode: 2,
			wantStderr:   "unrecognized color",
		},
		{
			name:         "invalid color above threshold value",
			args:         []string{"-ca", "red,notanumber"},
			wantExitCode: 2,
			wantStderr:   "invalid threshold value",
		},
		{
			name:         "invalid color below format",
			args:         []string{"-cb", "invalidcolor,2"},
			wantExitCode: 2,
			wantStderr:   "unrecognized color",
		},
		{
			name:         "invalid caption color",
			args:         []string{"-cc", "invalidcolor"},
			wantExitCode: 2,
			wantStderr:   "unrecognized color",
		},
		{
			name:         "invalid axis color",
			args:         []string{"-ac", "invalidcolor"},
			wantExitCode: 2,
			wantStderr:   "unrecognized color",
		},
		{
			name:         "invalid label color",
			args:         []string{"-lc", "invalidcolor"},
			wantExitCode: 2,
			wantStderr:   "unrecognized color",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, stderr, exitCode := runCLIWithExitCode(t, "1\n2\n", tt.args...)
			if exitCode != tt.wantExitCode {
				t.Fatalf("expected exit code %d, got %d; stderr: %s", tt.wantExitCode, exitCode, stderr)
			}
			if !strings.Contains(stderr, tt.wantStderr) {
				t.Fatalf("expected stderr to contain %q; got: %s", tt.wantStderr, stderr)
			}
		})
	}
}

func TestParseColors(t *testing.T) {
	colors, ok := parseColors("red,green,blue")
	if !ok || len(colors) != 3 {
		t.Fatalf("expected 3 parsed colors, got %v, ok=%v", colors, ok)
	}
	if colors[0] != asciigraph.Red || colors[1] != asciigraph.Green || colors[2] != asciigraph.Blue {
		t.Fatalf("unexpected parsed colors: %v", colors)
	}

	colors, ok = parseColors(" red ")
	if !ok || len(colors) != 1 || colors[0] != asciigraph.Red {
		t.Fatalf("expected single red color, got %v, ok=%v", colors, ok)
	}

	_, ok = parseColors("notacolor")
	if ok {
		t.Fatal("expected failure on unrecognized color")
	}

	_, ok = parseColors("red,notacolor")
	if ok {
		t.Fatal("expected failure on list containing unrecognized color")
	}
}

func TestParseColorThreshold(t *testing.T) {
	color, threshold, err := parseColorThreshold("red,4.5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if color != asciigraph.Red || threshold != 4.5 {
		t.Fatalf("got color %v, threshold %f, want Red, 4.5", color, threshold)
	}

	_, _, err = parseColorThreshold("red")
	if err == nil {
		t.Fatal("expected error on missing threshold value")
	}

	_, _, err = parseColorThreshold("badcolor,4")
	if err == nil {
		t.Fatal("expected error on invalid color")
	}

	_, _, err = parseColorThreshold("red,invalid")
	if err == nil {
		t.Fatal("expected error on non-numeric threshold")
	}

	_, _, err = parseColorThreshold("red,NaN")
	if err == nil {
		t.Fatal("expected error on NaN threshold")
	}

	_, _, err = parseColorThreshold("red,+Inf")
	if err == nil {
		t.Fatal("expected error on Inf threshold")
	}
}

func TestParseColor(t *testing.T) {
	color, ok := parseColor("blue")
	if !ok || color != asciigraph.Blue {
		t.Fatalf("got %v, ok=%v, want Blue", color, ok)
	}

	_, ok = parseColor("nonexistentcolor")
	if ok {
		t.Fatal("expected false for nonexistent color")
	}
}

func TestClearPreviousGraph(t *testing.T) {
	lastGraphLines = 5
	clearPreviousGraph()
	lastGraphLines = 0
	clearPreviousGraph()
}

func TestMainDirectExecution(t *testing.T) {
	stdout, stderr := runMainInProcess(t, "10\n20\n30\n", "-h", "3", "-c", "Direct")
	want := asciigraph.PlotMany(
		[][]float64{{10, 20, 30}},
		asciigraph.Height(3),
		asciigraph.Caption("Direct"),
	) + "\n"
	if stdout != want {
		t.Fatalf("unexpected stdout:\ngot:\n%s\nwant:\n%s", stdout, want)
	}
	if stderr != "" {
		t.Fatalf("unexpected stderr: %s", stderr)
	}
}
