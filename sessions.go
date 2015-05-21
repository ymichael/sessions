package sessions

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

/**
 * Configuration Options for Sessions.
 */
type SessionOptions struct {
	Name          string
	Secret        string
	EnvKey        string
	Store         Store
	CookieOptions *CookieOptions
}

/**
 * Helper to create SessionOptions object with sensible defaults.
 */
func NewSessionOptions(secret string, store Store) *SessionOptions {
	return &SessionOptions{
		Name:          "gojisid",
		EnvKey:        "session",
		CookieOptions: &CookieOptions{"/", 0, true, false},
		Secret:        secret,
		Store:         store,
	}
}

/**
 * Get session object from context
 */
func (s SessionOptions) GetSession(c *web.C) map[string]interface{} {
	return c.Env[s.EnvKey].(map[string]interface{})
}

/**
 * Returns session middleware.
 */
func (s SessionOptions) Middleware() web.MiddlewareType {
	middlewareFn := func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if c.Env == nil {
				c.Env = make(map[interface{}]interface{})
			}

			var sessionId string
			var sessionObj map[string]interface{}
			// If cookie is set, retrieve session and set on context.
			// Otherwise, create new session.
			cookie, err := r.Cookie(s.Name)
			if err == http.ErrNoCookie {
				sessionId = GenerateRandomString(24)
				sessionObj = make(map[string]interface{})
			} else {
				sessionId = cookie.Value
				sessionObj, err = s.Store.Get(sessionId)
				if err != nil {
					sessionObj = make(map[string]interface{})
				}
			}

			// Set new cookie in response.
			cookie = NewCookie(s.Name, sessionId, s.CookieOptions)
			http.SetCookie(w, cookie)

			// Add session object to context.
			c.Env[s.EnvKey] = sessionObj
			h.ServeHTTP(w, r)
			// Persist session to store.
			s.Store.Save(sessionId, sessionObj)
		}
		return http.HandlerFunc(fn)
	}
	return middlewareFn
}
