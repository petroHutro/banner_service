package caches

import (
	"banner_service/internal/logger"
	"fmt"
	"net/http"
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

// the test piece
func (c *Cache) CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cachedData, err := c.Get(cacheKey(r))
		if err == nil {
			w.Write([]byte(cachedData))
			return
		} else if cachedData != "" { // проверить ошибки
			http.Error(w, "Ошибка при поиске данных в кеше", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)

		cached := w.(cacheWriter).Body()
		err = c.Set(cacheKey(r), cached, 5*time.Minute)
		if err != nil {
			logger.Error("cookies do not contain a token: %v", err)
			fmt.Println("Ошибка при сохранении данных в кеше:", err)
		}
	})
}
