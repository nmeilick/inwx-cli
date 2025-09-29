package api

import (
	"net/http"
	"sync"
)

type Session struct {
	cookies []*http.Cookie
	mutex   sync.RWMutex
}

func NewSession() *Session {
	return &Session{
		cookies: make([]*http.Cookie, 0),
	}
}

func (s *Session) StoreCookies(cookies []*http.Cookie) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, newCookie := range cookies {
		found := false
		for i, existingCookie := range s.cookies {
			if existingCookie.Name == newCookie.Name {
				s.cookies[i] = newCookie
				found = true
				break
			}
		}
		if !found {
			s.cookies = append(s.cookies, newCookie)
		}
	}
}

func (s *Session) GetCookies() []*http.Cookie {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	result := make([]*http.Cookie, len(s.cookies))
	copy(result, s.cookies)
	return result
}

func (s *Session) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.cookies = make([]*http.Cookie, 0)
}
