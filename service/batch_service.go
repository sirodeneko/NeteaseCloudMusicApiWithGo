package service

import (
	"github.com/gin-gonic/gin"
)

type BatchService struct {
}

func (service *BatchService) Batch(c *gin.Context) map[string]interface{} {

	reBody := make(map[string]interface{})
	reBody["code"] = "500"
	reBody["err"] = "该接口未实现"
	return reBody
}
