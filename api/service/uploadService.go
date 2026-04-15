package service

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"sync"
	"time"

	"example.com/t/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UploadService struct {
	db     *gorm.DB
	lock   sync.Mutex
	config *types.AppConfig
}

func NewUploadService(db *gorm.DB, appConfig *types.AppConfig) *UploadService {
	return &UploadService{
		db:     db,
		lock:   sync.Mutex{},
		config: appConfig,
	}
}

func (service *UploadService) Upload(c *gin.Context, file *multipart.FileHeader) error {
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), "upload", ext)
	savePath := filepath.Join(service.config.UploadPath, newFileName)
	return c.SaveUploadedFile(file, savePath)
}
