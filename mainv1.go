// package main

// import (
// 	"fmt"
// 	"io/fs"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// type Movie struct {
// 	Title string `json:"title"`
// 	Path  string `json:"path"`
// }

// func main() {
// 	movieFolder := "/Users/sandeepreddy/Movies" // Path to your movies folder

// 	// Ensure correct permissions on the folder and its files
// 	if err := ensurePermissions(movieFolder); err != nil {
// 		log.Fatalf("Failed to set permissions: %v", err)
// 	}

// 	r := gin.Default()

// 	// Endpoint to list all movies
// 	r.GET("/movies", func(c *gin.Context) {
// 		var movies []Movie
// 		err := filepath.WalkDir(movieFolder, func(path string, d fs.DirEntry, err error) error {
// 			if err != nil {
// 				return err
// 			}
// 			// Include only movie files
// 			if !d.IsDir() && isMovieFile(path) {
// 				movies = append(movies, Movie{
// 					Title: filepath.Base(path),
// 					Path:  path,
// 				})
// 			}
// 			return nil
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read movie directory"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, movies)
// 	})

// 	// Endpoint to stream a specific movie by file path
// 	r.GET("/movies/play", func(c *gin.Context) {
// 		moviePath := c.Query("path")
// 		if moviePath == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie path is required"})
// 			return
// 		}

// 		// Determine the content type based on file extension
// 		ext := strings.ToLower(filepath.Ext(moviePath))
// 		contentType := ""
// 		switch ext {
// 		case ".mp4":
// 			contentType = "video/mp4"
// 		case ".mkv":
// 			contentType = "video/x-matroska"
// 		case ".avi":
// 			contentType = "video/x-msvideo"
// 		case ".mov":
// 			contentType = "video/quicktime"
// 		default:
// 			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported file type"})
// 			return
// 		}

// 		// Set the Content-Type header to indicate the file type
// 		c.Header("Content-Type", contentType)

// 		// Prevent Content-Disposition header from forcing a download
// 		c.Header("Content-Disposition", "inline")

// 		// Stream the video file
// 		c.File(moviePath)
// 	})

// 	// Start the server
// 	r.Run(":8080") // Server runs on http://localhost:8080
// }

// // ensurePermissions checks and sets permissions for the folder and its files
// func ensurePermissions(folder string) error {
// 	// Set folder permissions to 755
// 	if err := os.Chmod(folder, 0755); err != nil {
// 		return fmt.Errorf("failed to set folder permissions: %w", err)
// 	}

// 	// Walk through all files and set their permissions to 644
// 	return filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		// If it's a file, set permissions to 644
// 		if !d.IsDir() {
// 			if err := os.Chmod(path, 0644); err != nil {
// 				return fmt.Errorf("failed to set file permissions for %s: %w", path, err)
// 			}
// 		}
// 		return nil
// 	})
// }

// // isMovieFile checks if a file is a movie based on its extension
// func isMovieFile(path string) bool {
// 	extensions := []string{".mp4", ".mkv", ".avi", ".mov"} // Add more extensions as needed
// 	for _, ext := range extensions {
// 		if strings.HasSuffix(strings.ToLower(path), ext) {
// 			return true
// 		}
// 	}
// 	return false
// }
