package schemas

type Process struct {
	Pid  int    `json:"pid"`
	Name string `json:"name"`
}
