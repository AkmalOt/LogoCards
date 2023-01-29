package models

//type mediaType string
//
//const (
//	PHOTO mediaType = "photo"
//	VIDEO mediaType = "video"
//)

type Config struct {
	LocalHost struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
}

type DbData struct {
	DbConnection struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	}
}
