package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"unicode/utf8"

	"github.com/Carlos-CJC/decision/internal/config"
)

// runList 列出所有可用集合。
func runList(args []string, opt options) int {
	if len(args) > 0 {
		return fail("用法: decision list")
	}

	entries := config.Available(opt.configDir)
	if len(entries) == 0 {
		fmt.Fprintln(os.Stderr, "没有找到任何集合。已查找:")
		for _, d := range config.SearchDirs(opt.configDir) {
			fmt.Fprintln(os.Stderr, "  "+d)
		}
		return 0
	}

	// 命令名(通常为 ASCII)定宽左对齐;中文名称放行尾,避免等宽错位。
	width := 0
	for _, e := range entries {
		if w := utf8.RuneCountInString(e.Command); w > width {
			width = w
		}
	}
	for _, e := range entries {
		name, count := "(无法加载)", "-"
		if cfg, err := config.Load(e.Path); err == nil {
			name = cfg.Name
			count = fmt.Sprintf("%d", len(cfg.Options))
		}
		fmt.Fprintf(os.Stdout, "%-*s  %s (%s)\n", width, e.Command, name, count)
	}
	return 0
}

// runShow 展示某个集合的详细配置。
func runShow(args []string, opt options) int {
	if len(args) != 1 {
		return fail("用法: decision show <集合名>")
	}

	path, err := config.Resolve(args[0], opt.configDir)
	if err != nil {
		return fail(err.Error())
	}
	cfg, err := config.Load(path)
	if err != nil {
		return fail(err.Error())
	}

	typ := cfg.Type
	if typ == "" {
		typ = "choice"
	}
	command := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "命令:\t%s\n", command)
	fmt.Fprintf(w, "名称:\t%s\n", cfg.Name)
	fmt.Fprintf(w, "类型:\t%s\n", typ)
	fmt.Fprintf(w, "路径:\t%s\n", path)
	w.Flush()

	fmt.Fprintf(os.Stdout, "\n选项 (%d):\n", len(cfg.Options))
	for _, o := range cfg.Options {
		fmt.Fprintf(os.Stdout, "  - %s\n", o)
	}
	return 0
}
