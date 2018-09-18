package Models

type Cfg struct {
	DataBaseType string
	Host         string
	Username     string
	Password     string
	Port         string `default:"8080"`
	Table        string
	Param        string
}
