package notepad

import (
	"log"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
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
	Title   string         `json:"title"`
	Content string         `json:"content"`
	Date    time.Time      `json:"date"`
}

//NewDocument is model of added text content.
type NewDocument struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
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
	Keyword string `json:"keyword" endpoints:"d=''"`
}

//SearchResult is searched information by SearchWithKeywordRequest.
type SearchResult struct {
	Key    *datastore.Key `json:"id"`
	Author string         `json:"author"`
	Title  string         `json:"title"`
	Date   time.Time      `json:"date"`
}

// SearchResultList is a collection of searched result by SearchWithKeywordRequest.
type SearchResultList struct {
	Results []*SearchResult `json:"results"`
}

//AddDcoment adds a document.
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

//UpdateDcoment updates a document.
func (service *Service) UpdateDcoment(c context.Context, doc *Document) error {
	k := doc.Key
	_, err := datastore.Put(c, k, doc)
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

//SearchResults contains a collection of searched result by SearchWithKeywordRequest.
func (service *Service) SearchResults(c context.Context, r *SearchWithKeywordRequest) (*SearchResultList, error) {
	if len(strings.TrimSpace(r.Keyword)) == 0 || r.Keyword == "''" {
		return nil, nil
	}
	filters := []string{"Author =", "Title =", "Content ="}
	documents := make([]*Document, 0)
	cnt := len(filters)
	for i := 0; i < cnt; i++ {
		q := datastore.NewQuery(TABLE).Filter(filters[i], r.Keyword)
		searched := make([]*Document, 0)
		keys, err := q.GetAll(c, &searched)
		if err != nil {
			return nil, err
		}
		for i, k := range keys {
			searched[i].Key = k
		}
		documents = append(documents, searched...)
	}
	results := make([]*SearchResult, 0)
	for _, d := range documents {
		result := &SearchResult{
			d.Key,
			d.Author,
			d.Title,
			d.Date}
		results = append(results, result)
	}
	return &SearchResultList{results}, nil
}

func init() {
	api, err := endpoints.RegisterService(&Service{},
		"documents", "v1", "Documents API", true)
	if err != nil {
		log.Fatalf("Register service: %v", err)
	}

	register := func(orig, name, method, path, desc string) {
		m := api.MethodByName(orig)
		if m == nil {
			log.Fatalf("Missing method %s", orig)
		}
		i := m.Info()
		i.Name, i.HTTPMethod, i.Path, i.Desc = name, method, path, desc
	}

	register("ListDocument", "documents.list", "GET", "documents/list", "List most recent documents.")
	register("AddDcoment", "documents.add", "POST", "documents/add", "Add a document.")
	register("SearchResults", "documents.search", "POST", "documents/search", "Search with keyword.")
	register("UpdateDcoment", "documents.update", "POST", "documents/update", "Update a document.")

	endpoints.HandleHTTP()
}
