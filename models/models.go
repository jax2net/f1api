package models

type User struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Circuit struct {
	ID       int64  `json:"id"`
	Ref      string `json:"ref"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	Url      string `json:"url"`
}

type Driver struct {
	ID          int64  `json:"id"`
	Ref         string `json:"ref"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Dob         string `json:"dob"`
	Nationality string `json:"nationality"`
	Url         string `json:"url"`
}

type Constructor struct {
	ID          int64  `json:"id"`
	Ref         string `json:"ref"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Url         string `json:"url"`
}
