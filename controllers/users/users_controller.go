package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jamalakida20/bookstore_users-api/domain/users"
	"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context)  {
	var user users.User
	fmt.Println(user)
	bytes, err:=ioutil.ReadAll(c.Request.Body)
	fmt.Println(err)
	fmt.Println(string(bytes))
	c.String(http.StatusNotImplemented," implement me!")
}
func GetUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented," implement me!")
}
