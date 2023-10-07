package album

import (
	"gorm.io/gorm"
	"rest_api/app"
)

func NewAlbumRepository(db *gorm.DB) *AlbumRepository {
	return &AlbumRepository{
		db: db,
	}
}

type AlbumRepositoryImpl interface {
	Migrate() error
	CreateAlbum(title, artist string, price float64) (Album, error)
	GetAlbums() ([]Album, error)
	DeleteAlbum(pk uint) (int64, error)
}

type AlbumRepository struct {
	db *gorm.DB
}

func init() {
	db, err := app.GetDBEngine()
	if err != nil {
		panic(err)
	}
	repo := AlbumRepository{db: db}
	err = repo.Migrate()
	if err != nil {
		panic(err)
	}
}

func (repo *AlbumRepository) Migrate() error {
	return repo.db.AutoMigrate(&Album{})
}

func (repo *AlbumRepository) CreateAlbum(title, artist string, price float64) (Album, error) {
	album := Album{Title: title, Artist: artist, Price: price}
	result := repo.db.Create(&album)
	if result.Error != nil {
		return album, result.Error
	}
	return album, nil
}

func (repo *AlbumRepository) GetAlbums() ([]Album, error) {
	var albums []Album
	if err := repo.db.Find(&albums).Error; err != nil {
		return nil, err
	}

	return albums, nil
}

func (repo *AlbumRepository) DeleteAlbum(pk uint) (int64, error) {
	deleteRes := repo.db.Delete(&Album{}, pk)
	if err := deleteRes.Error; err != nil {
		return 0, err
	}

	return deleteRes.RowsAffected, nil
}
