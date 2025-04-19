package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Stream info
type Stream struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	PlaybackURL string `json:"playbackUrl"`
}

var (
	streams   = make([]Stream, 0)
	streamsMu sync.RWMutex
)

func main() {
	// Poblamos ejemplo
	streamsMu.Lock()
	streams = append(streams, Stream{
		ID:          "live_abc",
		Title:       "Demo Stream",
		PlaybackURL: "http://localhost:8080/hls/live_abc.m3u8",
	})
	streamsMu.Unlock()

	r := gin.Default()

	// Sirve estáticos HLS (no necesario si lo hace Nginx)
	// r.Static("/hls", "../hls")

	// Rutas API
	r.GET("/streams", func(c *gin.Context) {
		streamsMu.RLock()
		defer streamsMu.RUnlock()
		c.JSON(http.StatusOK, streams)
	})
	r.GET("/streams/:id", func(c *gin.Context) {
		id := c.Param("id")
		streamsMu.RLock()
		defer streamsMu.RUnlock()
		for _, s := range streams {
			if s.ID == id {
				c.JSON(http.StatusOK, s)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "stream not found"})
	})
	r.POST("/streams", func(c *gin.Context) {
		var s Stream
		if err := c.ShouldBindJSON(&s); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		streamsMu.Lock()
		streams = append(streams, s)
		streamsMu.Unlock()
		c.JSON(http.StatusCreated, s)
	})

	r.Run(":8080")
}
