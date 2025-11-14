package crud

type GuitarModel struct {
    ID    uint   `gorm:"primaryKey" json:"id"`
    Brand string `json:"brand"`
    Name  string `json:"name"`
    Year  int    `json:"year"`
}
