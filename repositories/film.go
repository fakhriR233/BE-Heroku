package repositories

import (
	"dumbflix_be/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilms() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film, ID int) (models.Film, error)
	DeleteFilm(film models.Film, ID int) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilms() ([]models.Film, error) {
	var films []models.Film
	err := r.db.Preload("Category").Find(&films).Error

	return films, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
var film models.Film
err := r.db.Preload("Category").First(&film, ID).Error

return film, err
}

func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Create(&film).Error

	return film, err
}

func (r *repository) UpdateFilm(film models.Film, ID int) (models.Film, error) {
	err := r.db.Preload("Category").Raw("UPDATE films SET title=?, thumbnail_film=?, year=?,category_id=?, description=? WHERE id=?", film.Title,film.ThumbnailFilm, film.Year, film.CategoryID, film.Description, ID).Scan(&film).Error

	return film, err
}

func (r *repository) DeleteFilm(film models.Film, ID int) (models.Film, error) {
	err := r.db.Delete(&film).Error

	return film, err
}