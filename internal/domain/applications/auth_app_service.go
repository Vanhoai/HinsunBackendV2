package applications

type AuthAppService interface{}

type appAppService struct{}

func NewAuthAppService() AuthAppService {
	return &appAppService{}
}
