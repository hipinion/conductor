package conductor

type Sites struct {
	Sites []Site

	Domain string
}

type Site struct {
	ID              int64
	DomainID        int64
	Name            string
	PrimaryColor    string
	PrimaryAccent   string
	SecondaryColor  string
	SecondaryAccent string

	RegistrationOpen  bool
	InvitationsActive bool
}

func NewSites() Sites {
	var ss Sites
	return ss
}

func NewSite() Site {
	var s Site
	return s
}

func (ss *Sites) Get() error {

	var params []interface{}
	w := NewWheres()

	if ss.Domain != "" {
		w.Add("host=?")
		params = append(params, ss.Domain)
	}

	sql := `SELECT s.id, d.id, s.name, sop.site_primary_color, sop.site_primary_accent,
	sop.site_registration_open, sop.site_invitations_active 
	 FROM domains d LEFT JOIN sites s ON s.id=d.site_id LEFT JOIN site_options sop ON sop.site_id=s.id ` + w.Compile()

	rows, err := DB.Query(sql, params...)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var s Site
		var ro int
		var ia int
		rows.Scan(&s.ID, &s.DomainID, &s.Name, &s.PrimaryColor, &s.PrimaryAccent, &ro, &ia)

		s.RegistrationOpen = ro == 1
		s.InvitationsActive = ia == 1

		ss.Sites = append(ss.Sites, s)
	}

	return nil
}
