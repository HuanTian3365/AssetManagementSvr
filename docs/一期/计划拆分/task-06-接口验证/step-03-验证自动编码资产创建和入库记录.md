# Task 6 / Step 3: 验证自动编码资产创建和入库记录

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [接口验证](README.md)

- [ ] **Step 3: 验证自动编码资产创建和入库记录**

Run:

```bash
curl -X POST http://127.0.0.1:8000/asset/item/create -H "Content-Type: application/json" -d "{\"productId\":1,\"name\":\"练习用笔记本\",\"status\":1,\"roomId\":1,\"purchasePrice\":0,\"remark\":\"一期验证\"}"
curl "http://127.0.0.1:8000/asset/item/location-record/list?assetId=1&page=1&pageSize=20"
```

Expected:

```text
资产创建成功，返回 assetCode；位置记录列表包含 actionType=0 的入库记录。
```
