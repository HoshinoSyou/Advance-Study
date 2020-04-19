package dao_sql

type Table struct {
	Sno     int      `json:"sno"`
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	Name          string `json:"name"`
	Id            string `json:"id"`
	Type          string `json:"type"`
	Class         string `json:"class"`
	Required      string `json:"required"`
	Teacher       string `json:"teacher"`
	TimeAndPlaces []TimeAndPlace
}

type TimeAndPlace struct {
	Day     string `json:"day"`
	Section string `json:"section"`
	Week    string `json:"week"`
	Place   string `json:"place"`
}

func AddInSql(t Table)  {
	db = Init()
	db.Create(&t)
}
