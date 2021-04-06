package campaign

type Service interface {
	FindCampaign(userId int) ([]Campaign, error)
}

type service struct {
	repositiory Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaign(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaign, err := s.repositiory.FindById(userId)
		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}

	campaign, err := s.repositiory.FindAll()
	if err != nil {
		return campaign, err
	}

	return campaign, nil

}