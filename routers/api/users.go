package api

import (
	"context"
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
	"github.com/proto"
	"net/http"
	"time"
)

/*
	客户端下单前，需要先获取一个订单令牌
	客户端拿着令牌进行下单，这样消息就可以不断重发而不重复下单
*/

type Users struct{}

func (api *Users) Info(ctx *gin.Context) {

}

func (api *Users) OrderToken(ctx *gin.Context) {
	var operatorId uint64 = 13

	// 生成一个订单号
	cli := proto.NewOrderNoService("micro.service.order.no", app.Service().Client())
	rsp, err := cli.Gen(context.TODO(), &proto.OrderNoGenRequest{})
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}
	no := rsp.No

	fmt.Println(no)

	// 加密订单号
	token := ""
	{
		rsp, err := cli.Encrypt(context.TODO(), &proto.OrderNoEncryptReq{
			Data: &proto.OrderNoEncryptFormat{
				No:      no,
				Id:      operatorId,
				Code:    "",
				Time:    uint64(time.Now().Unix()),
				Timeout: 300,
			},
		})
		if err != nil {
			app.ResponseError(ctx, err)
			return
		}
		token = rsp.Token
	}

	if token == "" {
		fmt.Println(rsp)
	}

	// 返回消息
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"token":   token,
	})
}

/*
	消费者，下单操作，生成一个订单，有商品的明细
*/
func (api *Users) Order(ctx *gin.Context) {
	type _params struct {
		Token     string
		ShopId    uint64
		GoodsList []uint64
	}

	params := _params{}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	token := params.Token
	cli := proto.NewOrderNoService("micro.service.order.no", app.Service().Client())
	rsp, err := cli.Decrypt(context.TODO(), &proto.OrderNoDecryptReq{
		Token: token,
	})
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	if uint64(time.Now().Unix()) > rsp.Data.Time+rsp.Data.Timeout {
		app.ResponseError(ctx, fmt.Errorf("订单号已经超时"))
		return
	}

	//id := rsp.Data.Id
	//no := rsp.Data.No
	//
	//{
	//	cli := proto.NewUsersOrdersService("micro.service.users", app.Service().Client())
	//	rsp, err := cli.Order(ctx, &proto.UsersOrdersReq{
	//		UserId: id,
	//		OrderNo:    no,
	//		ShopId:     params.ShopId,
	//		GoodsId:    params.GoodsList,
	//	})
	//	if err != nil {
	//		app.ResponseMicroServicesError(ctx, err)
	//		return
	//	}
	//	if rsp.Code != 0 {
	//		app.Response(ctx, rsp.Code, rsp.Message)
	//		return
	//	}
	//
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code":    0,
	//		"message": "下单成功",
	//		"orderNo": rsp.OrderNo,
	//		"cost":    rsp.Cost,
	//	})
	//}
}

/*
	只能查询一定时间范围
	近半年，近1年，近3年，这样的时间区
	页数
	只能查询成功消费的订单
*/
func (api *Users) OrderSearch(ctx *gin.Context) {
	type _params struct {
		TimeZone string
		app.PageParams
	}
	params := _params{}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	now := time.Now()
	day, err := time.ParseDuration("-24h")
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	var total uint64 = 0
	var orders []models.UsersOrders
	db := models.DB.Model(models.UsersOrders{})
	db = db.Where("consumer_id = ?", 1)
	switch params.TimeZone {
	case "halfYear":
		t0 := fmt.Sprintf("%s 00:00:00", now.Add(day*180).Format("2006-01-02"))
		t1 := fmt.Sprintf("%s 23:59:59", now.Format("2006-01-02"))
		db = db.Where("payment_at between ? and ?", t0, t1)
		break
	case "year":
		t0 := fmt.Sprintf("%s 00:00:00", now.Add(day*365).Format("2006-01-02"))
		t1 := fmt.Sprintf("%s 23:59:59", now.Format("2006-01-02"))
		db = db.Where("payment_at between ? and ?", t0, t1)
		break
	case "threeYears":
		t0 := fmt.Sprintf("%s 00:00:00", now.Add(day*365*3).Format("2006-01-02"))
		t1 := fmt.Sprintf("%s 23:59:59", now.Format("2006-01-02"))
		db = db.Where("payment_at between ? and ?", t0, t1)
		break
	default:
		app.ResponseError(ctx, fmt.Errorf("未知的时间区间"))
		return
	}
	db.Count(&total)
	res := db.Offset(params.Page * 20).Limit(20).Debug().Find(&orders)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, total, orders)
}

/*
	用户接口，需要登录状态
	通过订单号，查询订单下的商品明细
*/
func (api *Users) OrderDetails(ctx *gin.Context) {
	//consumerId := uint64(1)
	//
	//type _params struct {
	//	OrderNo uint64
	//}
	//params := _params{}
	//err := ctx.BindJSON(&params)
	//if err != nil {
	//	app.ResponseError(ctx, err)
	//}

	//// 需要校验订单号，是否属于此用户
	//order := models.UsersOrders{
	//	//Model:      models.Model{},
	//	ConsumerId: &consumerId,
	//	OrderNo:    &params.OrderNo,
	//	Status:     nil,
	//	Total:      nil,
	//	ShopId:     nil,
	//	ShopAddr:   nil,
	//	ShopCode:   nil,
	//	PaymentAt:  nil,
	//}
	//res := models.DB.Model(models.UsersOrders{}).Where(order).Debug().Find(&models.UsersOrders{})
	//if res.Error != nil {
	//	if res.RecordNotFound() {
	//		app.ResponseError(ctx, fmt.Errorf("订单号不存在"))
	//	} else {
	//		app.ResponseError(ctx, res.Error)
	//	}
	//	return
	//}
	//
	//// 查询订单明细
	//var goods []models.UsersOrdersDetails
	//res = models.DB.Model(models.UsersOrdersDetails{}).Where("order_no = ?", params.OrderNo).Find(&goods)
	//if res.Error != nil {
	//	app.ResponseError(ctx, res.Error)
	//	return
	//}
	//app.ResponseSearchSuccess(ctx, uint64(len(goods)), goods)
}
