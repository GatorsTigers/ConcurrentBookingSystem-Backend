package models

type Seat struct {
	SeatId        uint32 `json:"seatId" gorm:"primaryKey;"`
	SeatName      string `json:"seatName"`
	ScreenReferId uint32 `json:"-" gorm:"not null"`
	Screen        Screen `json:"-"  gorm:"ForeignKey:ScreenReferId;References:ScreenId"`
}
