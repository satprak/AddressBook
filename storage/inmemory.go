package storage

import (
	"addressbook/models"
	"addressbook/utils"
	"strings"
	"sync"
	"unicode"
)

type InMemoryStorage struct {
	contacts  map[string]models.Contact
	nameIndex map[string]map[string]struct{}
	mu        sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		contacts:  make(map[string]models.Contact),
		nameIndex: make(map[string]map[string]struct{}),
	}
}

func tokenize(text string) []string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	tokens := strings.FieldsFunc(text, f)
	for i, t := range tokens {
		tokens[i] = strings.ToLower(t)
	}
	return tokens
}

func (s *InMemoryStorage) addToIndex(id string, contact models.Contact) {
	for _, token := range tokenize(contact.Name) {
		if s.nameIndex[token] == nil {
			s.nameIndex[token] = make(map[string]struct{})
		}
		s.nameIndex[token][id] = struct{}{}
	}
}

func (s *InMemoryStorage) removeFromIndex(id string, contact models.Contact) {
	for _, token := range tokenize(contact.Name) {
		delete(s.nameIndex[token], id)
		if len(s.nameIndex[token]) == 0 {
			delete(s.nameIndex, token)
		}
	}

}

func (s *InMemoryStorage) Create(contact models.ContactRequest) models.Contact {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := utils.GenerateUUID()
	newContact := models.Contact{
		ID:    id,
		Name:  contact.Name,
		Phone: contact.Phone,
		Email: contact.Email,
	}
	s.contacts[id] = newContact
	s.addToIndex(id, newContact)
	return newContact
}

func (s *InMemoryStorage) Update(update models.UpdateRequest) (models.Contact, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	contact, exists := s.contacts[update.ID]
	if !exists {
		return models.Contact{}, false
	}

	s.removeFromIndex(update.ID, contact)

	if update.Name != "" {
		contact.Name = update.Name
	}
	if update.Phone != "" {
		contact.Phone = update.Phone
	}
	if update.Email != "" {
		contact.Email = update.Email
	}

	s.contacts[update.ID] = contact
	s.addToIndex(update.ID, contact)
	return contact, true
}

func (s *InMemoryStorage) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	contact, exists := s.contacts[id]
	if !exists {
		return false
	}

	s.removeFromIndex(id, contact)
	delete(s.contacts, id)
	return true
}

func (s *InMemoryStorage) Search(query string) []models.Contact {
	s.mu.RLock()
	defer s.mu.RUnlock()

	queryTokens := tokenize(query)
	if len(queryTokens) == 0 {
		return []models.Contact{}
	}

	var resultIDs map[string]struct{}
	for _, token := range queryTokens {
		ids := s.nameIndex[token]
		if ids == nil {
			return []models.Contact{}
		}
		if resultIDs == nil {
			resultIDs = make(map[string]struct{})
			for id := range ids {
				resultIDs[id] = struct{}{}
			}
		} else {
			for id := range resultIDs {
				if _, exists := ids[id]; !exists {
					delete(resultIDs, id)
				}
			}
			if len(resultIDs) == 0 {
				return []models.Contact{}
			}
		}
	}

	var results []models.Contact
	for id := range resultIDs {
		results = append(results, s.contacts[id])
	}
	return results
}
