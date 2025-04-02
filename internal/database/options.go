package database

type options struct {
	Driver string
	DSN    string
	User   string
	Pwd    string
	Host   string
	Port   int
	DB     string
}

type Option func(*options)

func WithDriver(driver string) Option {
	return func(opts *options) {
		opts.Driver = driver
	}
}

func WithDSN(dsn string) Option {
	return func(opts *options) {
		opts.DSN = dsn
	}
}

func WithUser(user string) Option {
	return func(opts *options) {
		opts.User = user
	}
}

func WithPwd(pwd string) Option {
	return func(opts *options) {
		opts.Pwd = pwd
	}
}

func WithHost(host string) Option {
	return func(opts *options) {
		opts.Host = host
	}
}

func WithPort(port int) Option {
	return func(opts *options) {
		opts.Port = port
	}
}

func WithDB(db string) Option {
	return func(opts *options) {
		opts.DB = db
	}
}

func NewOptions(opts ...Option) *options {
	opt := &options{}
	for _, o := range opts {
		o(opt)
	}
	return opt
}
