package conductor

import (
	"net/http"

	"github.com/gorilla/mux"
)

func forumsHandler(w http.ResponseWriter, r *http.Request) {
	RefreshTemplates()
	site, err := GetSite(r)
	if err != nil {

	}
	_, AuthUser := GetSessionUser(w, r)

	f := NewForums()
	f.SiteID = site.ID
	f.Get()

	Page := struct {
		Forums []Forum
	}{
		f.Forums,
	}

	PL := Payload{
		Title:       "Forums",
		Site:        site,
		User:        AuthUser,
		Page:        Page,
		Breadcrumbs: []Breadcrumb{{Label: "Forums", Link: "/forums"}},
	}

	Templates.ExecuteTemplate(w, "forums.html", PL)
}

func forumHandler(w http.ResponseWriter, r *http.Request) {
	RefreshTemplates()
	site, err := GetSite(r)
	if err != nil {

	}
	_, AuthUser := GetSessionUser(w, r)

	vars := mux.Vars(r)

	f := NewForums()
	f.SiteID = site.ID
	f.Key = vars["forum_guid"]
	f.Get()

	t := NewTopics()
	t.Forum = f.Forums[0].ID
	t.Get()

	Page := struct {
		Forums   []Forum
		Topics   []Topic
		ForumKey string
	}{
		f.Forums,
		t.Topics,
		f.Forums[0].Key,
	}

	PL := Payload{
		Title:       f.Forums[0].Name,
		Site:        site,
		User:        AuthUser,
		Page:        Page,
		Breadcrumbs: []Breadcrumb{{Label: "Forums", Link: "/forums"}},
	}

	Templates.ExecuteTemplate(w, "forum.html", PL)
}
