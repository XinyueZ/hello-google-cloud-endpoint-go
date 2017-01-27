package notepad

import (
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

func init() {
	api, err := endpoints.RegisterService(&Service{},
		"documents", "v2", "Documents API", true)
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
	register("SearchResultList", "documents.search", "POST", "documents", "Search with keyword.")

	endpoints.HandleHTTP()
}
