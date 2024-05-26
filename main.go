package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/brpaz/echozap"
	"github.com/google/go-github/v39/github"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// Global map to store unique links and their expiry times
var uniqueLinks = make(map[string]time.Time)

// Function to generate a unique link
func generateUniqueLink() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Set up logger with zap
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	e.Use(echozap.ZapLogger(logger))

	// Middleware
	e.Use(middleware.Recover())

	// GitHub authentication setup
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "YOUR_GITHUB_ACCESS_TOKEN"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Route to generate unique links
	e.POST("/generate", func(c echo.Context) error {
		link := generateUniqueLink()
		expiry := time.Now().Add(24 * time.Hour) // Link expires in 24 hours
		uniqueLinks[link] = expiry
		return c.String(http.StatusOK, "http://localhost:8080/view/"+link)
	})

	// Proxy handler
	e.GET("/view/:link/*", func(c echo.Context) error {
		link := c.Param("link")
		expiry, exists := uniqueLinks[link]
		if !exists || time.Now().After(expiry) {
			return echo.ErrNotFound
		}

		// Extract the GitHub URL from the request
		path := strings.TrimPrefix(c.Param("*"), "/")
		target := "https://api.github.com/" + path
		targetURL, err := url.Parse(target)
		if err != nil {
			return err
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			c.Logger().Error("Error in proxy", zap.Error(err))
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		// Add GitHub authentication token to the request
		proxy.ModifyResponse = func(r *http.Response) error {
			r.Header.Set("Authorization", "token YOUR_GITHUB_ACCESS_TOKEN")
			return nil
		}

		proxy.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
