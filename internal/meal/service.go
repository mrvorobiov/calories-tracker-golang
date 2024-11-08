package meal

import "errors"

type Service struct {
	repository *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repository: r}
}

func (s Service) Track(meal Meal) error {
	if meal.Calories <= 0 {
		return errors.New(ErrCaloriesAreNotPositiveInteger)
	}
	return s.repository.Create(meal)
}

func (s Service) GetAll() ([]Meal, error) {
	return s.repository.GetAll()
}

func (s Service) Delete(id uint) error {
	return s.repository.Delete(Meal{Id: id})
}
