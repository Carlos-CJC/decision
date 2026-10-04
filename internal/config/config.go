// Package config 负责决策集合配置的定位与加载。
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config 是一个决策集合的配置。
type Config struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Options []string `yaml:"options"`
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

// Resolve 把集合名或显式路径解析为可读的配置文件路径。
// 查找顺序见 introduction.md 4.2:显式路径 → --config-dir → $DECISION_CONFIG_DIR
// → ~/.config/decision/ → ./config/。
func Resolve(name, configDir string) (string, error) {
	if looksLikePath(name) {
		if fi, err := os.Stat(name); err == nil && !fi.IsDir() {
			return name, nil
		}
		return "", fmt.Errorf("找不到配置文件: %s", name)
	}

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

	var tried []string
	for _, dir := range dirs {
		for _, ext := range []string{".yaml", ".yml"} {
			p := filepath.Join(dir, name+ext)
			if fi, err := os.Stat(p); err == nil && !fi.IsDir() {
				return p, nil
			}
			tried = append(tried, p)
		}
	}
	return "", fmt.Errorf("找不到集合 %q,已查找:\n  %s", name, strings.Join(tried, "\n  "))
}

func looksLikePath(name string) bool {
	return strings.ContainsAny(name, `/\`) ||
		strings.HasSuffix(name, ".yaml") ||
		strings.HasSuffix(name, ".yml")
}
