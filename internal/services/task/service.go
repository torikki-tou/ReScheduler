package task

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Get() string {
	return "get task"
}

func (s *Service) Search() string {
	return "search task"
}

func (s *Service) Create() string {
	return "create task"
}

func (s *Service) Update() string {
	return "update task"
}

func (s *Service) Delete() string {
	return "delete task"
}
