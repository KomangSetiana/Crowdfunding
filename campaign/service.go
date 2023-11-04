package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {

	if userID != 0 {
		campaigs, err := s.repository.GetByUserID(userID)

		if err != nil {
			return campaigs, err
		}
		return campaigs, nil
	}
	campaigs, err := s.repository.GetAll()

	if err != nil {
		return campaigs, err
	}
	return campaigs, nil

}
