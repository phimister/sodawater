package soda

type Option func(*Model)

func WithTPS(tps int) Option {
	if tps <= 0 {
		panic("[clockModel] TPS must be greater than 0")
	}
	return func(m *Model) {
		m.clock = clockModel{tps: tps}
	}
}
