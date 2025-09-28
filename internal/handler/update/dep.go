package update

type Storage interface {
	UpdateGauge(name string, val float64)
	UpdateCounter(name string, val int64)
}
