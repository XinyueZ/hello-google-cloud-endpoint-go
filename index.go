package notepad

import (
	"log"
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
	Content string         `json:"content" datastore:",noindex"`
	Date    time.Time      `json:"date"`
}

//NewDocument is model of added text content.
type NewDocument struct {
	Author  string `json:"author"`
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

//AddDcoment adds adds a document.
func (service *Service) AddDcoment(c context.Context, doc *NewDocument) error {
	k := datastore.NewIncompleteKey(c, TABLE, nil)
	g := &Document{
		Author:  doc.Author,
		Content: doc.Content,
		Date:    time.Now(),
	}
	_, err := datastore.Put(c, k, g)
	return err
}

//ListDocument returns list of most recent documents.
func (service *Service) ListDocument(c context.Context, r *DocumentListRequest) (*DocmentList, error) {
	limit := r.Limit
	if limit <= 0 {
		limit = 10
	}

	q := datastore.NewQuery(TABLE).Order("-Date").Limit(limit)
	documents := make([]*Document, 0, limit)
	keys, err := q.GetAll(c, &documents)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		documents[i].Key = k
	}
	return &DocmentList{documents}, nil
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

	register("ListDocument", "documents.list", "GET", "documents", "List most recent documents.")
	register("AddDcoment", "documents.add", "POST", "documents", "Add a document.")

	endpoints.HandleHTTP()
}
