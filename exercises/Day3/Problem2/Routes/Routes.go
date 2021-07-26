package Routes

import (
	"github.com/gin-gonic/gin"
	"local/exercises/Day3/Problem2/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/students-api")
	{
		grp1.GET("", Controllers.GetAllStudent)
		grp1.POST("student", Controllers.CreateStudent)
		grp1.GET("student/:id", Controllers.GetStudentByID)
		grp1.PUT("student/:id", Controllers.UpdateStudent)
		grp1.DELETE("student/:id", Controllers.DeleteStudent)
	}
	return r
}
