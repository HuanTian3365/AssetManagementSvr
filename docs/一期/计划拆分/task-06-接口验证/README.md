# Task 6: 接口验证

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`

## 概览

## Task 6: 接口验证

**Files:**
- No source changes

## Step 预览

- [Step 1: 启动服务](step-01-启动服务.md)
- [Step 2: 验证基础数据创建](step-02-验证基础数据创建.md)
- [Step 3: 验证自动编码资产创建和入库记录](step-03-验证自动编码资产创建和入库记录.md)
- [Step 4: 验证转移失败场景](step-04-验证转移失败场景.md)
- [Step 5: 验证删除引用保护](step-05-验证删除引用保护.md)
- [Step 6: 提交验证记录](step-06-提交验证记录.md)

## 后续说明

## 执行注意事项与自检结果

- 设计文档中的一期范围均已覆盖：位置、产品、资产、编码、入库、转移、位置记录、删除引用保护。
- 项目初始化和配置已按用户要求排除，由用户自行完成。
- 当前项目 `go.mod` 的 module 为 `asset_management_svr`，代码示例中的导入路径按该 module 编写。
- 计划没有保留未说明用途的占位项。
- GoFrame CLI 用法参考了官方 `gf gen dao`、`gf gen ctrl` 文档；本机已验证 `gf v2.10.2` 可用。
