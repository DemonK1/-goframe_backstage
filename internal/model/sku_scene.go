package model

type SkuSceneReq struct{}
type SkuSceneRes struct {
	Id   int    `json:"id"   `
	Name string `json:"name" ` // 分类名称
}
