package item_manager

type ItemService struct {
	repo Repository
}

func New(repo Repository) *ItemService {
	return &ItemService{repo: repo}
}
