package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rohitbisht01/url-shortener/api/database"
	"github.com/rohitbisht01/url-shortener/api/models"
)

func EditUrl(c *gin.Context) {
	shortId := c.Param("shortId")
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Parse Json"})
	}

	r := database.CreateClient(0)
	defer r.Close()

	// checking if the shortID exists in the DB or not
	val, err := r.Get(database.Ctx, shortId).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "ShortIds doesn;t exists"})
	}

	// udpate the content of the URL with the shortID, expiry time with shortID
	err = r.Set(database.Ctx, shortId, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update the shortend content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The content has been updated"})

}
