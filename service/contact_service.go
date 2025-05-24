package service

import (
	"addressbook/models"
	"addressbook/storage"
)

type ContactService struct {
	store *storage.InMemoryStorage
}

func NewContactService(store *storage.InMemoryStorage) *ContactService {
	return &ContactService{store}
}

func (s *ContactService) Create(contactsReq models.CreateContactsRequest) models.CreateContactsResponse {
	response := make(models.CreateContactsResponse, 0, len(contactsReq))
	for _, contactReq := range contactsReq {
		contact := s.store.Create(contactReq)
		response = append(response, contact)
	}

	return response
}

func (s *ContactService) Update(updateReq models.UpdateContactsRequest) []models.Contact {
	response := make(models.UpdateContactsResponse, 0, len(updateReq))
	for _, update := range updateReq {
		if contact, ok := s.store.Update(update); ok {
			response = append(response, contact)
		}
	}

	return response
}

func (s *ContactService) Delete(ids []string) int {
	deleted := 0
	for _, id := range ids {
		if s.store.Delete(id) {
			deleted++
		}
	}

	return deleted
}

func (s *ContactService) Search(searchReq models.SearchRequest) []models.Contact {
	return s.store.Search(searchReq.Query)
}
