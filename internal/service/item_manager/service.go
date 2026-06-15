package item_manager

type ItemService struct {
	repo repository
}

func New(repo repository) *ItemService {
	return &ItemService{repo: repo}
}
