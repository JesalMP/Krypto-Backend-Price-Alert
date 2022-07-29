package dataStructs

type user struct {
	Name, Email string
	Alerts      []alert
}

type alert struct {
	AlertPrice int
}
