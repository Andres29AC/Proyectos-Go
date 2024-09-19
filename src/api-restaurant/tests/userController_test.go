package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/Andres29AC/Proyectos-Go/src/api-restaurant/controllers"
)

func TestGetUser(t *testing.T) {
	h := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(h)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "66ecb27fac31ebbbd146f625"}}
	c.Request, _ = http.NewRequest("GET", "/users/66ecb27fac31ebbbd146f625", nil)

	controllers.GetUser()(c)

	assert.Equal(t, http.StatusOK, h.Code)
	assert.Contains(t, h.Body.String(), "Juan")
}

//NOTE:go test ./tests -v
