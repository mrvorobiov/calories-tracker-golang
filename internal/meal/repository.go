package meal

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetAll() ([]Meal, error) {
	var meals []Meal
	res := r.db.Find(&meals)
	return meals, res.Error
}

func (r Repository) Create(meal Meal) error {
	return r.db.Create(meal).Error
}

func (r Repository) Delete(meal Meal) error {
	return r.db.Delete(meal).Error
}

func (r Repository) Update(meal Meal) error {
	return r.db.Save(meal).Error
}
