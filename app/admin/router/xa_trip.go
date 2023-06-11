package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerXaTripRouter)
}

// registerXaTripRouter
func registerXaTripRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.XaTrip{}
	r := v1.Group("/xa-trip").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("/add", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
		r.POST("/list", api.GetTripList)
		r.POST("/review/:id", actions.PermissionAction(), api.Review)
	}

	//r := v1.Group("/xa-trip")
	//{
	//	r.GET("", actions.PermissionAction(), api.GetPage)
	//	r.GET("/:id", actions.PermissionAction(), api.Get)
	//	r.POST("/add", api.Insert)
	//	r.PUT("/:id", actions.PermissionAction(), api.Update)
	//	r.DELETE("", api.Delete)
	//	r.POST("/list", api.GetTripList)
	//	r.POST("/review/:id", actions.PermissionAction(), api.Review)
	//}
}
