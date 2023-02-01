package shipper

type FormCountry struct {
	ID    int `form:"id"`
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

type FormProvince struct {
	CountryID int `form:"country_id" binding:"required"`
	ID        int `form:"id"`
	Limit     int `form:"limit"`
	Page      int `form:"page"`
}
