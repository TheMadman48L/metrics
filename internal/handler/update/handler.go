package update

import (
	"net/http"
	"strconv"
)

type Handler struct {
	repo Storage
}

func New(repo Storage) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	metricName := r.PathValue("name")
	metricValue := r.PathValue("value")
	metricType := r.PathValue("type")

	if metricName == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	switch metricType {
	case "gauge":
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		h.repo.UpdateGauge(metricName, value)
	case "counter":
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		h.repo.UpdateCounter(metricName, value)
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

}
