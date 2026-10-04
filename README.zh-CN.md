# Decision CLI

[English](README.md) | [简体中文](README.zh-CN.md)

[![CI](https://github.com/Carlos-CJC/decision/actions/workflows/ci.yml/badge.svg)](https://github.com/Carlos-CJC/decision/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Carlos-CJC/decision)](go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/Carlos-CJC/decision)](https://goreportcard.com/report/github.com/Carlos-CJC/decision)
[![License](https://img.shields.io/github/license/Carlos-CJC/decision)](LICENSE)

> 一个轻量的终端决策工具:把不值得消耗注意力的小选择,交给一次命令。

```console
$ decision breakfast
早餐:包子

$ decision choose 麦当劳 肯德基 沙县
选择:肯德基
```

## 特性

- **已保存的集合** —— 用 YAML 定义一次,之后按名字运行。
- **临时选择** —— `decision choose 甲 乙 丙`,无需配置文件。
- **简单定位** —— 集合以文件名定位,放在标准目录即可。
- **零依赖** —— 单个静态二进制,无需数据库、不联网、无需账号。
- **等概率** —— 每个选项机会相同。

## 安装

### go install

```bash
go install github.com/Carlos-CJC/decision/cmd/decision@latest
```

### 从源码编译

```bash
git clone https://github.com/Carlos-CJC/decision.git
cd decision
go build -o decision ./cmd/decision
```

## 用法

| 命令 | 说明 |
| --- | --- |
| `decision <集合名>` | 运行一个已保存的决策集合 |
| `decision choose <选项> [<选项>...]` | 临时选择,无需配置文件 |
| `decision --config-dir <path> <集合名>` | 使用自定义配置目录 |
| `decision help` | 显示帮助(同 `--help`、`-h`) |
| `decision version` | 显示版本(同 `--version`) |

### 已保存的集合

复制模板 [`config/example.yaml`](config/example.yaml) 后修改:

```yaml
name: 早餐

options:
  - 面包
  - 包子
  - 麦片
  - 鸡蛋
```

然后按文件名(去掉扩展名)运行:

```bash
decision breakfast   # -> 早餐:包子
```

命令名**就是**文件名 —— 想加一条命令,加一个文件即可:`study.yaml` 对应
`decision study`,中文文件名同样支持。

### 配置查找顺序

执行 `decision <集合名>` 时,按以下顺序查找 `<集合名>.yaml`:

1. `--config-dir` 指定的目录
2. `$DECISION_CONFIG_DIR`
3. `~/.config/decision/`
4. `./config/`

也可直接传入路径,如 `decision ./my.yaml`。

## 项目结构

```
├── cmd/decision/      # 可执行入口
├── internal/
│   ├── cli/           # 参数解析、命令分发、结果渲染
│   ├── config/        # 配置定位与加载
│   └── chooser/       # 等概率选择
└── config/            # 示例配置
```

## 参与贡献

见 [CONTRIBUTING.md](CONTRIBUTING.md)。

## 许可证

[MIT](LICENSE)
