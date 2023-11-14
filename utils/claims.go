package utils

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func SetClaimInResponse[K comparable](c *gin.Context, name K, value map[string]any) {
	c.Request = c.Request.WithContext(
		context.WithValue(
			c.Request.Context(),
			name,
			value,
		),
	)
}

func GetClaimFromResponse[K comparable, L any](c *gin.Context, name K) (L, bool) {
	value, ok := c.Request.Context().Value(name).(map[string]any)
	var result L
	if !ok {
		return result, false
	}
	str, marshalError := json.Marshal(value)
	err := json.Unmarshal(str, &result)
	if err != nil || marshalError != nil {
		return c.Request.Context().Value(name).(L), true
	}

	return result, true
}
