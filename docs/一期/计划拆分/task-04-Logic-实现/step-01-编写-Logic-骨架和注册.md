# Task 4 / Step 1: 编写 Logic 骨架和注册

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [Logic 实现](README.md)

- [ ] **Step 1: 编写 Logic 骨架和注册**

写入 `internal/logic/asset/asset.go`：

```go
package asset

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"asset_management_svr/internal/dao"
	"asset_management_svr/internal/model"
	"asset_management_svr/internal/service"
)

type sAsset struct{}

func init() {
	service.RegisterAsset(New())
}

func New() service.IAsset {
	return &sAsset{}
}

func normalizePage(page model.PageInput) model.PageInput {
	return page
}
```

当前项目 `go.mod` 的 module 为 `asset_management_svr`，示例导入路径已按该 module 编写。
