package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cc "github.com/ravayak/copee_backend/apis/controllers/clients"
	guc "github.com/ravayak/copee_backend/apis/controllers/groups/users"
	"github.com/ravayak/copee_backend/app/middlewares"
)

var globalPath = "/api/v1"
var middlewareList = []gin.HandlerFunc{
	middlewares.EnsureLoggedIn(), middlewares.EnsureHasPrivileges(),
}

func mapClientsUrls() {
	clientsRoutes := router.Group(globalPath + "/clients")
	clientsRoutes.GET("/:client_id", cc.Get)
	clientsRoutes.GET("", cc.List)
	clientsRoutes.POST("", cc.Create)
	clientsRoutes.PUT(":client_id", cc.Update)
	clientsRoutes.DELETE(":client_id", cc.Delete)
}

func mapGroupsUsersUrls() {
	groupsUsersRoutes := router.Group(globalPath + "/groups_users")
	groupsUsersRoutes.POST("", guc.Create)
	// groupsUsersRoutes.GET("", guc.List)
	// groupsUsersRoutes.POST("", guc.Create)
	// groupsUsersRoutes.PUT(":group_user_id", guc.Update)
	// groupsUsersRoutes.DELETE(":group_user_id", guc.Delete)
}

func mapCompaniesUrls() {
	companiesRoutes := router.Group(globalPath + "/companies")
	companiesRoutes.POST("", cc.Create)
	// companiesRoutes.GET("", cc.List)
	// companiesRoutes.GET("/:company_id", cc.Get)
	// companiesRoutes.PUT(":company_id", cc.Update)
	// companiesRoutes.DELETE(":company_id", cc.Delete)
}

// MapUrls maps urls to the associated controllers
func MapUrls() {
	// No Route handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "route not found")
	})

	// No proxies trusted
	router.SetTrustedProxies(nil)
	// set release mode for deploy version
	gin.SetMode(gin.ReleaseMode)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AddAllowHeaders("X-Requested-By")
	router.Use(cors.New(corsConfig))
	// Commenter cette ligne pour tester les points d'entrée
	// de l'api sans les contrôles d'accès réalisés par
	//les middlewares
	router.Use(middlewareList...)

	mapClientsUrls()
	mapGroupsUsersUrls()
	mapCompaniesUrls()
}
