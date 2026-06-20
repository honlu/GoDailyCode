# Coding Interviews
本目录用于逐步迁移 `coding-interviews/` 中已经稳定的题解。

参考：https://leetcode.cn/studyplan/coding-interviews/

目录规则：

```text
lcr_159/
  solution.go
  solution_test.go
```

命名规则：

- package 使用题号，例如 `package lcr159`。
- 函数名可使用题目原始名称，不需要靠题号后缀避免冲突。
- 原 `coding-interviews/` 暂时保留为学习历史和迁移来源。

验证：

```bash
go test ./problems/coding_interviews/...
```

当前已迁移非 `_undo.go` 题解：

- `lcr_124`
- `lcr_126`
- `lcr_127`
- `lcr_129`
- `lcr_130`
- `lcr_133`
- `lcr_134`
- `lcr_135`
- `lcr_137`
- `lcr_139`
- `lcr_143`
- `lcr_144`
- `lcr_145`
- `lcr_150`
- `lcr_151`
- `lcr_152`
- `lcr_153`
- `lcr_155`
- `lcr_156`
- `lcr_157`
- `lcr_158`
- `lcr_159`
- `lcr_161`
- `lcr_163`
- `lcr_164`
- `lcr_165`
- `lcr_166`
- `lcr_167`
- `lcr_174`
- `lcr_175`
- `lcr_176`
- `lcr_177`
- `lcr_178`
- `lcr_180`
- `lcr_186`
- `lcr_188`
- `lcr_189`
- `lcr_190`
- `lcr_191`
- `lcr_194`
