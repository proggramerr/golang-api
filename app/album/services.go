package album

type AlbumService struct {
	repository AlbumRepositoryImpl
}

func NewAlbumService(repository AlbumRepositoryImpl) *AlbumService {
	return &AlbumService{repository: repository}
}

func (service *AlbumService) CreateAlbum(title, artist string, price float64) (Album, error) {
	return service.repository.CreateAlbum(title, artist, price)
}

func (service *AlbumService) GetAlbums() ([]Album, error) {
	return service.repository.GetAlbums()
}

func (service *AlbumService) DeleteAlbum(pk uint) (int64, error) {
	return service.repository.DeleteAlbum(pk)
}
