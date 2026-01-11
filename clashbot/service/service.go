package service

import "clashbot/cr"

type Service struct {
	Cr *cr.Client
}

func NewService(client *cr.Client) *Service {
	return &Service{Cr: client}
}

func (s *Service) GetPlayer(tag string) (*cr.Player, error) {
	return s.Cr.GetPlayer(tag)
}

func (s *Service) GetClan(tag string) (*cr.Clan, error) {
	return s.Cr.GetClan(tag)
}

func (s *Service) GetBattleLog(tag string) (*cr.BattleList, error) {
	return s.Cr.GetBattleLog(tag)
}
