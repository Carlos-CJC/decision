package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `Decision CLI —— 把不值得消耗注意力的小选择,交给一次命令。

用法:
  decision <集合名>                   运行一个已保存的决策集合
  decision choose <选项> [<选项>...]  临时选择,无需配置文件
  decision --help                     显示帮助
  decision --version                  显示版本

参数:
  --config-dir <path>   覆盖配置目录

配置查找顺序(集合名 → 文件):
  1. --config-dir 指定目录
  2. $DECISION_CONFIG_DIR
  3. ~/.config/decision/
  4. ./config/

示例:
  decision breakfast
  decision choose 麦当劳 肯德基 沙县
`

type options struct {
	configDir string
}

func run(args []string) int {
	var opt options
	var rest []string

	for i := 0; i < len(args); i++ {
		switch a := args[i]; {
		case a == "--help" || a == "-h":
			fmt.Fprint(os.Stdout, usage)
			return 0
		case a == "--version":
			fmt.Fprintln(os.Stdout, "decision "+version)
			return 0
		case a == "--config-dir":
			if i+1 >= len(args) {
				return fail("--config-dir 需要一个路径参数")
			}
			i++
			opt.configDir = args[i]
		case strings.HasPrefix(a, "--config-dir="):
			opt.configDir = strings.TrimPrefix(a, "--config-dir=")
		default:
			rest = append(rest, a)
		}
	}

	if len(rest) == 0 {
		fmt.Fprint(os.Stderr, usage)
		return 1
	}

	if rest[0] == "choose" {
		return runChoose(rest[1:])
	}
	return runSet(rest[0], rest[1:], opt)
}

func runSet(name string, extra []string, opt options) int {
	if len(extra) > 0 {
		return fail(fmt.Sprintf("%q 不接受额外参数: %s", name, strings.Join(extra, " ")))
	}
	path, err := resolveConfig(name, opt.configDir)
	if err != nil {
		return fail(err.Error())
	}
	cfg, err := loadConfig(path)
	if err != nil {
		return fail(err.Error())
	}
	fmt.Fprintln(os.Stdout, cfg.Name)
	return 0
}

func runChoose(args []string) int {
	return fail("暂未实现")
}

func fail(msg string) int {
	fmt.Fprintln(os.Stderr, "decision: "+msg)
	return 1
}
