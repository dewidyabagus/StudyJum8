package product

type service struct {
	repository     Repositorer
	userRepository UserRepositorer
}

func NewService(r Repositorer, u UserRepositorer) Servicer {
	return &service{
		repository:     r,
		userRepository: u,
	}
}

func (s *service) PublicFunc()
