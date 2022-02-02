package todo

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

func (repository *Repository) FindAll() []Todo {
	var todos []Todo
	repository.database.Find(&todos)
	return todos
}

func (repository *Repository) Find(id int) (Todo, error) {
	var todo Todo
	err := repository.database.Find(&todo, id).Error
	if todo.Name == "" {
		err = errors.New("Todo not found")
	}
	return todo, err
}

func (repository *Repository) Create(todo Todo) (Todo, error) {
	err := repository.database.Create(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (repository *Repository) Save(user Todo) (Todo, error) {
	err := repository.database.Save(user).Error
	return user, err
}

func (repository *Repository) Delete(id int) int64 {
	count := repository.database.Delete(&Todo{}, id).RowsAffected
	return count
}

func NewRepository(database *gorm.DB) *Repository {
	return &Repository{
		database: database,
	}
}
