package news_doc_service

import "context"

func (dao *newsDocESDAO) Delete(ctx context.Context, docID string) error {
	return dao.Client.Delete(ctx, dao.IndexIndex, docID)
}
