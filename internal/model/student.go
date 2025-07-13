package model

type Student struct {
	// ID          int     `db:"id" json:"id"`
	RollNumber  int     `db:"roll_number" json:"roll_number"`
	Name        string  `db:"name" json:"name"`
	Class       string  `db:"class" json:"class"` // Assuming enum is handled as string
	Age         *int    `db:"age" json:"age,omitempty"`
	PhoneNumber *int64  `db:"phone_number" json:"phone_number,omitempty"`
	Email       *string `db:"email" json:"email,omitempty"`
	Address     *string `db:"address" json:"address,omitempty"`
}
