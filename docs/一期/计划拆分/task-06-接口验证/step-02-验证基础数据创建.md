# Task 6 / Step 2: 验证基础数据创建

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [接口验证](README.md)

- [ ] **Step 2: 验证基础数据创建**

Run:

```bash
curl -X POST http://127.0.0.1:8000/asset/building/create -H "Content-Type: application/json" -d "{\"name\":\"家\",\"code\":\"HOME\",\"address\":\"本地\",\"remark\":\"\"}"
curl -X POST http://127.0.0.1:8000/asset/floor/create -H "Content-Type: application/json" -d "{\"buildingId\":1,\"name\":\"一楼\",\"code\":\"HOME-F1\",\"floorNo\":1,\"remark\":\"\"}"
curl -X POST http://127.0.0.1:8000/asset/room/create -H "Content-Type: application/json" -d "{\"buildingId\":1,\"floorId\":1,\"name\":\"书房\",\"code\":\"HOME-F1-STUDY\",\"roomNo\":\"101\",\"remark\":\"\"}"
curl -X POST http://127.0.0.1:8000/asset/category/create -H "Content-Type: application/json" -d "{\"parentId\":0,\"name\":\"电脑设备\",\"code\":\"COMPUTER\",\"sort\":1,\"remark\":\"\"}"
curl -X POST http://127.0.0.1:8000/asset/product/create -H "Content-Type: application/json" -d "{\"categoryId\":1,\"name\":\"笔记本电脑\",\"code\":\"NOTEBOOK\",\"brand\":\"\",\"model\":\"\",\"unit\":\"台\",\"remark\":\"\"}"
```

Expected:

```text
每个请求返回成功结果，并包含新增 ID 或成功状态。
```
