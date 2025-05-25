package repo

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/lib/pq"
	"github.com/mahdi-cpp/api-go-chat/config"
	"github.com/mahdi-cpp/api-go-chat/model"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

var startTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC) // Start date

func CreateFakes() {
	FakeUsers()
	FakeGroupChats()
	FakeMessages()
}

func FakeUsers() {

	// Generate and save fake users
	for i := 0; i < 90; i++ {
		user := model.User{
			Username:    Usernames[i], // Randomly select a message
			PhoneNumber: PhoneNumbers[i],
			Email:       Emails[RandomInt(1, 10)],
			FirstName:   FirstNames[RandomInt(1, 10)],
			LastName:    LastNames[RandomInt(1, 10)],
			Bio:         faker.Paragraph(),
			AvatarURL:   "chat_" + strconv.Itoa(i),
			IsOnline:    true,
			LastSeen:    time.Now().Add(-time.Duration(RandomInt(1, 100)) * time.Hour), // Random last seen time
			CreatedAt:   randomTime(startTime, time.Now()),                             // Set the current time
		}

		// Save the user to the database
		if err := config.DB.Create(&user).Error; err != nil {
			log.Printf("Failed to create user: %v", err)
		} else {
			fmt.Printf("Created user: %+v\n", user)
		}
	}
}

func FakePrivateChats() {

	for i := 3; i < 80; i++ {

		chat := model.Chat{
			Type:           "private",
			Title:          GroupChatNames[i], // Simple title with a random word
			Description:    faker.Paragraph(),
			UnreadMessages: RandomInt(0, 30), // Random unread messages count
			AdminId:        1,
			UserIDs:        pq.Int32Array{int32(i)},
			CreatedAt:      randomTime(startTime, time.Now()), // Set the current time
		}

		// Save the chat to the database
		if err := config.DB.Create(&chat).Error; err != nil {
			log.Printf("Failed to create chat: %v", err)
		} else {
			fmt.Printf("Created chat: %+v\n", chat)
		}
	}
}

func FakeGroupChats() {

	for i := 0; i < 80; i++ {

		min := 2
		max := 80
		count := rand.Intn(5) + 1
		var randomNumbers pq.Int32Array
		if i == 1 || i == 2 || i == 7 || i == 9 || i == 13 || i == 15 || i == 88 || i == 89 || i == 33 || i == 38 || i == 39 || i == 53 || i == 70 || i == 71 || i == 75 || i == 83 || i == 55 || i == 86 || i == 81 || i == 40 {
			randomNumbers = generateRandomInt32Array(min, max, count)
		} else {
			randomNumbers = generateRandomInt32Array(min, max, 1)
		}

		var chatType string
		if len(randomNumbers) > 1 {
			chatType = "group"
		} else {
			chatType = "private"
		}

		chat := model.Chat{
			Type:           chatType,
			Title:          GroupChatNames[i], // Simple title with a random word
			Description:    faker.Paragraph(),
			UnreadMessages: RandomInt(0, 30), // Random unread messages count
			AdminId:        1,
			UserIDs:        randomNumbers,
			CreatedAt:      randomTime(startTime, time.Now()), // Set the current time
		}

		// Save the chat to the database
		if err := config.DB.Create(&chat).Error; err != nil {
			log.Printf("Failed to create chat: %v", err)
		} else {
			fmt.Printf("Created chat: %+v\n", chat)
		}
	}
}

func FakeMessages() {

	for i := 0; i < 100; i++ {
		message := model.Message{
			ChatID:    rand.Intn(80) + 1,                                  // Random ChatID between 1 and 10
			UserID:    1,                                                  // Fixed UserID
			Content:   LocationMessages[rand.Intn(len(LocationMessages))], // Randomly select a message
			CreatedAt: randomTime(startTime, time.Now()),                  // Set the current time
			UpdatedAt: time.Now(),                                         // Set the current time
		}

		// Save the message to the database
		if err := config.DB.Create(&message).Error; err != nil {
			log.Printf("Failed to create message: %v", err)
		} else {
			fmt.Printf("Created message: %+v\n", message)
		}
	}
}

// Function to generate a random time within a specific range
func randomTime(start, end time.Time) time.Time {
	// Generate a random Unix timestamp between start and end
	randomUnix := rand.Int63n(end.Unix()-start.Unix()) + start.Unix()
	return time.Unix(randomUnix, 0)
}

// Function to generate a random pq.Int32Array
func generateRandomInt32Array(min, max, count int) pq.Int32Array {
	// Create a new random source
	rand.Seed(time.Now().UnixNano())
	uniqueNumbers := make(map[int32]struct{}) // To ensure uniqueness
	var numbers pq.Int32Array

	// Generate random numbers until we have the desired count
	for len(uniqueNumbers) < count {
		num := rand.Int31n(int32(max-min+1)) + int32(min) // Generate number in range [min, max]
		if _, exists := uniqueNumbers[num]; !exists {
			uniqueNumbers[num] = struct{}{}
			numbers = append(numbers, num) // Append to pq.Int32Array
		}
	}

	return numbers
}
