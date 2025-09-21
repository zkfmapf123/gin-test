package validate

import "github.com/gin-gonic/gin"

func BindJSON[T any](c *gin.Context) (T, error) {

	var obj T
	if err := c.ShouldBindJSON(&obj); err != nil {
		return obj, err
	}

	return obj, nil
}
