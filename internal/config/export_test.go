package config

func WithCustomConfPath(path string) Option {
	return func(o *options) {
		if path != "" {
			o.addUserConfPath = path
		}
	}
}
