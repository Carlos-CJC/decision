package cli

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// capture 运行 f 并捕获其写入 stdout 的内容。
func capture(t *testing.T, f func() int) (int, string) {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	os.Stdout = w
	code := f()
	w.Close()
	os.Stdout = old

	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	return code, string(out)
}

func TestRunHelp(t *testing.T) {
	for _, args := range [][]string{{"help"}, {"--help"}, {"-h"}} {
		code, out := capture(t, func() int { return Run(args) })
		if code != 0 {
			t.Errorf("Run(%v) code = %d, want 0", args, code)
		}
		if !strings.Contains(out, "用法") {
			t.Errorf("Run(%v) output missing usage, got %q", args, out)
		}
	}
}

func TestRunVersion(t *testing.T) {
	for _, args := range [][]string{{"version"}, {"--version"}} {
		code, out := capture(t, func() int { return Run(args) })
		if code != 0 {
			t.Errorf("Run(%v) code = %d, want 0", args, code)
		}
		if !strings.Contains(out, Version) {
			t.Errorf("Run(%v) output = %q, want it to contain %q", args, out, Version)
		}
	}
}

func TestRunChoose(t *testing.T) {
	code, out := capture(t, func() int { return Run([]string{"choose", "甲", "乙", "丙"}) })
	if code != 0 {
		t.Errorf("Run(choose) code = %d, want 0", code)
	}
	if !strings.HasPrefix(out, "选择:") {
		t.Errorf("Run(choose) output = %q, want prefix %q", out, "选择:")
	}
}

func TestRunChooseNoArgs(t *testing.T) {
	code, _ := capture(t, func() int { return Run([]string{"choose"}) })
	if code != 1 {
		t.Errorf("Run(choose) code = %d, want 1", code)
	}
}

func TestRunNoArgs(t *testing.T) {
	if code, _ := capture(t, func() int { return Run(nil) }); code != 1 {
		t.Errorf("Run(nil) code = %d, want 1", code)
	}
}

func TestRunList(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "breakfast.yaml"), []byte("name: 早餐\noptions:\n  - a\n"), 0o644)

	code, out := capture(t, func() int { return Run([]string{"--config-dir", dir, "list"}) })
	if code != 0 {
		t.Errorf("Run(list) code = %d, want 0", code)
	}
	if !strings.Contains(out, "breakfast") || !strings.Contains(out, "早餐") {
		t.Errorf("Run(list) output = %q, want it to contain breakfast and 早餐", out)
	}
}

func TestRunShow(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "breakfast.yaml"), []byte("name: 早餐\noptions:\n  - 面包\n"), 0o644)

	code, out := capture(t, func() int { return Run([]string{"--config-dir", dir, "show", "breakfast"}) })
	if code != 0 {
		t.Errorf("Run(show) code = %d, want 0", code)
	}
	if !strings.Contains(out, "早餐") || !strings.Contains(out, "面包") {
		t.Errorf("Run(show) output = %q, want it to contain 早餐 and 面包", out)
	}
}

func TestRunShowNoArg(t *testing.T) {
	if code, _ := capture(t, func() int { return Run([]string{"show"}) }); code != 1 {
		t.Errorf("Run(show) code = %d, want 1", code)
	}
}
