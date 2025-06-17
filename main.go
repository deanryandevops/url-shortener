// Main package declaration (similar to C# namespace)
package main

// Import statements (like C# using directives)
import (
	"encoding/json" // Like System.Text.Json in C#
	"fmt"           // Similar to System.Console or string formatting
	"log"           // Basic logging like ILogger
	"net/http"      // HTTP functionality like ASP.NET Core
	"sync"          // Concurrency primitives like System.Threading
)

// ShortURL struct - similar to a C# DTO or record
type ShortURL struct {
	Key string `json:"key"` // Field tags work like C# JsonProperty attributes
	URL string `json:"url"`
}

// URLStore class equivalent - our in-memory data store
type URLStore struct {
	urls map[string]string // Dictionary<string, string> in C#
	mu   sync.RWMutex      // Like ReaderWriterLockSlim in C#
}

// Global variables (similar to static fields in C#)
var (
	store = URLStore{urls: make(map[string]string)} // Initialize like new Dictionary<string,string>()
	port  = "8080"                                  // Like const string in C#
)

// Set method - similar to C# dictionary indexer but with thread safety
func (s *URLStore) Set(key, url string) {
	s.mu.Lock()         // Lock for writing (like lock() in C#)
	defer s.mu.Unlock() // Ensure unlock happens (like finally block)
	s.urls[key] = url   // Dictionary assignment
}

// Get method - similar to C#'s TryGetValue pattern
func (s *URLStore) Get(key string) (string, bool) {
	s.mu.RLock()               // Lock for reading (like ReaderLock in C#)
	defer s.mu.RUnlock()       // Ensure unlock
	url, exists := s.urls[key] // TryGetValue equivalent
	return url, exists
}

// Count method - property-like getter in C#
func (s *URLStore) Count() int {
	s.mu.RLock()         // Read lock
	defer s.mu.RUnlock() // Ensure unlock
	return len(s.urls)   // Like Count property in C#
}

// shortenHandler - similar to ASP.NET Core API endpoint
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	// Check HTTP method (like [HttpPost] attribute)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return // Early return like in C#
	}

	// Request DTO (similar to [FromBody] in C#)
	var request struct {
		URL string `json:"url"`           // Required field
		Key string `json:"key,omitempty"` // Optional field
	}

	// JSON deserialization (like JsonSerializer.Deserialize)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Default value logic (like ?? operator in C#)
	if request.Key == "" {
		request.Key = generateKey(store.Count())
	}

	// Store the URL (like adding to Dictionary)
	store.Set(request.Key, request.URL)

	// Return response (like return Ok() in C#)
	respondWithJSON(w, http.StatusCreated, ShortURL{
		Key: request.Key,
		URL: request.URL,
	})
}

// generateKey - extracted method similar to C# helper method
func generateKey(sequence int) string {
	return fmt.Sprintf("%x", sequence+1) // String formatting like string.Format
}

// redirectHandler - similar to redirect action in ASP.NET Core
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:] // Get path parameter (like [FromRoute])

	// Lookup URL (like dictionary lookup)
	if url, exists := store.Get(key); exists {
		http.Redirect(w, r, url, http.StatusMovedPermanently) // Like RedirectResult
		return
	}

	http.NotFound(w, r) // Like return NotFound()
}

// metricsHandler - simple endpoint like a health check
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")    // Set response header
	fmt.Fprintf(w, "Total URLs: %d", store.Count()) // Like string interpolation
}

// respondWithJSON - helper method similar to C# ControllerBase helpers
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json") // Set content type
	w.WriteHeader(statusCode)                          // Set status code (like StatusCodeResult)
	json.NewEncoder(w).Encode(data)                    // JSON serialization
}

// Main function - similar to C# Program.cs entry point
func main() {
	// Route setup (like app.MapControllers in C#)
	http.HandleFunc("/shorten", shortenHandler) // POST endpoint
	http.HandleFunc("/metrics", metricsHandler) // GET endpoint
	http.HandleFunc("/", redirectHandler)       // Catch-all route

	// Start server (similar to app.Run() in C#)
	log.Printf("Server starting on port %s...", port) // Logging
	log.Fatal(http.ListenAndServe(":"+port, nil))     // Start listening
}
