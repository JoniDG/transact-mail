package controller

type FileController interface {
}

type fileController struct {
}

func NewFileController() FileController {
	return &fileController{}
}
