package jsonx

type Option func(*config)

func WithJsonLibType(lib jsonLibType) Option {
	return func(c *config) {
		c.jsonLibType = lib
	}
}

func WithUseNumber(enable bool) Option {
	return func(c *config) {
		c.useNumber = enable
	}
}

func WithEscapeHTML(enable bool) Option {
	return func(c *config) {
		c.escapeHTML = enable
	}
}
