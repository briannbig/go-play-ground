package shortener

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type UrlShortener struct {
	Urls map[string]string
}

func (u *UrlShortener) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Should be POST", http.StatusMethodNotAllowed)
		return
	}

	originalUrl := r.FormValue("url")

	if originalUrl == "" {
		http.Error(w, "URL Parameter is missing", http.StatusBadRequest)
		return
	}

	shortKey := generateShortKey()
	u.Urls[shortKey] = originalUrl

	shortenedURL := fmt.Sprintf("http://localhost:3000/short/%s", shortKey)

	w.Header().Set("Content-Type", "text/html")
	responseHTML := fmt.Sprintf(`
        <h2>URL Shortener</h2>
        <p>Original URL: %s</p>
        <p>Shortened URL: <a href="%s">%s</a></p>
        <form method="post" action="/shorten">
            <input type="text" name="url" placeholder="Enter a URL">
            <input type="submit" value="Shorten">
        </form>
    `, originalUrl, shortenedURL, shortenedURL)
	fmt.Fprintf(w, responseHTML)

}
func (us *UrlShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	originalURL, found := us.Urls[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
