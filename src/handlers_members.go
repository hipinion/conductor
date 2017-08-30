package conductor

import (
	"net/http"
)

func membersHandler(w http.ResponseWriter, r *http.Request) {

	RefreshTemplates()
	site, _ := GetSite(r)

	_, AuthUser := GetSessionUser(w, r)

	US := NewUsers()
	US.Get()

	Page := struct {
		Users Users
	}{
		US,
	}

	PL := Payload{
		Title:       "Members",
		Page:        Page,
		Site:        site,
		User:        AuthUser,
		Breadcrumbs: []Breadcrumb{{Label: "Forums", Link: "/forums"}, {Label: "Members"}},
	}

	Templates.ExecuteTemplate(w, "members.html", PL)
}
