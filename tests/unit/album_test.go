package unit

import (
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"rest_api/app"
	"rest_api/app/album"
	"testing"
)

func setupDatabase() *gorm.DB {
	db, _ := app.GetDBEngine()
	return db
}

func setup() (*album.AlbumRepository, *album.AlbumService, *gorm.DB) {
	db := setupDatabase()

	repo := album.NewAlbumRepository(db)
	service := album.NewAlbumService(repo)

	return repo, service, db
}

func TestAlbumCreate(t *testing.T) {
	repo, _, _ := setup()
	inputAlbum := album.CreateAlbumInput{Title: "Test Album", Artist: "Unknown", Price: &[]float64{1.52}[0]}
	newAlbum, _ := repo.CreateAlbum(inputAlbum.Title, inputAlbum.Artist, *inputAlbum.Price)
	assert2.NotNil(t, newAlbum)
	assert.Equal(t, inputAlbum.Title, newAlbum.Title)
	assert.Equal(t, inputAlbum.Artist, newAlbum.Artist)
	assert.Equal(t, inputAlbum.Price, newAlbum.Price)
}

func TestAlbumsGet(t *testing.T) {
	repo, _, _ := setup()
	inputAlbum := album.CreateAlbumInput{Title: "Test Album", Artist: "Unknown", Price: &[]float64{1.52}[0]}
	newAlbum, _ := repo.CreateAlbum(inputAlbum.Title, inputAlbum.Artist, *inputAlbum.Price)
	albums, err := repo.GetAlbums()
	if err != nil {
		panic(err)
	}
	lastCreatedAlbum := albums[len(albums)-1]
	assert.Equal(t, lastCreatedAlbum.ID, newAlbum.ID)
	assert.Equal(t, lastCreatedAlbum.Title, newAlbum.Title)
	assert.Equal(t, lastCreatedAlbum.Price, newAlbum.Price)
	assert.Equal(t, lastCreatedAlbum.Artist, newAlbum.Artist)
}

func TestAlbumDelete(t *testing.T) {
	repo, _, _ := setup()
	albums, err := repo.GetAlbums()
	if err != nil {
		panic(err)
	}
	countAlbumsBeforeDelete := len(albums)
	result, _ := repo.DeleteAlbum(albums[len(albums)-1].ID)
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
	inputAlbum := album.CreateAlbumInput{Title: "Test Album", Artist: "Unknown", Price: &[]float64{1.52}[0]}
	newAlbum, _ := service.CreateAlbum("Test Album", "Unknown", 1.52)
	assert.Equal(t, inputAlbum.Title, newAlbum.Title)
	assert.Equal(t, inputAlbum.Artist, newAlbum.Artist)
	assert.Equal(t, inputAlbum.Price, newAlbum.Price)
}