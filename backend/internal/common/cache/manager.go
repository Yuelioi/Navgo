package cache

var Manager *_Manager

type _Manager struct {
	Controllers map[string]Controller
}

func New() *_Manager {
	m := &_Manager{}

	m.RegisterController(NewInReviewCache())
	return m
}

func (m *_Manager) GetController(name string) Controller {
	if c, ok := m.Controllers[name]; ok {
		return c
	}
	return nil
}

func (m *_Manager) RegisterController(controllers ...Controller) error {
	for _, c := range controllers {
		err := c.Preload()
		if err != nil {
			return err
		}
	}
	return nil
}
