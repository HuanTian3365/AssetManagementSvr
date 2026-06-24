# Task 1 / Step 2: 导入 SQL

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [建表 SQL](README.md)

- [ ] **Step 2: 导入 SQL**

Run:

```bash
mysql -uroot -p your_database_name < manifest/sql/asset_init.sql
```

Expected:

```text
无 SQL 错误，7 张 asset_* 表创建成功。
```
