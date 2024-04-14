package caches

import (
	"banner_service/internal/logger"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func cacheKey(r *http.Request) string {
	return r.URL.String()
}

type cacheWriter struct {
	http.ResponseWriter
	body []byte
}

func (w cacheWriter) Write(b []byte) (int, error) {
	w.body = b
	return w.ResponseWriter.Write(b)
}

func (w cacheWriter) Body() []byte {
	return w.body
}

func isLast(query url.Values) bool {
	flag := query.Get("use_last_revision")

	if strings.ToLower(flag) == "true" {
		return false
	}

	return true
}

func (c *Cache) CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isLast(r.URL.Query()) {
			cachedData, err := c.Get(cacheKey(r))
			if err == nil {
				w.Write([]byte(cachedData))
				return
			} else if cachedData != "" { // отловить ошибку на пустоту
				logger.Error("the cache is not get: %v", err)
			}
		}

		next.ServeHTTP(w, r)

		cached := w.(cacheWriter).Body()
		err := c.Set(cacheKey(r), cached, 5*time.Minute)
		if err != nil {
			logger.Error("the cache is not saved: %v", err)
		}
	})
}
