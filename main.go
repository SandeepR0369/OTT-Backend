package main

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string `json:"title"`
	Path  string `json:"path"`
}

func main() {
	movieFolder := "/Users/sandeepreddy/Movies" // Path to your movies folder

	r := gin.Default()

	// Endpoint to list all movies (only .mp4 files)
	r.GET("/movies", func(c *gin.Context) {
		var movies []Movie
		err := filepath.WalkDir(movieFolder, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// Include only .mp4 files
			if !d.IsDir() && isMP4File(path) {
				movies = append(movies, Movie{
					Title: filepath.Base(path),
					Path:  path,
				})
			}
			return nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read movie directory"})
			return
		}

		c.JSON(http.StatusOK, movies)
	})

	// Endpoint to stream a specific movie by file path
	r.GET("/movies/play", func(c *gin.Context) {
		moviePath := c.Query("path")
		if moviePath == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie path is required"})
			return
		}

		// Ensure the file is .mp4
		if !isMP4File(moviePath) {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Only .mp4 files are supported"})
			return
		}

		// Stream the file
		c.Header("Content-Type", "video/mp4")
		c.Header("Content-Disposition", "inline")
		c.File(moviePath)
	})

	// Start the server
	r.Run(":8080") // Server runs on http://localhost:8080
}

// isMP4File checks if a file is an .mp4 file
func isMP4File(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), ".mp4")
}
