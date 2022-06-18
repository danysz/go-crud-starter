package crud

import (
	"fmt"
	"github.com/ElegantSoft/go-crud-starter/common"
	"github.com/ElegantSoft/go-crud-starter/db"
	"github.com/ElegantSoft/go-crud-starter/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", func(ctx *gin.Context) {
		var api GetAllRequest
		if err := ctx.ShouldBindQuery(&api); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var s = NewRepository[models.Post](db.DB, &models.Post{})

		var result []models.Post
		err := s.Find(api, &result)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//var data interface{}
		//var totalRows int64
		//tx.Count(&totalRows)
		//if api.Page > 0 {
		//	data = map[string]interface{}{
		//		//"data":       result,
		//		//"total":      totalRows,
		//		//"totalPages": int(math.Ceil(float64(totalRows) / float64(api.Limit))),
		//	}
		//} else {
		//	//data = result
		//}
		ctx.JSON(200, gin.H{"data": result})
	})

	routerGroup.GET(":id", func(ctx *gin.Context) {
		var api GetAllRequest
		var item ById
		if err := ctx.ShouldBindQuery(&api); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := ctx.ShouldBindUri(&item); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": common.ValidateErrors(err)})
			return
		}
		var s = NewRepository[models.Post](db.DB, &models.Post{})

		api.Filter = append(api.Filter, fmt.Sprintf("id||$eq||%s", item.ID))

		var result models.Post

		err := s.FindOne(api, &result)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//var data interface{}
		//var totalRows int64
		//tx.Count(&totalRows)
		//if api.Page > 0 {
		//	data = map[string]interface{}{
		//		//"data":       result,
		//		//"total":      totalRows,
		//		//"totalPages": int(math.Ceil(float64(totalRows) / float64(api.Limit))),
		//	}
		//} else {
		//	//data = result
		//}
		ctx.JSON(200, result)
	})
}
