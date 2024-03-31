package auth

import "net/http"

// SetToken sets the Authorization header in the request if not set already
func SetToken(req *http.Request, url string) error {
	if req.Header.Get("Authorization") != "" {
		return nil
	}
	token, err := BuildJWTForRESTAPI(url)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}
