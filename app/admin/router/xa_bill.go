package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerXaBillRouter)
}

// registerXaBillRouter
func registerXaBillRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.XaBill{}
	r := v1.Group("/xa-bill").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.POST("/review/:id", actions.PermissionAction(), api.Review)
		r.DELETE("", api.Delete)
	}

	//r := v1.Group("/xa-bill")
	//{
	//	r.GET("", actions.PermissionAction(), api.GetPage)
	//	r.GET("/:id", actions.PermissionAction(), api.Get)
	//	r.POST("", api.Insert)
	//	r.PUT("/:id", actions.PermissionAction(), api.Update)
	//	r.POST("/review/:id", actions.PermissionAction(), api.Review)
	//	r.DELETE("", api.Delete)
	//}
}
