package components

type Services []ServiceProvider

type ServiceProvider interface {
	InitService(u *Utils) error
	ShutdownService() error
}

type Service struct {
	*Utils
	onInit     func() error
	onShutdown func() error
}

func (s *Service) InitService(utils *Utils) error {
	s.Utils = utils
	if s.onInit != nil {
		return s.onInit()
	}
	return nil
}

func (s *Service) ShutdownService() error {
	if s.onShutdown != nil {
		return s.onShutdown()
	}
	return nil
}

func (s *Service) OnInit(callback func() error) {
	s.onInit = callback
}

func (s *Service) OnShutdown(callback func() error) {
	s.onShutdown = callback
}
