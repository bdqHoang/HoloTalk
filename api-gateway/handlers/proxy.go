package handlers

import (
	"net/http/httputil"
	"net/url"
	
	"github.com/gin-gonic/gin"
)

func ProxyHandler(target string) gin.HandlerFunc {
	url, _:= url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	return func(c *gin.Context) {
		// add header x-reequest-id
		if reequestID, exists := c.Get("x-request-id"); exists {
			c.Request.Header.Add("x-request-id", reequestID.(string))
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}