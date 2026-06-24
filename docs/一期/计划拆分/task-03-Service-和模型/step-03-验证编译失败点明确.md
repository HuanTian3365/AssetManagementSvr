# Task 3 / Step 3: 验证编译失败点明确

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Service 和模型](README.md)

- [ ] **Step 3: 验证编译失败点明确**

Run:

```bash
go test ./internal/service ./internal/model
```

Expected:

```text
如果导入路径正确，service 和 model 编译通过。
```
