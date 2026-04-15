package controller

import (
	dto "example.com/t/api/DTO"
	"example.com/t/api/service"
	"example.com/t/core"
	"example.com/t/core/logger"
	"example.com/t/types"
	"example.com/t/utility/reponse"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConfigController struct {
	BaseController
	service        *service.ConfigService
	upload_service *service.UploadService
}

func NewConfigController(
	app *core.AppServer,
	db *gorm.DB,
	service *service.ConfigService,
	upload_service *service.UploadService,
) *ConfigController {
	return &ConfigController{
		BaseController: BaseController{
			App: app,
			DB:  db,
		},
		service:        service,
		upload_service: upload_service,
	}
}

func (config *ConfigController) RegisterRouter() {
	group := config.App.Engine.Group("/config")
	group.POST("insert", config.insert)
	group.POST("insertMany", config.insertMany)
	group.POST("upload", config.upload)
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
	logger.GlobalLog.Info("====================日志测试=====================")
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

// @Summary      上传单个文件
// @Description  通过 multipart/form-data 上传文件，支持图片、文档等类型
// @Tags         Config
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "要上传的文件"
// @Success      200   {object}  map[string]string  "上传成功"
// @Failure      400   {object}  map[string]string  "请求参数错误"
// @Failure      500   {object}  map[string]string  "服务器内部错误"
// @Router       /config/upload [post]
func (config *ConfigController) upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reponse.ERROR(c, types.InvalidArgs)
	}
	err2 := config.upload_service.Upload(c, file)
	if err2 != nil {
		reponse.ERROR(c, types.UploadFaild)
		return
	}
	reponse.SUCCESS(c, types.OkMsg)
}
