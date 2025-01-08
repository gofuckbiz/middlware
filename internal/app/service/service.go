package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Days() int64 {

	data := time.Date(2029, time.April, 1, 0, 0, 0, 0, time.UTC)

	duration := time.Until(data)

	return int64(duration.Hours()) / 24
}
