package views

import (
	"github.com/gal/timber/controllers"
	"github.com/gal/timber/middlewares"
	"github.com/gal/timber/utils"
	"github.com/gal/timber/utils/customresponse"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserController        controllers.UserController
	ProjectController     controllers.ProjectController
	ApplicationController controllers.ApplicationController
	TokenController       controllers.TokenController
	TagController         controllers.TagController
}

type Config struct {
	R                     *gin.Engine
	UserController        controllers.UserController
	ProjectController     controllers.ProjectController
	ApplicationController controllers.ApplicationController
	TokenController       controllers.TokenController
	TagController         controllers.TagController
}

func NewHandler(c *Config) {
	h := &Handler{
		c.UserController,
		c.ProjectController,
		c.ApplicationController,
		c.TokenController,
		c.TagController,
	}

	g := c.R.Group("/api")
	g.Use(middlewares.AuthUser(h.TokenController))

	// if gin.Mode() != gin.TestMode {
	g.GET("/profile", h.Profile)
	g.GET("/profile/:userID", h.ProfileByID)

	g.GET("/projects", h.GetProjects)
	g.GET("/projects/:projectID", h.GetProject)
	g.GET("/users/:ownerID/projects", h.GetProjectsByOwnerID)
	g.POST("/projects", h.NewProject)
	g.POST("/projects/:projectID/apply", h.NewApplication)

	g.GET("/applications", h.GetOwnApplications)
	g.GET("/applications/:appID", h.GetApplicationByID)
	g.POST("/applications", h.NewApplication)
	g.GET("/applications/all", h.GetAllApplications)

	g.GET("/projects/:projectID/applications", h.GetApplicationsByProjectID)
	g.GET("/tags", h.GetTags)
	g.GET("/tags/:tagID", h.GetTag)
	g.POST("/tags", h.NewTag)
	g.GET("/projects/:projectID/tags", h.GetTagsByProjectID)

	c.R.NoRoute(func(c *gin.Context) {
		utils.Respond(c, customresponse.NewInternal(), nil)
	})

	authHandler := c.R.Group("/api/auth")

	authHandler.GET("/signin/:provider", h.SignIn)
	authHandler.GET("/callback/:provider", h.OauthCallback)
	authHandler.POST("/tokens", h.Tokens)
}
