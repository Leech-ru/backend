package cookie

type cookieService struct {
	accessCookieName  string
	refreshCookieName string
}

func NewService() *cookieService {
	return &cookieService{
		accessCookieName:  "user_auth_access_token",
		refreshCookieName: "user_auth_refresh_token",
	}
}
