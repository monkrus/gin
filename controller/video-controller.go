package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/monkrus/gin/entity"
	"github.com/monkrus/gin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

// implement that interface
type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return service.New().FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
