package posts

import (
	"github.com/gin-gonic/gin"
)

type PostsController struct{}

func (p *PostsController) DeleteHandler() gin.HandlerFunc {
	return nil
}

func (p *PostsController) GetHandler() gin.HandlerFunc {
	panic("implement me")
}

func (p *PostsController) PostHandler() gin.HandlerFunc {
	panic("implement me")
}

func (p *PostsController) PutHandler() gin.HandlerFunc {
	panic("implement me")
}
