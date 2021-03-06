package user

import (
	"net/http"
	"strconv"

	"github.com/5107293001/contact-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type resource struct {
	service ServiceInterface
}

func RegisterRoutes(r *gin.Engine, service ServiceInterface) {
	resource := &resource{
		service: service,
	}
	r.GET("/users", resource.Query)
	r.POST("/users", resource.Create)
	r.GET("/users/:id", resource.Get)
	r.PUT("/users/:id", resource.Put)
	r.DELETE("/users/:id", resource.Delete)
}

func (resource *resource) Query(c *gin.Context) {
	page, limit, err := utils.Pagination(c)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	userList, err := resource.service.Query(page*limit, limit, c.Query("q"))
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "users list",
		"data":    userList,
	})
}
func (resource *resource) Create(c *gin.Context) {
	user := User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := resource.service.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully Created User",
		"data":    user,
	})
}

func (resource *resource) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	user, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get user",
		"data":    user,
	})

}

func (resource *resource) Put(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	updateUser := User{}
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := resource.service.Update(uint(id), &updateUser)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"message": "Update the user",
		"data":    user,
	})
}

func (resource *resource) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	resource.service.Delete(uint(id))
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"message": "successfully deleted",
	})

}
