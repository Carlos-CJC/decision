# Decision CLI

[English](README.md) | [简体中文](README.zh-CN.md)

[![CI](https://github.com/Carlos-CJC/decision/actions/workflows/ci.yml/badge.svg)](https://github.com/Carlos-CJC/decision/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Carlos-CJC/decision)](go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/Carlos-CJC/decision)](https://goreportcard.com/report/github.com/Carlos-CJC/decision)
[![License](https://img.shields.io/github/license/Carlos-CJC/decision)](LICENSE)

> A tiny terminal tool that makes the small, low-stakes choices for you, so you
> can save your attention for the ones that matter.

```console
$ decision breakfast
早餐:包子

$ decision choose 麦当劳 肯德基 沙县
选择:肯德基
```

## Features

- **Saved sets** — define your choices once in YAML, run them by name.
- **Ad-hoc picks** — `decision choose a b c`, no config file needed.
- **Simple lookup** — sets are found by file name, from a standard directory.
- **Zero setup** — one static binary, no database, no network, no account.
- **Equal probability** — every option has the same chance.

## Install

### go install

```bash
go install github.com/Carlos-CJC/decision/cmd/decision@latest
```

### From source

```bash
git clone https://github.com/Carlos-CJC/decision.git
cd decision
go build -o decision ./cmd/decision
```

## Usage

| Command | Description |
| --- | --- |
| `decision <name>` | Run a saved decision set |
| `decision choose <option> [<option>...]` | Pick one on the fly, no config needed |
| `decision --config-dir <path> <name>` | Use a custom config directory |
| `decision --help` | Show help |
| `decision --version` | Show version |

### Saved sets

Copy the template [`config/example.yaml`](config/example.yaml) and edit it:

```yaml
name: 早餐

options:
  - 面包
  - 包子
  - 麦片
  - 鸡蛋
```

Then run it by file name (without the extension):

```bash
decision breakfast   # -> 早餐:包子
```

The command name **is** the file name — to add a command, add a file.
`study.yaml` gives you `decision study`. Non-ASCII names work too.

### Config lookup

For `decision <name>`, the file `<name>.yaml` is searched in this order:

1. the directory passed to `--config-dir`
2. `$DECISION_CONFIG_DIR`
3. `~/.config/decision/`
4. `./config/`

An explicit path such as `decision ./my.yaml` is loaded directly.

## Project layout

```
├── cmd/decision/      # executable entry point
├── internal/
│   ├── cli/           # argument parsing, command dispatch, rendering
│   ├── config/        # config lookup and loading
│   └── chooser/       # equal-probability selection
└── config/            # example configuration
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE)
