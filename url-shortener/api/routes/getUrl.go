package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohitbisht01/url-shortener/api/database"
)

func GetByShortId(c *gin.Context) {
	shortId := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortId).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found for given shortId"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": val})

}
