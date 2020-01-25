package FarmLandController

import (
	"../../Helpers"
	"../../Models/FarmLandModel"
	"github.com/gin-gonic/gin"
)

func ListLand(c *gin.Context) {
	var land []FarmLandModel.FarmLand

	err := FarmLandModel.Get(&land)
	if err != nil {
		Helpers.RespondJSON(c, 200, land)
	} else {
		Helpers.RespondJSON(c, 404, land)
	}
}

func Create(c *gin.Context) {
	var land FarmLandModel.FarmLand
	c.Bind(&land)
	example := c.MustGet("accountid").(string)
	err := FarmLandModel.Create(&land)
	if err != nil {
		Helpers.RespondJSON(c, 200, example)
	} else {
		Helpers.RespondJSON(c, 404, land)
	}
}
