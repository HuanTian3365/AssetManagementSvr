package asset

import "asset_management_svr/internal/service"

type sAsset struct{}

func New() *sAsset {
	return &sAsset{}
}

func init() {
	service.RegisterAsset(New())

}
