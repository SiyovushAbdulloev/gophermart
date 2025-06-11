package postgres

type Option func(*Postgres)

func MaxPoolSize(maxPoolSize int) Option {
	return func(p *Postgres) {
		p.maxPoolSize = maxPoolSize
	}
}

func ConnAttempts(attempts int) Option {
	return func(p *Postgres) {
		p.connAttempts = attempts
	}
}

func ConnDelay(delay int) Option {
	return func(p *Postgres) {
		p.connDelay = delay
	}
}
