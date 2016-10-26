package utils

type TableInfo struct {
	Url           string
	PageName      string
	Title         map[string]interface{}
	Body          map[string]interface{}
	TitleLen      int
	TitleIndexLen int
	BodyLen       int64
}
