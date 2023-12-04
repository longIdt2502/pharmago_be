package entities

type AccountTypeEntity struct {
	Id    int64  `json:"id" binding:"required"`
	Code  string `json:"code" binding:"required"`
	Title string `json:"title"`
}
