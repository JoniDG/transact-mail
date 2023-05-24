package controller

import (
	"github.com/JoniDG/transact-mail/internal/service"
)

type FileController interface {
	HandlerFile(fileName string) error
}

type fileController struct {
	svc service.FileService
}

func NewFileController(svc service.FileService) FileController {
	return &fileController{svc: svc}
}

func (c *fileController) HandlerFile(fileName string) error {
	return c.svc.HandlerFile(fileName)
}
