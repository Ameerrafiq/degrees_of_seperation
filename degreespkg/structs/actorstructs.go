package structs

type ActorMoviesList struct {
	Url    string       `json:"url,omitempty"`
	Type   string       `json:"type,omitempty"`
	Name   string       `json:"name,omitempty"`
	Movies []MoviesList `json:"movies,omitempty"`
}

type MoviesList struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
	Role string `json:"role,omitempty"`
}
