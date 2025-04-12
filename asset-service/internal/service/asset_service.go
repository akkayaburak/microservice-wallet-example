package service

import (
	"asset-service/internal/model"
	"asset-service/internal/repository"
)

type AssetService struct {
	repo *repository.AssetRepository
}

func NewAssetService(repo *repository.AssetRepository) *AssetService {
	return &AssetService{repo: repo}
}

func (s *AssetService) CreateAsset(asset model.Asset) (model.Asset, error) {
	err := s.repo.CreateAsset(asset)
	if err != nil {
		return model.Asset{}, err
	}
	return asset, nil
}

func (s *AssetService) GetAssets() ([]model.Asset, error) {
	assets, err := s.repo.GetAssets()
	if err != nil {
		return []model.Asset{}, err
	}
	return assets, nil
}

func (s *AssetService) UpdateAsset(asset *model.Asset) (*model.Asset, error) {
	asset, err := s.repo.UpdateAsset(asset)
	if err != nil {
		return &model.Asset{}, err
	}
	return asset, nil
}
