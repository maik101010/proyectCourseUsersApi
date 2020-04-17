package service

var (
	//ItemService Interface service user
	ItemService itemsService = itemsService{}
)

type intemsServiceInterface interface {
	GetItem()
	SaveItem()
}
type itemsService struct {
}

func (s *itemsService) GetItem() {

}
func (s *itemsService) SaveItem() {

}
