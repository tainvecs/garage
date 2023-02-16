package newssvc

import (
	"context"

	"github.com/tainvecs/garage/apisrv/pkg/data_access/sqldao"
)

// GetAll is the service func for getting all news doc from psql
func (dao *psqlDAO) Get(ctx context.Context, queryConf *sqldao.QueryConfig) ([]*PsqlNewsDoc, error) {

	var docSlice []*PsqlNewsDoc

	err := queryConf.
		Apply(dao.Client).
		WithContext(ctx).
		Model(PsqlNewsDoc{}).
		Order("id").
		Find(&docSlice).Error

	return docSlice, err
}
