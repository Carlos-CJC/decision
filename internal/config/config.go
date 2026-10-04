// Package config 负责决策集合配置的定位、加载与枚举。
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

var exts = []string{".yaml", ".yml"}

// Config 是一个决策集合的配置。
type Config struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Options []string `yaml:"options"`
}

// Entry 是一个可用的决策集合。
type Entry struct {
	Command string // 命令名(文件名去掉扩展名)
	Path    string // 配置文件路径
}

// Load 读取并校验一个配置文件。
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析 YAML 失败 (%s): %w", path, err)
	}
	if strings.TrimSpace(cfg.Name) == "" {
		return nil, fmt.Errorf("配置 %s 缺少必填字段 name", path)
	}
	if len(cfg.Options) == 0 {
		return nil, fmt.Errorf("配置 %s 的 options 不能为空", path)
	}
	return &cfg, nil
}

// SearchDirs 返回配置查找目录,按优先级从高到低:
// --config-dir → $DECISION_CONFIG_DIR → ~/.config/decision/ → ./config/。
func SearchDirs(configDir string) []string {
	var dirs []string
	if configDir != "" {
		dirs = append(dirs, configDir)
	}
	if env := os.Getenv("DECISION_CONFIG_DIR"); env != "" {
		dirs = append(dirs, env)
	}
	if home, err := os.UserHomeDir(); err == nil {
		dirs = append(dirs, filepath.Join(home, ".config", "decision"))
	}
	dirs = append(dirs, filepath.Join(".", "config"))
	return dirs
}

// Resolve 把集合名或显式路径解析为可读的配置文件路径。
func Resolve(name, configDir string) (string, error) {
	if looksLikePath(name) {
		if fi, err := os.Stat(name); err == nil && !fi.IsDir() {
			return name, nil
		}
		return "", fmt.Errorf("找不到配置文件: %s", name)
	}

	var tried []string
	for _, dir := range SearchDirs(configDir) {
		for _, ext := range exts {
			p := filepath.Join(dir, name+ext)
			if fi, err := os.Stat(p); err == nil && !fi.IsDir() {
				return p, nil
			}
			tried = append(tried, p)
		}
	}
	return "", fmt.Errorf("找不到集合 %q,已查找:\n  %s", name, strings.Join(tried, "\n  "))
}

// Available 扫描查找目录,返回可用的集合,按命令名去重(优先级高者先出现)
// 并按命令名排序。
func Available(configDir string) []Entry {
	seen := map[string]bool{}
	var entries []Entry

	for _, dir := range SearchDirs(configDir) {
		matches := []string{}
		for _, ext := range exts {
			found, _ := filepath.Glob(filepath.Join(dir, "*"+ext))
			matches = append(matches, found...)
		}
		for _, p := range matches {
			base := strings.TrimSuffix(filepath.Base(p), filepath.Ext(p))
			if seen[base] {
				continue
			}
			seen[base] = true
			entries = append(entries, Entry{Command: base, Path: p})
		}
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].Command < entries[j].Command })
	return entries
}

func looksLikePath(name string) bool {
	return strings.ContainsAny(name, `/\`) ||
		strings.HasSuffix(name, ".yaml") ||
		strings.HasSuffix(name, ".yml")
}
