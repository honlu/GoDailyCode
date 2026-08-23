# 题解稳定目录

这里放逐步整理后的题解。新迁移的题目应尽量做到：

- 每题一个独立 package。
- 每题目录包含 `solution.go` 和必要的 `solution_test.go`。
- 保留原实现思路和关键注释，只做必要的 package、命名和测试调整。

当前已迁移：

- `coding_interviews/`：Interview 75 / LCR 题单。
- `hot100/`：Hot100。
- `codetop/`：CodeTop 公司高频面试题。
- `leetcode/`：个人收藏题。
- `code_carl/`：代码随想录题解。
- `sword_offer/`：剑指 Offer。

迁移状态：

- `coding_interviews/` 已迁移全部非 `_undo.go` 题解；未完成题解以 `undo` build tag 保留。
- `code_carl/`、`hot100/`、`leetcode/`、`sword_offer/`、`cracking_the_coding_interview/` 已完成历史镜像迁移。
- 默认实现参与 `go test ./problems/...`；尚未完成的题解和历史原文通过 build tag 保留。
- 目前仍有较多 package 仅通过编译检查、没有独立测试，后续可随复习进度逐题补充。
