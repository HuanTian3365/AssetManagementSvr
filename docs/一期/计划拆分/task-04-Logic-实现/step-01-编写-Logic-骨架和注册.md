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

	"your_project/internal/dao"
	"your_project/internal/model"
	"your_project/internal/service"
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

将 `your_project/internal/...` 替换为项目实际 module 路径对应的导入路径。
