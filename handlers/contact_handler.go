package handlers

import (
	"addressbook/models"
	"addressbook/service"
	"encoding/json"
	"net/http"
)

type ContactHandler struct {
	service *service.ContactService
}

func NewContactHandler(s *service.ContactService) *ContactHandler {
	return &ContactHandler{s}
}

func (h *ContactHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var contactsReq models.CreateContactsRequest
	if err := json.NewDecoder(r.Body).Decode(&contactsReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := h.service.Create(contactsReq)

	writeJSON(w, response)
}

func (h *ContactHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatesReq models.UpdateContactsRequest
	if err := json.NewDecoder(r.Body).Decode(&updatesReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := h.service.Update(updatesReq)
	writeJSON(w, response)
}

func (h *ContactHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var ids models.DeleteContactsRequest
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	deletedCount := h.service.Delete(ids)

	writeJSON(w, map[string]int{"deleted": deletedCount})
}

func (h *ContactHandler) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var searchReq models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&searchReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := h.service.Search(searchReq)
	writeJSON(w, response)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
