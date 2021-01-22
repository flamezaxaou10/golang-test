package beer

import (
	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.Engine) {
	b := r.Group("/beer")
	b.GET("/", getAllBeer)
}

type BeerResponse struct {
	Message string `json:message`
}

// get all beer
// @Summary Retrieves users from mongodb
// @Description Get ALL Beers
// @Produce json
// @Success 200 {object} BeerResponse
// @Router /beer/ [get]
func getAllBeer(c *gin.Context) {
	//TODO :: call service
	c.JSON(200, gin.H{
		"message": "getAllBeer",
	})
}
