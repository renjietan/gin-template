package controller

import (
	dto "example.com/t/api/DTO"
	"example.com/t/api/service"
	"example.com/t/core"
	"example.com/t/types"
	"example.com/t/utility/reponse"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConfigController struct {
	BaseController
	service *service.ConfigService
}

func NewConfigController(
	app *core.AppServer,
	db *gorm.DB,
	service *service.ConfigService,
) *ConfigController {
	return &ConfigController{
		BaseController: BaseController{
			App: app,
			DB:  db,
		},
		service: service,
	}
}

func (config *ConfigController) RegisterRouter() {
	group := config.App.Engine.Group("/config")
	group.POST("insert", config.insert)
	group.POST("insertMany", config.insertMany)
}

// @Summary 创建 配置
// @Description 配置 描述
// @Tags Config
// @Accept json
// @Produce json
// @Param request body dto.ConfigDTO true "配置结构体"
// @Success 200 {object} entity.ConfigEntity
// @Router /config/insert [post]
func (config *ConfigController) insert(c *gin.Context) {
	var d dto.ConfigDTO
	if err := c.ShouldBindJSON(&d); err != nil {
		reponse.ERROR(c, types.InvalidArgs)
		return
	}
	res := config.service.Insert(d)
	reponse.SUCCESS(c, res)
}

// @Summary 批量创建 配置
// @Description 批量配置 描述
// @Tags Config
// @Accept json
// @Produce json
// @Param request body dto.ConfigsDTO true "配置结构体"
// @Success 200 {object} []entity.ConfigEntity
// @Router /config/insertMany [post]
func (config *ConfigController) insertMany(c *gin.Context) {
	var d dto.ConfigsDTO
	if err := c.ShouldBindJSON(&d); err != nil {
		reponse.ERROR(c, types.InvalidArgs)
		return
	}
	res := config.service.InsertMany(d)
	reponse.SUCCESS(c, res)
}
