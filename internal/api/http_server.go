package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"olinker/internal/core"
)

type Server struct {
	port    int
	service *core.EncodeService
}

func NewServer(port int, service *core.EncodeService) *Server {
	return &Server{
		port:    port,
		service: service,
	}
}

// enable CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /{vendor}/write_card", s.handleEncode)
	mux.HandleFunc("POST /{vendor}/cancel_card", s.handleCancel)
	mux.HandleFunc("POST /{vendor}/extend_card", s.handleExtend)
	
	// Legacy Main.cs used POST for read_card, but GET is also fine if needed. We register POST to match:
	mux.HandleFunc("POST /{vendor}/read_card", s.handleStatus)
	mux.HandleFunc("GET /{vendor}/read_card", s.handleStatus)

	mux.HandleFunc("GET /config", s.handleGetConfig)
	mux.HandleFunc("POST /config", s.handleSaveConfig)

	// Web UI
	mux.Handle("/", http.FileServer(http.Dir("web")))

	addr := fmt.Sprintf(":%d", s.port)
	log.Printf("Starting HTTP API Gateway on http://localhost%s\n", addr)
	return http.ListenAndServe(addr, corsMiddleware(mux))
}

func (s *Server) handleEncode(w http.ResponseWriter, r *http.Request) {
	var req core.EncodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vendorName := r.PathValue("vendor")
	log.Printf("[API] Processing encode card request for room: %s (Endpoint target vendor: %s)", req.RoomName, vendorName)
	res, err := s.service.EncodeCard(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (s *Server) handleCancel(w http.ResponseWriter, r *http.Request) {
	var req core.CancelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vendorName := r.PathValue("vendor")
	log.Printf("[API] Processing cancel card request for card: %s (vendor: %s)", req.CardID, vendorName)
	err := s.service.CancelCard(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": 0, "message": "Success"})
}

func (s *Server) handleExtend(w http.ResponseWriter, r *http.Request) {
	var req core.ExtendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vendorName := r.PathValue("vendor")
	log.Printf("[API] Processing extend card request for card: %s (vendor: %s)", req.CardID, vendorName)
	err := s.service.ExtendCard(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": 0, "message": "Success"})
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	vendorName := r.PathValue("vendor")
	log.Printf("[API] Processing read card request (vendor: %s)", vendorName)

	res, err := s.service.ReadCard()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (s *Server) handleGetConfig(w http.ResponseWriter, r *http.Request) {
	configData, err := os.ReadFile("configs/config.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(configData)
}

func (s *Server) handleSaveConfig(w http.ResponseWriter, r *http.Request) {
	var config core.VendorConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile("configs/config.json", configData, 0644); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
