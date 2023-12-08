package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	CreateCampaignImage(input CampaignImageInput, FileLocation string) (CampaignImage, error)
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

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.GetCampaignByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil

}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	//pembuatan slug
	slugCanidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCanidate) // Nama Campaign
	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil

}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.GetCampaignByID(inputID.ID)

	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount
	campaign.UserID = inputData.User.ID

	updatedCampaign, err := s.repository.Update(campaign)

	if err != nil {
		return campaign, err
	}
	return updatedCampaign, nil

}

func (s *service) CreateCampaignImage(input CampaignImageInput, fileLocation string) (CampaignImage, error) {

	campaign, err := s.repository.GetCampaignByID(input.CampaignID)

	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary == true {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesNonPrimary(input.CampaignID)

		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}

	campaignImage.CampaignID = input.CampaignID

	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)

	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
