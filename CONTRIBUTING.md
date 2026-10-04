# Contributing

Thanks for your interest in Decision CLI.

## Development setup

```bash
asdf install            # installs Go 1.22.0 (see .tool-versions)
go run ./cmd/decision breakfast
go test -race ./...
```

Requires Go 1.22+. The only third-party dependency is `gopkg.in/yaml.v3`.

## Branching

- `main` — stable, released code.
- `dev` — integration branch. **All new work lands here first.**
- `feat/*`, `fix/*`, `chore/*` — short-lived branches, merged into `dev` via pull request.

`dev` is merged into `main` for a release.

## Commits

Follow [Conventional Commits](https://www.conventionalcommits.org/):
`type(scope): subject`.

Common types: `feat`, `fix`, `refactor`, `test`, `docs`, `chore`, `ci`,
`perf`, `build`.

Keep each commit small and focused on one change.

## Pull requests

- Open pull requests against **`dev`**, not `main`.
- CI must pass: formatting (`gofmt`), `go vet`, `go test -race`, `go build`.
- Keep the description short — what changed and why.

---

# 参与贡献

感谢你对 Decision CLI 的兴趣。

## 开发环境

```bash
asdf install            # 按 .tool-versions 安装 Go 1.22.0
go run ./cmd/decision breakfast
go test -race ./...
```

需要 Go 1.22+。唯一的第三方依赖是 `gopkg.in/yaml.v3`。

## 分支

- `main` —— 稳定、已发布的代码。
- `dev` —— 集成分支,**所有新工作先合入这里**。
- `feat/*`、`fix/*`、`chore/*` —— 短生命周期分支,通过 PR 合入 `dev`。

发布时再把 `dev` 合入 `main`。

## 提交

遵循 [Conventional Commits](https://www.conventionalcommits.org/):
`type(scope): subject`。

常用类型:`feat`、`fix`、`refactor`、`test`、`docs`、`chore`、`ci`、
`perf`、`build`。

每个提交尽量小而聚焦,一次只做一件事。

## 拉取请求

- PR 目标分支为 **`dev`**,不是 `main`。
- CI 必须通过:格式检查(`gofmt`)、`go vet`、`go test -race`、`go build`。
- 描述保持简短 —— 改了什么、为什么。
