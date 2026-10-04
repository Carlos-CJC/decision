// Package cli 负责命令行参数解析、命令分发与结果渲染。
package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/Carlos-CJC/decision/internal/chooser"
	"github.com/Carlos-CJC/decision/internal/config"
)

// Version 是 CLI 版本号。
const Version = "0.1.0"

const usage = `Decision CLI —— 把不值得消耗注意力的小选择,交给一次命令。

用法:
  decision <集合名>                   运行一个已保存的决策集合
  decision choose <选项> [<选项>...]  临时选择,无需配置文件
  decision list                       列出所有可用集合
  decision show <集合名>              查看某个集合的详细配置
  decision help                       显示帮助(同 --help)
  decision version                    显示版本(同 --version)

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

// Run 解析参数并执行对应命令,返回进程退出码。
func Run(args []string) int {
	var opt options
	var rest []string

	for i := 0; i < len(args); i++ {
		switch a := args[i]; {
		case a == "--help" || a == "-h":
			fmt.Fprint(os.Stdout, usage)
			return 0
		case a == "--version":
			fmt.Fprintln(os.Stdout, "decision "+Version)
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
	switch rest[0] {
	case "choose":
		return runChoose(rest[1:])
	case "list":
		return runList(rest[1:], opt)
	case "show":
		return runShow(rest[1:], opt)
	case "help":
		fmt.Fprint(os.Stdout, usage)
		return 0
	case "version":
		fmt.Fprintln(os.Stdout, "decision "+Version)
		return 0
	default:
		return runSet(rest[0], rest[1:], opt)
	}
}

func runSet(name string, extra []string, opt options) int {
	if len(extra) > 0 {
		return fail(fmt.Sprintf("%q 不接受额外参数: %s", name, strings.Join(extra, " ")))
	}

	path, err := config.Resolve(name, opt.configDir)
	if err != nil {
		return fail(err.Error())
	}
	cfg, err := config.Load(path)
	if err != nil {
		return fail(err.Error())
	}
	result, err := chooser.Pick(cfg.Options)
	if err != nil {
		return fail(err.Error())
	}

	fmt.Fprintln(os.Stdout, render(cfg.Name, result))
	return 0
}

func runChoose(args []string) int {
	if len(args) == 0 {
		return fail("用法: decision choose <选项> [<选项>...]")
	}
	result, err := chooser.Pick(args)
	if err != nil {
		return fail(err.Error())
	}
	fmt.Fprintln(os.Stdout, render("选择", result))
	return 0
}

// render 按默认模板 {name}:{result} 生成单行输出。
func render(name, result string) string {
	return fmt.Sprintf("%s:%s", name, result)
}

func fail(msg string) int {
	fmt.Fprintln(os.Stderr, "decision: "+msg)
	return 1
}
