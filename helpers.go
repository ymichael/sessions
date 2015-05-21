package sessions

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

/**
 * Generates random string of length n
 */
func GenerateRandomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

/**
 * Reverse http.Cookie.String(),
 * Taken from: http://play.golang.org/p/YkW_z2CSyE
 */
func CookieFromString(line string) (*http.Cookie, error) {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", line))))
	if err != nil {
		return nil, err
	}
	cookies := req.Cookies()
	if len(cookies) == 0 {
		return nil, fmt.Errorf("no cookies")
	}
	return cookies[0], nil
}
