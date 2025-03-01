package model

import (
	"github.com/lib/pq"
	"time"
)

type Device struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"type:varchar(50);unique" json:"name"`
	AdminId int    `json:"AdminId"`
	Users   []User `gorm:"many2many:user_devices;" json:"users"` // Many-to-many relationship
}

type User struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string    `gorm:"type:varchar(50);unique" json:"username"`
	PhoneNumber string    `gorm:"type:varchar(20);unique;not null" json:"phoneNumber"`
	Email       string    `gorm:"type:varchar(100)" json:"email"`
	FirstName   string    `gorm:"type:varchar(50)" json:"firstName"`
	LastName    string    `gorm:"type:varchar(50)" json:"lastName"`
	Bio         string    `gorm:"type:text" json:"bio"`
	AvatarURL   string    `gorm:"type:varchar(255)" json:"avatarUrl"`
	IsOnline    bool      `gorm:"default:false" json:"isOnline"`
	LastSeen    time.Time `json:"lastSeen"`
	CreatedAt   time.Time `gorm:"default:now()" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"default:now()" json:"updatedAt"`
}

type Chat struct {
	ID             int           `gorm:"primaryKey;autoIncrement" json:"id"`
	AdminId        int           `gorm:"references:users(id);onDelete:SET NULL" json:"adminId"`
	UserIDs        pq.Int32Array `gorm:"type:integer[]" json:"userIDs"`
	Type           string        `gorm:"type:varchar(20);not null;check:type in ('private', 'group', 'channel')" json:"type"`
	Title          string        `gorm:"type:varchar(100)" json:"title"` // For group/channel
	Description    string        `gorm:"type:text" json:"description"`   // For group/channel
	UnreadMessages int           `gorm:"default:0" json:"unreadMessages"`
	IsPin          bool          `gorm:"default:false" json:"isPin"`
	LastPinChange  time.Time     `gorm:"default:now()" json:"lastPinChange"`
	CreatedAt      time.Time     `gorm:"default:now()" json:"createdAt"`
	UpdatedAt      time.Time     `gorm:"default:now()" json:"updatedAt"`
}

type Message struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ChatID    int       `json:"chatId"`  // Identifier of the chat where the message belongs
	UserID    int       `json:"userId"`  // Identifier of the user who sent the message
	Content   string    `json:"content"` // The actual message content
	Type      string    `gorm:"default:text" json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CommandChatPin struct {
	ChatID int  `json:"chatID"`
	IsPin  bool `json:"isPin"`
}
