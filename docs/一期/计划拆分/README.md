# Personal Asset Management API Plan 拆分预览

> 原始计划: [`../2026-06-24-personal-asset-management-api-plan.md`](../计划.md)

## 原始计划前置内容

# Personal Asset Management API Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 在一个已创建并可运行的 GoFrame + MySQL 项目中，实现个人资产管理系统一期后端 API。

**Architecture:** 项目初始化、数据库连接和基础配置由用户自行完成。本计划从资产模块开始，使用 MySQL 建表、`gf gen dao` 生成 DAO/Entity/DO，再按 GoFrame 分层实现 API、Controller、Service、Logic，并用接口请求验证核心业务闭环。

**Tech Stack:** GoFrame v2.10.x、MySQL、Go、`gf` CLI、REST API。

---

## 前置约定

本计划不负责创建项目和配置数据库。执行前，用户需要已经完成：

- GoFrame 项目已创建。
- 项目可以执行 `go test ./...`。
- MySQL 数据库已创建。
- 项目数据库连接已配置。
- `gf -v` 可正常输出版本。

下文所有路径均相对于你的 GoFrame 项目根目录。

## 文件结构

- Create: `manifest/sql/asset_init.sql`  
  保存一期资产模块建表 SQL。
- Generated: `internal/dao/*`  
  由 `gf gen dao` 生成，不手工编辑。
- Generated: `internal/model/do/*`  
  由 `gf gen dao` 生成，不手工编辑。
- Generated: `internal/model/entity/*`  
  由 `gf gen dao` 生成，不手工编辑。
- Create: `api/asset/v1/asset.go`  
  定义一期资产模块全部请求和响应结构。
- Create: `internal/model/asset.go`  
  定义 Logic 层输入结构和常量。
- Create: `internal/service/asset.go`  
  定义资产模块业务接口。
- Create: `internal/logic/asset/asset.go`  
  实现 CRUD、引用保护、自动编码、入库、转移。
- Create: `internal/controller/asset/asset.go`  
  实现 HTTP Controller，调用 service。
- Modify: `internal/cmd/cmd.go`  
  注册资产模块控制器路由。

## 当前实际进度（2026-08-25）

> 本节记录代码仓库的实际状态，不改变下方原始 Task/Step 的计划定义。当前开发方式已经从按分层横向实施，调整为按“建筑 → 楼层 → 房间 → 分类”逐个资源纵向打通，因此不能仅根据 Task 编号判断完成度。

最后一次提交为 `04af881 feat: 房间增删改查`。提交之后仍有分类 CRUD 和房间校验相关的未提交改动，恢复开发前不得直接清理或重置工作区。

| 功能范围 | 当前状态 | 说明 |
| --- | --- | --- |
| 建表 SQL、DAO、DO、Entity | 基本完成 | 7 张表及生成文件已存在；实际 SQL 位于 `sql/`，与原计划的 `manifest/sql/asset_init.sql` 路径不同 |
| 建筑 CRUD | 已提交 | API、Controller、Service、Logic 和路由均已实现 |
| 楼层 CRUD | 已提交 | API、Controller、Service、Logic 和路由均已实现 |
| 房间 CRUD | 已提交，有未提交修正 | 未提交改动增加了楼层归属建筑校验；删除引用保护仍需在资产功能完成后验收 |
| 分类 CRUD | 进行中 | Create、Delete 已有实现；Update Controller 尚未实现；View、List Logic 仍为空实现 |
| 产品 CRUD | 未开始 | 当前只有生成的数据访问层 |
| 资产 CRUD、自动编码、入库、转移、位置记录 | 未开始 | 一期核心业务尚未实现 |
| 接口集成验证 | 未开始 | 2026-08-24 检查时本机 MySQL 3306 端口没有监听，未执行数据库写入验证 |

### 最近一次检查结果

- `go test ./...` 通过，但项目不存在 `_test.go`，结果只能证明全部包可以编译。
- `go vet ./...` 通过。
- `gofmt -d` 未发现格式差异。
- `git diff --check` 通过，仅提示部分文件后续可能从 LF 转换为 CRLF。
- 服务可以在 `:8000` 启动，OpenAPI 实际注册建筑、楼层、房间、分类共 20 条路由；检查后服务已停止。
- 分类的空命名返回值可以通过编译，因此编译通过不能视为分类 CRUD 已完成。

### 推荐恢复顺序

1. 先保护当前未提交改动，避免重置或覆盖分类半成品。
2. 以房间模块为样例，按 API → Controller → Service → Logic → DAO/Entity → 数据表的顺序恢复调用链知识。
3. 补齐并验收分类 CRUD，再形成独立提交。
4. 依次实现产品 CRUD、资产基础 CRUD、自动编码与入库、资产转移与位置记录。
5. MySQL 可用后执行 Task 6 的完整接口验证。

待确认的设计问题和技术债统一记录在 [`../TODO.md`](../TODO.md)。

## Task 目录

- [Task 1: 建表 SQL](task-01-建表-SQL/README.md)(4 个 Step)
- [Task 2: API 定义](task-02-API-定义/README.md)(4 个 Step)
- [Task 3: Service 和模型](task-03-Service-和模型/README.md)(4 个 Step)
- [Task 4: Logic 实现](task-04-Logic-实现/README.md)(7 个 Step)
- [Task 5: Controller 和路由注册](task-05-Controller-和路由注册/README.md)(4 个 Step)
- [Task 6: 接口验证](task-06-接口验证/README.md)(6 个 Step)

## Step 快速入口

### Task 1: 建表 SQL

- [Step 1: 创建资产模块建表 SQL](task-01-建表-SQL/step-01-创建资产模块建表-SQL.md)
- [Step 2: 导入 SQL](task-01-建表-SQL/step-02-导入-SQL.md)
- [Step 3: 生成 DAO](task-01-建表-SQL/step-03-生成-DAO.md)
- [Step 4: 提交](task-01-建表-SQL/step-04-提交.md)

### Task 2: API 定义

- [Step 1: 编写 API 请求响应结构](task-02-API-定义/step-01-编写-API-请求响应结构.md)
- [Step 2: 生成 Controller 模板](task-02-API-定义/step-02-生成-Controller-模板.md)
- [Step 3: 验证 API 定义可编译](task-02-API-定义/step-03-验证-API-定义可编译.md)
- [Step 4: 提交](task-02-API-定义/step-04-提交.md)

### Task 3: Service 和模型

- [Step 1: 编写内部模型](task-03-Service-和模型/step-01-编写内部模型.md)
- [Step 2: 编写 Service 接口](task-03-Service-和模型/step-02-编写-Service-接口.md)
- [Step 3: 验证编译失败点明确](task-03-Service-和模型/step-03-验证编译失败点明确.md)
- [Step 4: 提交](task-03-Service-和模型/step-04-提交.md)

### Task 4: Logic 实现

- [Step 1: 编写 Logic 骨架和注册](task-04-Logic-实现/step-01-编写-Logic-骨架和注册.md)
- [Step 2: 编写通用 CRUD 辅助方法](task-04-Logic-实现/step-02-编写通用-CRUD-辅助方法.md)
- [Step 3: 编写基础资料 CRUD](task-04-Logic-实现/step-03-编写基础资料-CRUD.md)
- [Step 4: 编写分类和产品 CRUD](task-04-Logic-实现/step-04-编写分类和产品-CRUD.md)
- [Step 5: 编写位置校验、自动编码、资产入库和转移](task-04-Logic-实现/step-05-编写位置校验、自动编码、资产入库和转移.md)
- [Step 6: 运行编译检查](task-04-Logic-实现/step-06-运行编译检查.md)
- [Step 7: 提交](task-04-Logic-实现/step-07-提交.md)

### Task 5: Controller 和路由注册

- [Step 1: 编写 Controller](task-05-Controller-和路由注册/step-01-编写-Controller.md)
- [Step 2: 注册路由](task-05-Controller-和路由注册/step-02-注册路由.md)
- [Step 3: 验证全项目编译](task-05-Controller-和路由注册/step-03-验证全项目编译.md)
- [Step 4: 提交](task-05-Controller-和路由注册/step-04-提交.md)

### Task 6: 接口验证

- [Step 1: 启动服务](task-06-接口验证/step-01-启动服务.md)
- [Step 2: 验证基础数据创建](task-06-接口验证/step-02-验证基础数据创建.md)
- [Step 3: 验证自动编码资产创建和入库记录](task-06-接口验证/step-03-验证自动编码资产创建和入库记录.md)
- [Step 4: 验证转移失败场景](task-06-接口验证/step-04-验证转移失败场景.md)
- [Step 5: 验证删除引用保护](task-06-接口验证/step-05-验证删除引用保护.md)
- [Step 6: 提交验证记录](task-06-接口验证/step-06-提交验证记录.md)

