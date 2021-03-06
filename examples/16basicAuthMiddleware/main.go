package main

import "github.com/gin-gonic/gin"

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123456"},
	"austin": gin.H{"email": "austin@xample.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			// {"secret":{"email":"foo@bar.com","phone":"123456"},"user":"foo"}
			c.JSON(200, gin.H{"user": user, "secret": secret})
		} else {
			// {"secret":"NO SECRET :(","user":"manu"}
			c.JSON(200, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	r.Run()
}
