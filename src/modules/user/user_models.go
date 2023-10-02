package user

type User struct {
	ID    int64  `json:"id" gorm:"primaryKey;auto_incerment:true;index"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Email string `json:"email" gorm:"type:varchar(255)"`
}
