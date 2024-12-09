package repository

import (
	"context"
	"pikachu/model"

	"github.com/juju/errors"
	"gorm.io/gorm"
)

type gormCompanyReadOnlyRepository struct {
	conn *gorm.DB
}

// NewGormCompanyReadOnlyRepository ...
func NewGormCompanyReadOnlyRepository(conn *gorm.DB) CompanyReadOnlyRepository {
	return &gormCompanyReadOnlyRepository{conn: conn}
}

// GetCompanyByUserID ...
func (g *gormCompanyReadOnlyRepository) GetCompanyByUserID(ctx context.Context, uid string) (rcompany *model.Company, err error) {
	zlog.With(ctx).Infow("[New Repository Request]", "uid", uid)

	scope := g.conn.WithContext(ctx)
	scope = scope.Joins("JOIN user_companies ON companies.id = user_companies.company_id").
		Where("user_companies.user_id = ?", uid).
		First(&rcompany)

	if err = scope.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zlog.With(ctx).Errorw("GetCompanyByUserID Not Found", "uid", uid, "err", err)
			return nil, errors.NotFoundf("Company is not exist")
		}
		zlog.With(ctx).Errorw("GetCompanyByUserID Error", "uid", uid, "err", err)
		return nil, err
	}

	return rcompany, nil
}

// GetCompany ...
func (g *gormCompanyReadOnlyRepository) GetCompany(ctx context.Context, id uint64) (rcompany *model.Company, err error) {
	zlog.With(ctx).Infow("[New Repository Request]", "id", id)

	scope := g.conn.WithContext(ctx)
	scope = scope.First(&rcompany, id)

	if err = scope.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zlog.With(ctx).Errorw("GetCompany Not Found", "id", id, "err", err)
			return nil, errors.NotFoundf("Company is not exist")
		}
		zlog.With(ctx).Errorw("GetCompany Error", "id", id, "err", err)
		return nil, err
	}

	return rcompany, nil
}
