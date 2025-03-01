package repo

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-chat/config"
	"github.com/mahdi-cpp/api-go-chat/model"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func InitChatRepository() {
	db = config.DatabaseInit()
}

func MessageAdd(message model.Message) (int, error) {

	result := config.DB.Create(&message)

	// Check for errors
	if result.Error != nil {
		return 0, result.Error // Return 0 ID and the error
	}

	// Return the ID of the newly created record
	return message.ID, nil
}

type ChatWithUsers struct {
	Chat          model.Chat    `json:"chat"`
	Users         []model.User  `json:"users"`         // Change to a slice of users
	LatestMessage model.Message `json:"latestMessage"` // Still hold the latest message
}

func FetchChatsForUser(userId int) ([]ChatWithUsers, error) {
	var chatsWithUsers []ChatWithUsers
	var chats []model.Chat

	// Fetch the chats associated with the user
	//err := config.DB.Where("user_ids @> ?", pq.Int32Array{int32(userId)}).Limit(100).Find(&chats).Error
	//if err != nil {
	//	return nil, err
	//}

	// Query the database to find chats where AdminId matches the provided adminID
	if err := db.Where("admin_id = ?", userId).Find(&chats).Error; err != nil {
		return nil, err
	}

	// Loop through each chat and get the details of the other users
	for _, chat := range chats {
		// Slice to hold the other users
		var chatUsers []model.User

		// Fetch details of the other users, limiting to a maximum of 3
		for _, otherUserId := range chat.UserIDs {
			if otherUserId != int32(userId) { // Compare to userId of type int
				var otherUser model.User
				if err := db.First(&otherUser, otherUserId).Error; err != nil {
					fmt.Println("Error fetch details of the other users:", err)
					return nil, err
				}
				chatUsers = append(chatUsers, otherUser)

				// Limit to a maximum of 3 users
				if len(chatUsers) >= 4 {
					break
				}
			}
		}

		// Fetch the latest message for this chat
		var latestMessage model.Message
		if err := db.Where("chat_id = ?", chat.ID).Order("created_at DESC").First(&latestMessage).Error; err != nil {
			//return nil, err
			fmt.Println("Error Fetch the latest message for this chat:", err)
		}

		// Append the chat with the other user's details to the result slice
		chatsWithUsers = append(chatsWithUsers, ChatWithUsers{
			Chat:          chat,
			Users:         chatUsers,     // This will be the other users in this case
			LatestMessage: latestMessage, // Latest message if found
		})
	}

	return chatsWithUsers, nil
}

func UpdateLastPinChange(chatID int, isPin bool) error {
	// Define the current time for LastPinChange
	currentTime := time.Now()

	// Update the LastPinChange and IsPin fields for the chat with the specified ID
	if err := db.Model(&model.Chat{}).Where("id = ?", chatID).Updates(map[string]interface{}{
		"last_pin_change": currentTime,
		"is_pin":          isPin,
	}).Error; err != nil {
		return err
	}

	return nil
}
