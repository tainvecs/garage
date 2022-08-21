package news_doc_service

import "context"

func (dao *ESDAO) Index(ctx context.Context, doc *NewsDoc) error {
	return dao.Client.Index(ctx, dao.IndexIndex, doc.ID, doc)
}
