package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohitbisht01/url-shortener/api/database"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag"`
}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest

	if err := c.ShouldBind(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
		return
	}

	shortID := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	// checking if shortID already exists in the database
	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shortID not found in the database"})
		return
	}

	var data map[string]interface{}

	if err := json.Unmarshal([]byte(val), &data); err != nil {
		// data is comming in string format not in json format
		data = make(map[string]interface{})
		data["data"] = val
	}

	// cehcking if tag field already exists with the shortId and is a slice
	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTags, ok := t.(string); ok {
				tags = append(tags, strTags)
			}
		}
	}

	// check for duplicate tags
	for _, existingTag := range tags {
		if existingTag == tag {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tag already exists"})
			return
		}
	}

	// no duplicate tag then append the new tag in already existing tag
	tags = append(tags, tag)
	data["tags"] = tags

	updatedData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Marshal updated data"})
		return
	}

	// shortId, data, expiration time
	err = r.Set(database.Ctx, shortID, updatedData, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the tags"})
		return
	}

	c.JSON(http.StatusOK, updatedData)
}
