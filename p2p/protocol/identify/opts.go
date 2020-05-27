package identify

type config struct {
	userAgent string
	groupid   string // add by liangc
}

// Option is an option function for identify.
type Option func(*config)

// UserAgent sets the user agent this node will identify itself with to peers.
func UserAgent(ua string) Option {
	return func(cfg *config) {
		cfg.userAgent = ua
	}
}

// add by liangc
func Groupid(gid string) Option {
	return func(cfg *config) {
		cfg.groupid = gid
	}
}
