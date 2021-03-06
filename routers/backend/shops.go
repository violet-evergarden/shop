package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type ApiShop struct{}


func (api *ApiShop) Insert(ctx *gin.Context) {
	var params struct {
		Status            uint64
		FirstBusinessTime string
		ShopCode          string
		ShopAddr          string
		ShopCoor          string
		NumOfDesk         uint64
		NumOfChair        uint64
		Desc              string
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	businessTime, err := app.ParserDate(params.FirstBusinessTime)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	newData := models.Shops{
		Status:            &params.Status,
		FirstBusinessTime: businessTime,
		ShopCode:          &params.ShopCode,
		ShopAddr:          &params.ShopAddr,
		ShopCoor:          &params.ShopCoor,
		NumOfDesk:         &params.NumOfDesk,
		NumOfChair:        &params.NumOfChair,
		Desc:              &params.Desc,
	}

	res := models.DB.Create(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "新增店铺成功")
}


func (api *ApiShop) Search(ctx *gin.Context) {
	var params struct {
		Status   int64
		ShopCode string
		ShopAddr string
		Desc     string
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	db := models.DB.Table(models.ShopTableName)
	//if params.Status >= 0 {
	//	db = db.Where("status = ?", params.Status)
	//}
	if params.ShopCode != "" {
		db = db.Where("shop_code = ?", params.ShopCode)
	}
	if params.ShopAddr != "" {
		db = db.Where("shop_addr like ?", params.ShopAddr)
	}
	if params.Desc != "" {
		db = db.Where("desc like %?%", params.Desc)
	}

	var total uint64
	var values []models.Shops
	db.Count(&total)
	res := db.Find(&values)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, total, values)
}

func (api *ApiShop) Update(ctx *gin.Context) {
	var params struct {
		Id                uint64
		Status            uint64
		FirstBusinessTime string
		ShopCode          string
		ShopAddr          string
		ShopCoor          string
		NumOfDesk         uint64
		NumOfChair        uint64
		Desc              string
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	businessTime, err := app.ParserDate(params.FirstBusinessTime)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	fmt.Println(businessTime)

	// 查询ID是否存在
	oldData := models.Shops{}
	res := models.DB.First(&oldData, params.Id)
	if res.Error != nil {
		if res.RecordNotFound() {
			app.ResponseError(ctx, fmt.Errorf("ID不存在"))
		} else {
			app.ResponseError(ctx, err)
		}
		return
	}

	newData := models.Shops{
		Status:            &params.Status,
		FirstBusinessTime: businessTime,
		ShopCode:          &params.ShopCode,
		ShopAddr:          &params.ShopAddr,
		ShopCoor:          &params.ShopCoor,
		NumOfDesk:         &params.NumOfDesk,
		NumOfChair:        &params.NumOfChair,
		Desc:              &params.Desc,
	}
	res = models.DB.Model(models.Shops{Model: models.Model{ID: &params.Id}}).Updates(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "更新成功")
}
