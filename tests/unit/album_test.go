package unit

import (
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	album_repo "rest_api/internal/app/adapters/db/gorm/repo"
	album2 "rest_api/internal/app/controller/http/album"
	"rest_api/internal/app/domain/album"
	"rest_api/pkg/client/postgres"
	"testing"
)

func setupDatabase() *gorm.DB {
	db, _ := postgres.GetPostgresEngine()
	return db
}

func setup() (*album_repo.AlbumRepository, *album.AlbumService, *gorm.DB) {
	db := setupDatabase()

	repo := album_repo.NewAlbumRepository(db)
	service := album.NewAlbumService(repo)

	return repo, service, db
}

func TestAlbumCreate(t *testing.T) {
	repo, _, _ := setup()
	inputAlbum := album2.CreateAlbumInput{Title: "Test Album", Artist: "Unknown", Price: &[]float64{1.52}[0]}
	newAlbum, _ := repo.CreateAlbum(inputAlbum.Title, inputAlbum.Artist, *inputAlbum.Price)
	assert2.NotNil(t, newAlbum)
	assert.Equal(t, inputAlbum.Title, newAlbum.Title)
	assert.Equal(t, inputAlbum.Artist, newAlbum.Artist)
	assert.Equal(t, inputAlbum.Price, newAlbum.Price)
}

func TestAlbumsGet(t *testing.T) {
	repo, _, _ := setup()
	inputAlbum := album2.CreateAlbumInput{Title: "Test Album", Artist: "Unknown", Price: &[]float64{1.52}[0]}
	newAlbum, _ := repo.CreateAlbum(inputAlbum.Title, inputAlbum.Artist, *inputAlbum.Price)
	albums, err := repo.GetAlbums()
	if err != nil {
		panic(err)
	}
	lastCreatedAlbum := albums[len(albums)-1]
	assert.Equal(t, lastCreatedAlbum.Title, newAlbum.Title)
	assert.Equal(t, lastCreatedAlbum.Price, newAlbum.Price)
	assert.Equal(t, lastCreatedAlbum.Artist, newAlbum.Artist)
}

func TestAlbumDelete(t *testing.T) {
	repo, _, _ := setup()
	inputAlbum := album2.CreateAlbumInput{Title: "Test Album", Artist: "Unknown", Price: &[]float64{1.52}[0]}
	_, _ = repo.CreateAlbum(inputAlbum.Title, inputAlbum.Artist, *inputAlbum.Price)
	albums, err := repo.GetAlbums()
	if err != nil {
		panic(err)
	}
	countAlbumsBeforeDelete := len(albums)
	result, err := repo.DeleteAlbum(albums[len(albums)-1].ID)
	if err != nil {
		panic(err)
	}
	assert2.NotNil(t, result)
	albums, err = repo.GetAlbums()
	if err != nil {
		panic(err)
	}
	assert.NotEqual(t, countAlbumsBeforeDelete, len(albums))
	assert.Equal(t, countAlbumsBeforeDelete-1, len(albums))

}

func TestAlbumService(t *testing.T) {
	_, service, _ := setup()
	inputAlbum := album2.CreateAlbumInput{Title: "Test Album", Artist: "Unknown", Price: &[]float64{1.52}[0]}
	newAlbum, _ := service.CreateAlbum("Test Album", "Unknown", 1.52)
	assert.Equal(t, inputAlbum.Title, newAlbum.Title)
	assert.Equal(t, inputAlbum.Artist, newAlbum.Artist)
	assert.Equal(t, inputAlbum.Price, newAlbum.Price)
}
