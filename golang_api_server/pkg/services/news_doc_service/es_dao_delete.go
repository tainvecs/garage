package news_doc_service

import "context"

func (dao *esDAO) Delete(ctx context.Context, docID string) error {
	return dao.Client.Delete(ctx, dao.IndexIndex, docID)
}
