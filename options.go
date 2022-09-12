package go_differ

type options struct {
	// 忽略的字段
	ignoreFields []string
}

type Option func(*options)

func WithIgnoreFields(fields ...string) Option {
	return func(o *options) {
		o.ignoreFields = fields
	}
}
