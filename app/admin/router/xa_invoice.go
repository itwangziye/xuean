package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerXaInvoiceRouter)
}

// registerXaInvoiceRouter
func registerXaInvoiceRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.XaInvoice{}
	//r := v1.Group("/xa-invoice").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	//{
	//	r.GET("", actions.PermissionAction(), api.GetPage)
	//	r.GET("/:id", actions.PermissionAction(), api.Get)
	//	r.POST("", api.Insert)
	//	r.PUT("/:id", actions.PermissionAction(), api.Update)
	//	r.DELETE("", api.Delete)
	//	r.POST("/review/:id", api.Review)
	//}
	r := v1.Group("/xa-invoice")
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
		r.POST("/review/:id", api.Review)
	}
}
