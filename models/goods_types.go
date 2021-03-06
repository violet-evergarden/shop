package models

const (
	GoodsTypeTableName  = "goods_types"
	GoodsTypesTable = "goods_types"
)

type GoodsType struct {
	Model
	Sort      *uint64 `json:"sort"`
	Icon      *string `json:"icon"`
	Name      *string `json:"name"`
	Desc      *string `json:"desc"`
	Reference *uint64 `json:"reference"`
}

type GoodsTypes struct {
	Model
	Sort      *uint64 `json:"sort,omitempty"`
	Name      *string `json:"name,omitempty"`
	Icon      *string `json:"icon,omitempty"`
	Desc      *string `json:"desc,omitempty"`
	Reference *uint64 `json:"reference,omitempty"`
}
