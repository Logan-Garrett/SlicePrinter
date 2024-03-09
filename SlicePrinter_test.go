package SlicePrinter_test

import (
	Print "SlicePrinter"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSlicePrinter(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = old
	}()

	dummyStringSlice := []string{"apple", "banana", "cherry", "date", "elderberry"}

	Print.SlicePrint(dummyStringSlice)

	w.Close()
	capturedOutput, _ := io.ReadAll(r)
	capturedOutputStr := strings.ReplaceAll(string(capturedOutput), "\r\n", "\n")

	expect_from_file_controller := "apple\nbanana\ncherry\ndate\nelderberry\n"
	expect_from_file_controller = strings.TrimSpace(expect_from_file_controller)
	if strings.TrimSpace(string(capturedOutputStr)) != expect_from_file_controller {
		t.Errorf("Expected: %q, Got: %q", expect_from_file_controller, string(capturedOutputStr))
	}
}

func TestSlice(t *testing.T) {
	TestSlicePrinter(t)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
