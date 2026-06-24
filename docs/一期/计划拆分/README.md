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

