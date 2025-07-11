package e2e

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/mountain-reverie/honk4greenhouse/internal/server"
	playwrightci "github.com/mountain-reverie/playwright-ci-go"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var browser playwright.Browser

func TestMain(m *testing.M) {
	// Install Playwright dependencies
	err := playwrightci.Install(playwrightci.WithTimeout(2 * time.Minute))
	if err != nil {
		panic(fmt.Sprintf("Failed to install Playwright: %v", err))
	}

	// Launch browser once for all tests
	browser, err = playwrightci.Chromium()
	if err != nil {
		playwrightci.Uninstall()
		panic(fmt.Sprintf("Failed to launch browser: %v", err))
	}

	// Run tests
	code := m.Run()

	// Cleanup
	if browser != nil {
		browser.Close()
	}
	playwrightci.Uninstall()

	os.Exit(code)
}

func startTestServer(t *testing.T) (string, func()) {
	// Start the server on a free port
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port

	// Create and start the server using the existing listener
	r := server.New()
	srv := &http.Server{
		Handler: r,
	}

	go func() {
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			t.Logf("Server error: %v", err)
		}
	}()

	// Wait for server to be ready
	baseURL := fmt.Sprintf("http://localhost:%d", port)
	err = playwrightci.Wait4Port(baseURL)
	require.NoError(t, err)

	// Return cleanup function
	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}

	return baseURL, cleanup
}

func TestHome_E2E_BrowserInteraction(t *testing.T) {
	t.Parallel()
	baseURL, cleanup := startTestServer(t)
	defer cleanup()

	// Create a new page for this test
	page, err := browser.NewPage()
	require.NoError(t, err)
	defer page.Close()

	// Navigate to the home page
	response, err := page.Goto(baseURL)
	require.NoError(t, err)
	assert.Equal(t, 200, response.Status())

	// Check the page title
	title, err := page.Title()
	require.NoError(t, err)
	assert.Equal(t, "Honk4Greenhouse", title)

	// Check that the main heading is present
	heading, err := page.Locator("h1").TextContent()
	require.NoError(t, err)
	assert.Contains(t, heading, "Honk4Greenhouse")

	// Check that the welcome message is present
	welcomeMessage, err := page.Locator("p").TextContent()
	require.NoError(t, err)
	assert.Equal(t, "Hello World! Welcome to honk4greenhouse - your greenhouse design companion", welcomeMessage)

	// Verify the page structure
	container, err := page.Locator(".container").Count()
	require.NoError(t, err)
	assert.Equal(t, 1, container)

	// Verify the viewport meta tag for responsive design
	viewport, err := page.Locator("meta[name='viewport']").GetAttribute("content")
	require.NoError(t, err)
	assert.Equal(t, "width=device-width, initial-scale=1.0", viewport)

	// Check that the charset is UTF-8
	charset, err := page.Locator("meta[charset]").GetAttribute("charset")
	require.NoError(t, err)
	assert.Equal(t, "UTF-8", charset)
}

func TestHome_E2E_ResponseTime(t *testing.T) {
	t.Parallel()
	baseURL, cleanup := startTestServer(t)
	defer cleanup()

	// Create a new page for this test
	page, err := browser.NewPage()
	require.NoError(t, err)
	defer page.Close()

	startTime := time.Now()

	// Navigate to the home page
	response, err := page.Goto(baseURL)
	require.NoError(t, err)

	loadTime := time.Since(startTime)

	// Verify successful response
	assert.Equal(t, 200, response.Status())

	// Check that the page loads reasonably fast (under 1 second)
	assert.Less(t, loadTime, 1*time.Second, "Page should load within 1 second")

	// Wait for the page to be fully loaded
	err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})
	require.NoError(t, err)

	// Verify content is present after full load
	heading, err := page.Locator("h1").TextContent()
	require.NoError(t, err)
	assert.Contains(t, heading, "Honk4Greenhouse")
}

func TestHome_E2E_HealthEndpoint(t *testing.T) {
	t.Parallel()
	baseURL, cleanup := startTestServer(t)
	defer cleanup()

	// Create a new page for this test
	page, err := browser.NewPage()
	require.NoError(t, err)
	defer page.Close()

	// Navigate to the health endpoint
	response, err := page.Goto(baseURL + "/health")
	require.NoError(t, err)
	assert.Equal(t, 200, response.Status())

	// Check that we get JSON response
	contentType := response.Headers()["content-type"]
	assert.Contains(t, contentType, "application/json")

	// Get the response text (JSON)
	content, err := page.TextContent("body")
	require.NoError(t, err)
	assert.Contains(t, content, `"status":"ok"`)
	assert.Contains(t, content, `"service":"honk4greenhouse"`)
}