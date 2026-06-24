# Task 6 / Step 5: 验证删除引用保护

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [接口验证](README.md)

- [ ] **Step 5: 验证删除引用保护**

Run:

```bash
curl -X DELETE "http://127.0.0.1:8000/asset/room/delete?id=1"
```

Expected:

```text
返回业务错误：房间已被资产引用，不能删除。
```
