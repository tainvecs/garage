package newssvc

import "context"

// Delete ESNewsDoc in elasticsearch
func (dao *esDAO) Delete(ctx context.Context, docID string) error {
	return dao.DataAccessObject.Delete(ctx, docID)
}
