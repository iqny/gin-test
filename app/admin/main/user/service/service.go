package service

import "byobject/app/admin/main/user/conf"

type Service struct {
	c *conf.Config
}

func New(c *conf.Config) *Service {
	return &Service{}
}
func (s *Service) Close() error {

	return nil
}
