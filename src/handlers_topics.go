package conductor

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewTopicProcessHandler(w http.ResponseWriter, r *http.Request) {
	RefreshTemplates()
	site, err := GetSite(r)
	if err != nil {

	}
	r.ParseForm()
	_, AuthUser := GetSessionUser(w, r)

	title := r.PostFormValue("title")
	text := r.PostFormValue("text")
	subtitle := r.PostFormValue("subtitle")

	vars := mux.Vars(r)

	f := NewForums()
	f.SiteID = site.ID
	f.Key = vars["forum_guid"]
	f.Get()

	t := NewTopic()
	t.Title = title
	t.Subtitle = subtitle
	t.AuthorID = AuthUser.ID
	t.SiteID = site.ID
	t.ForumID = f.Forums[0].ID
	t.Key = GenerateKey(8, title)
	t.Set()

	p := NewPost()
	p.TopicID = t.ID
	p.AuthorID = AuthUser.ID
	p.Text = text
	p.Set()

	http.Redirect(w, r, "/forums/"+f.Forums[0].Key+"/"+t.Key, http.StatusSeeOther)

}

func NewTopicHandler(w http.ResponseWriter, r *http.Request) {
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

	Page := struct {
		Forums   []Forum
		ForumKey string
	}{
		f.Forums,
		f.Forums[0].Key,
	}

	PL := Payload{
		Title:       "New Topic in " + f.Forums[0].Name,
		Site:        site,
		User:        AuthUser,
		Page:        Page,
		Breadcrumbs: []Breadcrumb{{Label: "Forums", Link: "/forums"}, {Label: f.Forums[0].Name, Link: "/forums/" + f.Forums[0].Key}},
	}

	Templates.ExecuteTemplate(w, "topic_new.html", PL)
}

func NewPostProcessHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	_, AuthUser := GetSessionUser(w, r)

	vars := mux.Vars(r)

	forum := vars["forum_key"]

	t := NewTopics()
	t.Key = vars["topic_key"]
	t.Get()

	p := NewPost()
	p.AuthorID = AuthUser.ID
	p.Text = r.PostFormValue("text")
	p.TopicID = t.Topics[0].ID
	p.Set()

	http.Redirect(w, r, "/forums/"+forum+"/"+t.Key, http.StatusSeeOther)
}

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	RefreshTemplates()
	site, err := GetSite(r)
	if err != nil {

	}
	_, AuthUser := GetSessionUser(w, r)

	vars := mux.Vars(r)

	f := NewForums()
	f.SiteID = site.ID
	f.Key = vars["forum_key"]
	f.Get()

	t := NewTopics()
	t.Key = vars["topic_key"]
	t.Get()

	p := NewPosts()
	p.TopicID = t.Topics[0].ID
	p.Get()

	Page := struct {
		Forums   []Forum
		ForumKey string
		Topic    Topic
		Posts    []Post
	}{
		f.Forums,
		f.Forums[0].Key,
		t.Topics[0],
		p.Posts,
	}

	PL := Payload{
		Title:       t.Topics[0].Title,
		MetaTitle:   t.Topics[0].Title,
		Site:        site,
		User:        AuthUser,
		Page:        Page,
		Breadcrumbs: []Breadcrumb{{Label: "Forums", Link: "/forums"}, {Label: f.Forums[0].Name, Link: "/forums/" + f.Forums[0].Key}},
	}
	Templates.ExecuteTemplate(w, "topic.html", PL)
}
