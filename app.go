package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func main() {
	router := gin.Default()

	router.GET("/", homeHandler)
	router.GET("/login", loginHandler)
	router.GET("/logout", logoutHandler)

	router.Run(":8080")
}

func homeHandler(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	c.String(http.StatusOK, "Welcome to the home page!")
}

func loginHandler(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")

	session.Values["authenticated"] = true
	session.Save(c.Request, c.Writer)

	c.String(http.StatusOK, "Berhasil login")
}

func logoutHandler(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")

	delete(session.Values, "authenticated")
	session.Save(c.Request, c.Writer)

	c.String(http.StatusOK, "Berhasil logout")
}
