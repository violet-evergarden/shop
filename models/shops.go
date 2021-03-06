package models

import (
	"time"
)

const (
	ShopTableName = "shops"
	ShopsTable    = "shops"
)

type Shops struct {
	Model
	Status            *uint64    `json:"status"`
	FirstBusinessTime *time.Time `json:"firstBusinessTime"`
	ShopCode          *string    `json:"shopCode"`
	ShopAddr          *string    `json:"shopAddr"`
	ShopCoor          *string    `json:"shopCoor"`
	NumOfDesk         *uint64    `json:"numOfDesk"`
	NumOfChair        *uint64    `json:"numOfChair"`
	Desc              *string    `json:"desc"`
}

func (m *Shops) Table() string {
	return "shops"
}

//
//var ModelShop = modelShop{
//	Model{TableName:"shop"},
//}
//type modelShop struct {
//	Model
//}
//
//func (m *modelShop) Search(status int,shopCode string,firstBusinessTime *time.Time,shopAddr string) (list []Shop,total int,err error) {
//	db := DB.Table(m.TableName)
//
//	if shopCode != "" {
//		db =db.Where("shop_code = ?",shopCode)
//	}
//	if status >= 0 {
//		db = db.Where("status = ?",status)
//	}
//	if firstBusinessTime != nil {
//		db = db.Where("first_business_time = ?",firstBusinessTime)
//	}
//	if shopAddr != "" {
//		db = db.Where("shop_addr = ?",shopAddr)
//	}
//
//	db.Count(&total)
//
//	res := db.Find(&list)
//	if res.Error != nil {
//		err = res.Error
//		return
//	}
//
//	return
//}
//
//func (m *modelShop) Insert(shop *Shop) error {
//	res := DB.Table(m.TableName).Create(shop)
//	if res.Error != nil {
//		code,err := GetErrorCode(res.Error)
//		if err != nil {
//			return err
//		}
//		if code == 1062 {
//			return fmt.Errorf("添加失败，可能商店地址重复了")
//		}
//		return res.Error
//	}
//	return nil
//}
//
//func (m *modelShop) Goods() {
//
//}
