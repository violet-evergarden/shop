package models

const (
	GoodsTableName = "goods"
	GoodsTable     = "goods"
)

type Goods struct {
	Model
	TypeId            *uint64  `json:"typeId,omitempty"`
	Name              *string  `json:"name,omitempty"`
	ImportPrice       *float64 `json:"importPrice,omitempty"`
	ExportPriceAdvise *float64 `json:"exportPriceAdvise,omitempty"`
	Image             *string  `json:"image,omitempty"`
	Desc              *string  `json:"desc,omitempty"`
}

func (m *Goods) SelectFieldsForClient() []string {
	return []string{"`id`", "`name`", "`export_price`", "`image`", "`desc`"}
}

func (m *Goods) Table() string {
	return "goods"
}

//
//var ModelGoods = modelGoods{
//	Model{TableName:"goods"},
//}
//type modelGoods struct {
//	Model
//}
//
//func (m *modelGoods) Insert(goods *Goods) error {
//	res := DB.Model(m).Create(goods)
//	if res.Error != nil {
//		code,err := GetErrorCode(res.Error)
//		if err != nil {
//			return err
//		}
//		if code == 1062 {
//			return fmt.Errorf("商品唯一信息冲突")
//		}
//		return res.Error
//	}
//
//	return nil
//}
//
//func (m *modelGoods) Search(id uint,page uint) ([]Goods,error) {
//	var list []Goods
//
//	db := DB.Table(m.TableName)
//
//	if id > 0 {
//		db = db.Where("id = ?",id)
//	}
//
//	res := db.Offset(20*page).Limit(20).Find(&list)
//	if res.Error != nil {
//		return nil,res.Error
//	}
//	return list,nil
//}
//
//func (m *modelGoods) DeleteById(id int) {
//
//}
//
