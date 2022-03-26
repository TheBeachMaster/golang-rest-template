package delivery

import "com.thebeachmaster/golangrest/internal/user"

type Delivery struct {
	User interface{ user.UserHTTPRoutes }
}
