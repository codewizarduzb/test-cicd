package v1

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// GET PINGPONG
// @Summary GET PINGPONG
// @Description Api for getting PINGPONG
// @Tags PINGPONG
// @Accept json
// @Produce json
// @Param word query string true "word"
// @Success 200 {object} models.PingPong
// @Failure 404 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/pong [GET]
func (h handlerV1) PingPong(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	_, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	word := c.Query("word")


	c.JSON(200, gin.H{
		"message": word + "pong",
	})
}