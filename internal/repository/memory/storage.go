package memory

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func New() *MemStorage {
	return &MemStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func (s *MemStorage) UpdateGauge(name string, val float64) {
	s.gauge[name] = val
}

func (s *MemStorage) UpdateCounter(name string, val int64) {
	s.counter[name] += val
}
