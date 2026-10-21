# c1wubi — C1-Slim 五笔打字练习

看汉字输 86 版五笔编码，即时判分。字库内置一级简码 + 高频字（约 150）。

## 操作

| 键 | 作用 |
|---|---|
| 字母 | 输入编码 |
| Backspace | 删一字 |
| OK | 提交；答错后看答案，再 OK 下一题 |
| ← | 跳过 |
| N | 本局重开 |
| L | 切换难度 1/2/3 |
| Back | 退出 |

## 判分

- 简码或全码都算对（如「你」= `wq` 或 `wqiy`）
- 答错显示 `答案 简码 / 全码`
- 顶部显示 对/错/连胜

## 数据

`/storage/c1/local-apps/apps/wubi/data/stats.json` — 累计与最长连胜。

## 构建

```bat
set GOOS=linux& set GOARCH=mipsle& set GOMIPS=hardfloat& set CGO_ENABLED=0
go build -trimpath -ldflags "-s -w -buildid=" -o out\c1wubi .\cmd\c1wubi
```

不改 ELF e_flags。字库在 `wubi/table.go`，可继续加字。
