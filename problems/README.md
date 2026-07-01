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
- `cracking_the_coding_interview/`：程序员面试宝典。

迁移状态：

- `coding_interviews/` 已迁移全部非 `_undo.go` 题解；未完成题解以 `undo` build tag 保留。
- `code_carl/`、`hot100/`、`leetcode/`、`sword_offer/`、`cracking_the_coding_interview/` 已完成历史镜像迁移。
- 当前新结构中已有 438 个默认稳定实现，81 个历史原文仍以 `legacy` build tag 保留。
- 剩余 `legacy` 中只有 12 个目录暂无默认稳定实现，主要是旧测试拆分、设计模式拆分原文或依赖平台函数的题；后续可逐题整理后纳入稳定测试。
