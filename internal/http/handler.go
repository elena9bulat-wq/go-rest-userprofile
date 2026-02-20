package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/elena9bulat-wq/go-rest-userprofile/internal/service"
)

type Handler struct {
	profileService *service.ProfileService
}

func NewHandler(profileService *service.ProfileService) *Handler {
	return &Handler{profileService: profileService}
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.URL.Query().Get("userId")
	if userIDStr == "" {
		http.Error(w, "missing required query param: userId", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		http.Error(w, "invalid userId (must be a positive integer)", http.StatusBadRequest)
		return
	}

	profile, err := h.profileService.GetProfile(userID)
	if err != nil {
		http.Error(w, "profile not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}
