package url

type UrlService struct {
	UrlRepository UrlRepository
}

func ProvideUrlService(u UrlRepository) UrlService {
	return UrlService{UrlRepository: u}
}

func (u *UrlService) Create(url Url) Url {
	res := u.UrlRepository.Create(url)

	return res
}

func (u *UrlService) FindByID(id int) Url {
	url := u.UrlRepository.FindByID(id)

	return url
}
