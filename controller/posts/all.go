package posts

import (
	"github.com/gin-gonic/gin"
	"jun/dto"
)

type RootController struct {
	LowestRole dto.UserRole
}

func (p *RootController) DeleteHandler() gin.HandlerFunc {
	return nil
}

func (p *RootController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (p *RootController) PostHandler() gin.HandlerFunc {
	return nil
}

func (p *RootController) PutHandler() gin.HandlerFunc {
	return nil
}
