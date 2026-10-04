# Decision CLI

A tiny terminal tool that makes the small, low-stakes choices for you — so you
can save your attention for the ones that matter.

一个轻量的终端决策工具:把不值得消耗注意力的小选择,交给一次命令。

```console
$ decision breakfast
早餐:包子

$ decision choose 麦当劳 肯德基 沙县
选择:肯德基
```

## Install

```bash
go install github.com/Carlos-CJC/decision@latest
```

Or build from source:

```bash
git clone https://github.com/Carlos-CJC/decision.git
cd decision
go build -o decision
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

Copy the template [`config/example.yaml`](config/example.yaml) and edit it. A set is a YAML file, e.g. `breakfast.yaml`:

```yaml
name: 早餐

options:
  - 面包
  - 包子
  - 麦片
  - 鸡蛋
```

Then run it by its file name (without the extension):

```bash
decision breakfast   # -> 早餐:包子
```

The command name **is** the file name. To add a command, add a file.
`study.yaml` gives you `decision study`. Non-ASCII names work too.

### Config lookup

For `decision <name>`, the file `<name>.yaml` is searched in this order:

1. the directory passed to `--config-dir`
2. `$DECISION_CONFIG_DIR`
3. `~/.config/decision/`
4. `./config/`

An explicit path such as `decision ./my.yaml` is loaded directly.

## License

[MIT](LICENSE)

---

# Decision CLI(中文)

一个轻量的终端决策工具:当你面对低价值、重复性的选择时,让终端用一条命令
给出结果,把注意力留给真正重要的事。

```console
$ decision breakfast
早餐:包子

$ decision choose 麦当劳 肯德基 沙县
选择:肯德基
```

## 安装

```bash
go install github.com/Carlos-CJC/decision@latest
```

或从源码编译:

```bash
git clone https://github.com/Carlos-CJC/decision.git
cd decision
go build -o decision
```

## 用法

| 命令 | 说明 |
| --- | --- |
| `decision <集合名>` | 运行一个已保存的决策集合 |
| `decision choose <选项> [<选项>...]` | 临时选择,无需配置文件 |
| `decision --config-dir <path> <集合名>` | 使用自定义配置目录 |
| `decision --help` | 显示帮助 |
| `decision --version` | 显示版本 |

### 已保存的集合

复制模板 [`config/example.yaml`](config/example.yaml) 后修改即可。一个集合就是一个
YAML 文件,例如 `breakfast.yaml`:

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

命令名**就是**文件名。想加一条命令,加一个文件即可:`study.yaml` 对应
`decision study`,中文文件名同样支持。

### 配置查找顺序

执行 `decision <集合名>` 时,按以下顺序查找 `<集合名>.yaml`:

1. `--config-dir` 指定的目录
2. `$DECISION_CONFIG_DIR`
3. `~/.config/decision/`
4. `./config/`

也可直接传入路径,如 `decision ./my.yaml`。

## 许可证

[MIT](LICENSE)
