package config

import (
	"os"
	"path/filepath"
	"testing"
)

func write(t *testing.T, dir, name, body string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
	return path
}

func TestLoad(t *testing.T) {
	path := write(t, t.TempDir(), "breakfast.yaml", "name: 早餐\noptions:\n  - 面包\n  - 包子\n")

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Name != "早餐" {
		t.Errorf("Name = %q, want %q", cfg.Name, "早餐")
	}
	if len(cfg.Options) != 2 {
		t.Errorf("Options = %v, want 2 items", cfg.Options)
	}
}

func TestLoadErrors(t *testing.T) {
	cases := map[string]string{
		"missing name":  "options:\n  - a\n",
		"empty options": "name: x\noptions: []\n",
		"bad yaml":      "name: [unclosed\n",
	}
	for name, body := range cases {
		t.Run(name, func(t *testing.T) {
			path := write(t, t.TempDir(), "c.yaml", body)
			if _, err := Load(path); err == nil {
				t.Errorf("Load(%s) = nil error, want error", name)
			}
		})
	}
}

func TestLoadMissingFile(t *testing.T) {
	if _, err := Load(filepath.Join(t.TempDir(), "nope.yaml")); err == nil {
		t.Error("Load(missing) = nil error, want error")
	}
}

func TestResolveByName(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "lunch.yaml", "name: 午餐\noptions:\n  - a\n")

	got, err := Resolve("lunch", dir)
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if want := filepath.Join(dir, "lunch.yaml"); got != want {
		t.Errorf("Resolve = %q, want %q", got, want)
	}
}

func TestResolveExplicitPath(t *testing.T) {
	path := write(t, t.TempDir(), "x.yaml", "name: x\noptions:\n  - a\n")

	got, err := Resolve(path, "")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if got != path {
		t.Errorf("Resolve = %q, want %q", got, path)
	}
}

func TestResolveMissing(t *testing.T) {
	if _, err := Resolve("definitely-not-a-real-set-xyz", t.TempDir()); err == nil {
		t.Error("Resolve(missing) = nil error, want error")
	}
}
