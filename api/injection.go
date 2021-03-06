package main

import (
	"time"

	"github.com/Strum355/log"
	"github.com/gal/timber/controllers"
	"github.com/gal/timber/middlewares"
	"github.com/gal/timber/models"
	"github.com/gal/timber/views"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func inject(d *dataSources) (*gin.Engine, error) {
	log.Info("Injecting data sources")

	// Migrate postgres tables
	d.DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Application{}, &models.Tag{})

	d.DB.Preload("Projects").Preload("Applications").Find(&models.User{})
	d.DB.Preload("Collaborators").Preload("Applications").Find(&models.Project{})
	d.DB.Preload("Tags").Find(&models.User{})

	//// Drop all exising tables from the postgres database (d.DB) using raw SQL
	//err := d.DB.Unscoped().Where("id > ?", -1).Delete(&models.Tag{}).Error
	//if err != nil {
	//	log.WithError(err).Error("Error dropping tags")
	//}
	//err = d.DB.Unscoped().Where("id > ?", -1).Delete(&models.User{}).Error
	//if err != nil {
	//	log.WithError(err).Error("Error dropping users")
	//}
	//err = d.DB.Unscoped().Where("id > ?", -1).Delete(&models.Project{}).Error
	//if err != nil {
	//	log.WithError(err).Error("Error dropping projects")
	//}
	//err = d.DB.Unscoped().Where("id > ?", -1).Delete(&models.Application{}).Error
	//if err != nil {
	//	log.WithError(err).Error("Error dropping applications")
	//}

	/*
	 * Model layer
	 */
	userStore := models.NewUserStore(d.DB)
	projectStore := models.NewProjectStore(d.DB)
	tagStore := models.NewTagStore(d.DB)
	applicationStore := models.NewApplicationStore(d.DB)
	tokenStore := models.NewTokenStore(d.RedisClient)

	// Seeding
	seedStore := models.NewSeedStore(d.DB)
	seedStore.Seed()

	/*
	 * Controller layer
	 */
	userController := controllers.NewUserController(*userStore)
	projectController := controllers.NewProjectController(*projectStore)
	applicationController := controllers.NewApplicationController(*applicationStore)
	tagController := controllers.NewTagController(*tagStore)
	// Load auth key from env variable
	authKey := viper.GetString("auth.key")

	// Load expiration lengths
	accessTokenExp := int64(time.Minute * 10)
	refreshTokenExp := int64(time.Hour * 24 * 30)

	tokenController := controllers.NewTokenController(*tokenStore, *userStore, authKey, accessTokenExp, refreshTokenExp)

	// Initialize router
	router := gin.New()
	router.Use(middlewares.JSONLogMiddleware())
	router.Use(gin.Recovery())

	views.NewHandler(&views.Config{
		R:                     router,
		UserController:        *userController,
		ProjectController:     *projectController,
		ApplicationController: *applicationController,
		TokenController:       *tokenController,
		TagController:         *tagController,
	})

	return router, nil
}
