# Task 5 / Step 2: 注册路由

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Controller 和路由注册](README.md)

- [ ] **Step 2: 注册路由**

在 `internal/cmd/cmd.go` 的服务启动路由注册位置增加资产控制器导入和绑定。示例：

```go
import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"your_project/internal/controller/asset"
)
```

在 `s.Group("/", func(group *ghttp.RouterGroup) { ... })` 或项目现有路由分组中加入：

```go
group.Bind(
	asset.NewV1(),
)
```

如果文件还没有 `ghttp` 导入，补充：

```go
import "github.com/gogf/gf/v2/net/ghttp"
```
