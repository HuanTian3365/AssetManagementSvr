# Task 6 / Step 4: 验证转移失败场景

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [接口验证](README.md)

- [ ] **Step 4: 验证转移失败场景**

Run:

```bash
curl -X POST http://127.0.0.1:8000/asset/item/transfer -H "Content-Type: application/json" -d "{\"assetId\":1,\"roomId\":1,\"remark\":\"相同房间验证\"}"
```

Expected:

```text
返回业务错误：目标房间不能等于当前房间。
```
