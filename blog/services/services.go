package services

type Registerer interface {
	Register(name string, service *Service) bool
	Unregister(name string) bool
}

type Service struct {
	services map[string]*Service
}

func NewService() *Service {
	return &Service{
		services: make(map[string]*Service),
	}
}

func (self *Service) Register(name string, s *Service) bool {
	self.services[name] = s
	return true
}

func (self *Service) Unregister(name string) bool {
	delete(self.services, name)
	if _, ok := self.services[name]; ok {
		return !ok
	} else {
		return ok
	}
}
