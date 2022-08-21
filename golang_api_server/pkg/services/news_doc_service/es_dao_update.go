package news_doc_service

import "context"

func (dao *ESDAO) Update(ctx context.Context, doc *NewsDoc) error {
	return dao.Client.Update(ctx, dao.IndexIndex, doc.ID, doc)
}
