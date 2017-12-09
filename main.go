package main

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mjibson/goon"
	"lib/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"github.com/icza/session"
)

func init() {
	r := gin.New()

	r.Static("/assets", "./assets")
	r.Static("/node_modules", "./node_modules")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		id := getSessionUserId(c.Request)

		if (id > 0 ) {
			appengineContext := appengine.NewContext(c.Request)
			log.Debugf(appengineContext, "loginUserId: %v", id)
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "INDEX_PAGE",
				"message":"ログイン済み " + "UserID:" + strconv.Itoa(id),
			})
		} else {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "INDEX_PAGE",
				"message":"未ログイン",
			})
		}

	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/data_store_sample/:id", func(c *gin.Context) {
		id := c.Param("id")
		appengineContext := appengine.NewContext(c.Request)
		g := goon.FromContext(appengineContext)

		user := &models.User{Id:id}
		err := g.Get(user)
		if err!= nil {
			c.String(http.StatusNotFound, "NotFound")
		} else {
			c.JSON(http.StatusOK, user)
		}
	})

	r.POST("/data_store_sample", func(c *gin.Context) {
		id := c.PostForm("id")
		name := c.PostForm("name")
		appengineContext := appengine.NewContext(c.Request)
		g := goon.FromContext(appengineContext)

		var user = &models.User{
			Id: id,
			Name: name,
		}
		g.Put(user)
		c.JSON(http.StatusOK, user)
	})

	r.GET("/login", func(c *gin.Context) {
		// とりあえず、session作る処理だけ
		setSession(c.Writer,c.Request, 1)
		c.String(http.StatusOK, "create session")
	})

	r.GET("/logout", func(c *gin.Context) {
		// とりあえず、session消す処理だけ
		removeSession(c.Writer,c.Request)
		c.String(http.StatusOK, "delete session")
	})

	http.Handle("/", r)
}

func getSessionUserId ( r *http.Request)  int {
	appengineContext := appengine.NewContext(r)
	sessmgr := session.NewCookieManagerOptions(session.NewMemcacheStore(appengineContext), &session.CookieMngrOptions{AllowHTTP: true})
	defer sessmgr.Close() // This will ensure changes made to the session are auto-saved
	// in Memcache (and optionally in the Datastore).
	sess := sessmgr.Get(r) // Get current session
	log.Debugf(appengineContext, "debug: %v", sess)
	if sess != nil {
		userId := sess.CAttr("UserId")
		if userId, ok := userId.(int); ok{
			return userId
		}
	}
	return 0
}


func setSession (w http.ResponseWriter, r *http.Request, userId int)  {
	appengineContext := appengine.NewContext(r)
	sessmgr := session.NewCookieManagerOptions(session.NewMemcacheStore(appengineContext), &session.CookieMngrOptions{AllowHTTP: true})
	defer sessmgr.Close() // This will ensure changes made to the session are auto-saved
	// in Memcache (and optionally in the Datastore).
	//sess = session.NewSession()
	sess := session.NewSessionOptions(&session.SessOptions{
		CAttrs: map[string]interface{}{"UserId": userId},
	})
	sessmgr.Add(sess, w)
}

func removeSession (w http.ResponseWriter, r *http.Request)  {
	appengineContext := appengine.NewContext(r)
	sessmgr := session.NewCookieManagerOptions(session.NewMemcacheStore(appengineContext), &session.CookieMngrOptions{AllowHTTP: true})
	defer sessmgr.Close() // This will ensure changes made to the session are auto-saved
	// in Memcache (and optionally in the Datastore).
	sess := sessmgr.Get(r) // Get current session
	sessmgr.Remove(sess, w)
}


