package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/middleware/cors"
	"github.com/routers/api"
	"github.com/routers/backend"
)

func InitRouterForUsers() *gin.Engine {
	r := gin.New()
	r.Use(cors.CorsHandler())
	apiAuth := backend.ApiAuth{}


	apiAdmin := backend.ApiAdmins{}
	r.POST("/login",apiAdmin.Login)

	//user := r.Group("/user")
	//user.Use(backend.AuthVerify)
	//user.POST("/ping",backend.Ping)

	goods := r.Group("/goods")
	goods.Use(func(context *gin.Context) {
		// JWT 校验
		// 得到用户数据
		// 下面的接口就已经可以访问
		// 不然就 context.Abort()

	})
	{
		apiGoods := backend.Goods{}
		goods.POST("/insert",apiGoods.Insert)
		goods.POST("/delete",apiGoods.Delete)
		goods.POST("/update",apiGoods.Update)
		goods.POST("/search",apiGoods.Search)

		apiGoodsType := backend.ApiGoodsType{}
		goods.POST("/type/insert",apiGoodsType.Insert)
		goods.POST("/type/update",apiGoodsType.Update)
		goods.POST("/type/search",apiGoodsType.Search)
	}

	franchisees := r.Group("/franchisees")
	franchisees.Use(func(ctx *gin.Context) {
		if !apiAuth.Franchisees(ctx) {return}
	})
	{
		apiFranchisees := backend.ApiFranchisees{}
		franchisees.POST("/insert",apiFranchisees.Insert)
		franchisees.POST("/update",apiFranchisees.Update)
		franchisees.POST("/search",apiFranchisees.Search)
	}

	apiUsers := api.Users{}
	users := r.Group("/users")
	{
		users.POST("/order/token",apiUsers.OrderToken)
		users.POST("/order/order",apiUsers.Order)
		users.POST("/order/search",apiUsers.OrderSearch)
		users.POST("/order/details",apiUsers.OrderDetails)
		users.POST("/info",apiUsers.Info)
	}

	frontendUsersOrders := r.Group("/frontend/users")
	{
		apiUsers := api.Users{}
		frontendUsersOrders.POST("/order/token",apiUsers.OrderToken)
		frontendUsersOrders.POST("/order/order",apiUsers.Order)
		frontendUsersOrders.POST("/order/search",apiUsers.OrderSearch)
		frontendUsersOrders.POST("/order/details",apiUsers.OrderDetails)
		frontendUsersOrders.POST("/info",apiUsers.Info)
	}

	frontendShops := r.Group("/shops")
	{
		apiShops := api.Shops{}
		frontendShops.POST("/goods/types",apiShops.GoodsTypes)
		frontendShops.POST("/goods/list",apiShops.GoodsList)
		frontendShops.POST("/goods/search",apiShops.Search)
	}

	//apiConsumers := backend.ApiConsumers{}
	//consumers := r.Group("/consumers")
	//{
	//	consumers.POST("/order",apiConsumers.Order)
	//}

	return r
}