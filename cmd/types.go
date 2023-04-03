package cmd

type globalSettingsT struct {
	Setting1 string `json:"setting1"`
	Setting2 string `json:"setting2"`
}

type faxDocumentT struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type destinationsT struct {
	To        string `json:"to"`
	ToName    string `json:"toName"`
	ToOrgName string `json:"toCompanyName"`
}

type jobSettingsT struct {
	Documents        []faxDocumentT
	DeliveryProtocol string `json:"api"`
	Destinations     []destinationsT
	Sender           string `json:"sender"`
	Subject          string `json:"subject"`
}

type loggedInUserDetailsT struct {
	UserEmailAddress string `json:"ownerEmailAddress"`
}

type faxResponseT struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
	Message string `json:"message"`
}

/* Future Use
  type contactT struct {
	Name string `json:"name"`
	OrgName string `json:"company"`
	FaxNumber string `json:"faxNumber"`
	IsShared bool `json:"isShared"`
	IsGroup bool `json:"isGroup"`
}
*/
