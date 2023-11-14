package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetClaimInResponse[K comparable, L any](c *gin.Context, name K, value L) {
	c.Request = c.Request.WithContext(
		context.WithValue(
			c.Request.Context(),
			name,
			value,
		),
	)
}

func GetClaimFromResponse[K comparable, L any](c *gin.Context, name K) (L, bool) {
	value, ok := c.Request.Context().Value(name).(map[string]string)
	var result L
	if !ok {
		return result, false
	}
	str, marshalError := json.Marshal(value)
	err := json.Unmarshal(str, &result)
	if err != nil || marshalError != nil {
		fmt.Print(err)
		return c.Request.Context().Value(name).(L), false
	}

	return result, true
}
