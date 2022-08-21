package news_doc_service

import "context"

// should check if id is empty
// should not able to update id, createdAt, updatedAt, deletedAt
// should auto generate updatedAt
func (dao *ESDAO) Update(ctx context.Context, doc *NewsDoc) error {
	return dao.Client.Update(ctx, dao.IndexIndex, doc.ID, doc)
}
