package service

import "clashbot/cr"

type Service struct {
	cr *cr.Client
}

func NewService(client *cr.Client) *Service {
	return &Service{cr: client}
}

func (s *Service) GetPlayer(tag string) (*cr.Player, error) {
	return s.cr.GetPlayer(tag)
}

func (s *Service) GetClan(tag string) (*cr.Clan, error) {
	return s.cr.GetClan(tag)
}
