package service

import "errors"

type UserProfile struct {
	UserID          int    `json:"userId"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	DeliveryAddress string `json:"deliveryAddress"`
	PhotoURL        string `json:"photoUrl"`
	CardLast4       string `json:"cardLast4"`
}

// ProfileService tine logica de business (aici mock/in-memory).
type ProfileService struct {
	profiles map[int]UserProfile
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		profiles: map[int]UserProfile{
			1: {
				UserID:          1,
				FirstName:       "Ana",
				LastName:        "Popescu",
				DeliveryAddress: "Str. Exemplu 10, Chisinau",
				PhotoURL:        "https://example.com/photo.jpg",
				CardLast4:       "1234",
			},
		},
	}
}

func (s *ProfileService) GetProfile(userID int) (UserProfile, error) {
	p, ok := s.profiles[userID]
	if !ok {
		return UserProfile{}, errors.New("profile not found")
	}
	return p, nil
}
