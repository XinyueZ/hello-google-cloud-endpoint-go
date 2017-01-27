package notepad

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

//TABLE is table-name in database that contains all documents.
const TABLE = "Document"

//Service provides endpoint.
type Service struct {
}

//Document is model of one text content.
type Document struct {
	Key     *datastore.Key `json:"id" datastore:"-"`
	Author  string         `json:"author"`
	Title   string         `json:"title" datastore:",noindex"`
	Content string         `json:"content" datastore:",noindex"`
	Date    time.Time      `json:"date"`
}

//NewDocument is model of added text content.
type NewDocument struct {
	Author  string `json:"author"`
	Title   string `json:"title" datastore:",noindex"`
	Content string `json:"content" datastore:",noindex"`
}

// DocumentListRequest is request asking for list of documents.
type DocumentListRequest struct {
	Limit int `json:"limit" endpoints:"d=10"`
}

// DocmentList is a collection of most recent documents.
type DocmentList struct {
	Items []*Document `json:"items"`
}

//SearchWithKeywordRequest gives us list of name of documents.
type SearchWithKeywordRequest struct {
	Keyword string `json:"keyword" endpoints:"required"`
	Limit   int    `json:"limit" endpoints:"d=10"`
}

//SearchResult is searched information by SearchWithKeywordRequest.
type SearchResult struct {
	Key    *datastore.Key `json:"id" datastore:"-"`
	Author string         `json:"author"`
	Title  string         `json:"title" datastore:",noindex"`
	Date   time.Time      `json:"date"`
}

// SearchResultList is a collection of searched result by SearchWithKeywordRequest.
type SearchResultList struct {
	Items []*SearchResult `json:"results"`
}

//AddDcoment adds adds a document.
func (service *Service) AddDcoment(c context.Context, doc *NewDocument) error {
	k := datastore.NewIncompleteKey(c, TABLE, nil)
	g := &Document{
		Author:  doc.Author,
		Title:   doc.Title,
		Content: doc.Content,
		Date:    time.Now(),
	}
	_, err := datastore.Put(c, k, g)
	return err
}

//ListDocument returns list of most recent documents.
func (service *Service) ListDocument(c context.Context, r *DocumentListRequest) (*DocmentList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery(TABLE).Order("-Date").Limit(r.Limit)
	documents := make([]*Document, 0, r.Limit)
	keys, err := q.GetAll(c, &documents)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		documents[i].Key = k
	}
	return &DocmentList{documents}, nil
}

//SearchResultList contains a collection of searched result by SearchWithKeywordRequest.
func (service *Service) SearchResultList(c context.Context, r *SearchWithKeywordRequest) (*SearchResultList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery(TABLE).Filter("Title =", r.Keyword).Filter("Content =", r.Keyword).Order("-Date").Limit(r.Limit)
	results := make([]*SearchResult, 0, r.Limit)
	keys, err := q.GetAll(c, &results)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		results[i].Key = k
	}
	return &SearchResultList{results}, nil
}
