package repo

//import (
//	"gorm.io/gorm"
//	"rest_api/internal/app/adapters/db/gorm/models"
//	album_domain "rest_api/internal/app/domain/album"
//	"rest_api/pkg/client/postgres"
//)
//
//func init() {
//	db, err := postgres.GetPostgresEngine()
//	if err != nil {
//		panic(err)
//	}
//	repo := AlbumRepository{db: db}
//	err = repo.Migrate()
//	if err != nil {
//		panic(err)
//	}
//}
//
//func NewAlbumRepository(db *gorm.DB) *AlbumRepository {
//	return &AlbumRepository{
//		db: db,
//	}
//}
//
//type AlbumRepository struct {
//	db *gorm.DB
//}
//
//func (repo *AlbumRepository) Migrate() error {
//	return repo.db.AutoMigrate(&models.Album{})
//}
//
//func (repo *AlbumRepository) CreateAlbum(title, artist string, price float64) (album_domain.Album, error) {
//	album := models.Album{Title: title, Artist: artist, Price: price}
//	result := repo.db.Create(&album)
//	if result.Error != nil {
//		return album_domain.Album{}, result.Error
//	}
//	albumDomain := album_domain.Album{
//		ID:     album.ID,
//		Title:  album.Title,
//		Artist: album.Artist,
//		Price:  album.Price,
//	}
//	return albumDomain, nil
//}
//
//func (repo *AlbumRepository) GetAlbums() ([]album_domain.Album, error) {
//	var albums []album_domain.Album
//	if err := repo.db.Find(&albums).Error; err != nil {
//		return nil, err
//	}
//
//	return albums, nil
//}
//
//func (repo *AlbumRepository) DeleteAlbum(pk uint) (int64, error) {
//	deleteRes := repo.db.Unscoped().Delete(&models.Album{}, pk)
//	if err := deleteRes.Error; err != nil {
//		return 0, err
//	}
//
//	return deleteRes.RowsAffected, nil
//}
