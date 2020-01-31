package Land

import (
	"../../Helpers"
	"github.com/gin-gonic/gin"
)

func ListLand(c *gin.Context) {
	var land []Land

	err := Get(&land)
	if err != nil {
		Helpers.RespondJSON(c, 200, land)
	} else {
		Helpers.RespondJSON(c, 404, land)
	}
}

func Create(c *gin.Context) {
	var land Land
	c.Bind(&land)
	example := c.MustGet("accountid").(string)
	err := CreateData(&land)
	if err != nil {
		Helpers.RespondJSON(c, 200, example)
	} else {
		Helpers.RespondJSON(c, 404, land)
	}
}
