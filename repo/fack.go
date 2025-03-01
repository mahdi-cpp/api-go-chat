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
	FakePrivateChats()
	FakeMessages()
}

func FakeUsers() {

	// Generate and save fake users
	for i := 0; i < 90; i++ {
		user := model.User{
			Username:    usernames[i], // Randomly select a message
			PhoneNumber: phoneNumbers[i],
			Email:       emails[RandomInt(1, 10)],
			FirstName:   firstNames[RandomInt(1, 10)],
			LastName:    lastNames[RandomInt(1, 10)],
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
			Title:          groupChatNames[i], // Simple title with a random word
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
			Title:          groupChatNames[i], // Simple title with a random word
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

func FakeMessages() {

	for i := 0; i < 190; i++ {
		message := model.Message{
			ChatID:    rand.Intn(80) + 1,                                  // Random ChatID between 1 and 10
			UserID:    1,                                                  // Fixed UserID
			Content:   locationMessages[rand.Intn(len(locationMessages))], // Randomly select a message
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

func GenerateRandomArray() []int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random length between 1 and 10
	length := rand.Intn(10) + 1 // rand.Intn(10) gives a number from 0 to 9, so we add 1

	// Create a slice of uint32 with the random length
	randomArray := make([]int, 0, length) // Use a zero-length slice with capacity

	// Fill the array with random uint32 integers, ignoring 1
	for len(randomArray) < length {
		num := rand.Intn(80) // Generate a random uint32 integer from 0 to 79
		if num != 1 {        // Check if the number is not 1
			randomArray = append(randomArray, num) // Only add if it's not 1
		}
	}

	return randomArray
}

var firstNames = []string{
	"Ali",
	"Fatemeh",
	"Mohammad",
	"Sajad",
	"Saeid",
	"Niloofar",
	"Omid",
	"Shirin",
	"Kaveh",
	"Yasaman",
	"Peyman",
	"Nasrin",
	"Sami",
	"Parisa",
	"Reza",
	"Elham",
	"Farhad",
	"Mahsa",
	"Navid",
	"Simin",
}
var lastNames = []string{
	"Rezaei",
	"Mohammadi",
	"Ahmadi",
	"Hosseini",
	"Javid",
	"Sadeghi",
	"Amiri",
	"Faraji",
	"Malek",
	"Karimi",
	"Safavi",
	"Ghaffari",
	"Nouri",
	"Azizi",
	"Sharifi",
	"Bahrami",
	"Zand",
	"Khanlari",
	"Nourian",
	"Yazdani",
}

var usernames = []string{
	"Ali Reza",
	"Fatemeh Mohammadi",
	"Reza Ghasemi",
	"Sajad Esmaili",
	"Mohammad Jafari",
	"Parsa Nasiri",
	"Ahmad Hashemi",
	"Niloofar Sharifi",
	"Masoud Mohammadi",
	"Sara Ebrahimi",
	"Amir Hossein",
	"Parisa Khosravi",
	"Behzad Salehi",
	"Yasmin Azimi",
	"Farhad Zare",
	"Leila Sadeghi",
	"Kamran Vaziri",
	"Shirin Moghaddam",
	"Navid Rahimi",
	"Shirin Fadaei",
	"Arash Noor",
	"Zeynab Safavi",
	"Babak Fathi",
	"Mahsa Gholami",
	"Milad Amini",
	"Fariba Kiani",
	"Vahid Farahani",
	"Elham Gholizadeh",
	"Javad Saidi",
	"Sima Marzban",
	"Hossein Shafiei",
	"Shadi Mortezaei",
	"Navid Kamali",
	"Mahdi Abdolmaleki",
	"Amir Ali",
	"Samaneh Nazari",
	"Armin Jahanbakhsh",
	"Fatemeh Nouri",
	"Behzad Karami",
	"Zahra Nikpour",
	"Ali Akbar",
	"Parvin Rahmani",
	"Kamran Fadaei",
	"Leila Hosseini",
	"Reza Poursaeedi",
	"Maryam Yaghoubi",
	"Milad Rahimi",
	"Shirin Moshiri",
	"Arash Khosrowshahi",
	"Fatemeh Mohsenifar",
	"Farzad Rahmani",
	"Sahar Mohammadi",
	"Hossein Khoshnood",
	"Zeynab Shafiei",
	"Javad Khosravi",
	"Niloufar Fadavi",
	"Vahid Mohsenifar",
	"Elham Mohammadi",
	"Amir Mohsen",
	"Sara Gholipour",
	"Samira Aghaei",
	"Armin Moradi",
	"Farah Pourali",
	"Babak Bahrami",
	"Shadi Salimi",
	"Ali Tabrizi",
	"Zahra Nematollahi",
	"Ahmad Azizi",
	"Yasaman Ranjbar",
	"Navid Safarzadeh",
	"Fatemeh Nasr",
	"Masoud Khoshbakht",
	"Mahsa Khodabakhsh",
	"Farhad Behzad",
	"Sima Rahimi",
	"Reza Amini",
	"Leila Aghaei",
	"Kamran Javad",
	"Shirin Zand",
	"Vahid Ebrahimi",
	"Milad Zeynali",
	"Samaneh Heidari",
	"Niloofar Mohammadi",
	"Amir Reza",
	"Parisa Rahmani",
	"Yasmin Javid",
	"Ahmad Zare",
	"Shadi Gholizadeh",
	"Farzad Amiri",
	"Elham Azizi",
	"Ali Shafiei",
	"Fatemeh Rahimian",
	"Javad Sadeghi",
	"Zahra Ebrahimi",
	"Navid Moradi",
	"Behzad Gholami",
	"Samira Mohammadi",
	"Mahsa Zandi",
	"Armin Golestani",
	"Zeynab Khosravi",
	"Farhad Rahimi",
	"Shirin Tabrizi",
	"Milad Mohammadi",
	"Vahid Aghaei",
	"Ali Karami",
	"Samaneh Fadaei",
	"Niloofar Ghasemi",
	"Ahmad Rahimi",
}

var phoneNumbers = []string{
	"+989123456789",
	"+989234567890",
	"+989345678901",
	"+989456789012",
	"+989567890123",
	"+989678901234",
	"+989789012345",
	"+989890123456",
	"+989901234567",
	"+989012345678",
	"+989123456780",
	"+989234567801",
	"+989345678912",
	"+989456789023",
	"+989567890134",
	"+989678901245",
	"+989789012356",
	"+989890123467",
	"+989901234578",
	"+989012345689",
	"+989123456791",
	"+989234567802",
	"+989345678913",
	"+989456789024",
	"+989567890135",
	"+989678901246",
	"+989789012357",
	"+989890123468",
	"+989901234579",
	"+989012345690",
	"+989123456792",
	"+989234567803",
	"+989345678914",
	"+989456789025",
	"+989567890136",
	"+989678901247",
	"+989789012358",
	"+989890123469",
	"+989901234580",
	"+989012345691",
	"+989123456793",
	"+989234567804",
	"+989345678915",
	"+989456789026",
	"+989567890137",
	"+989678901248",
	"+989789012359",
	"+989890123470",
	"+989901234581",
	"+989012345692",
	"+989123456794",
	"+989234567805",
	"+989345678916",
	"+989456789027",
	"+989567890138",
	"+989678901249",
	"+989789012360",
	"+989890123471",
	"+989901234582",
	"+989012345693",
	"+989123456795",
	"+989234567806",
	"+989345678917",
	"+989456789028",
	"+989567890139",
	"+989678901250",
	"+989789012361",
	"+989890123472",
	"+989901234583",
	"+989012345694",
	"+989123456796",
	"+989234567807",
	"+989345678918",
	"+989456789029",
	"+989567890140",
	"+989678901251",
	"+989789012362",
	"+989890123473",
	"+989901234584",
	"+989012345695",
	"+989123456797",
	"+989234567808",
	"+989345678919",
	"+989456789030",
	"+989567890141",
	"+989678901252",
	"+989789012363",
	"+989890123474",
	"+989901234585",
	"+989012345696",
	"+989123456798",
	"+989234567809",
	"+989345678920",
	"+989456789031",
}

var emails = []string{
	"ali.reza@example.com",
	"fatemeh.sardar@example.com",
	"mohammad.nasr@example.com",
	"zahra.pour@example.com",
	"saeid.khaki@example.com",
	"niloofar.javid@example.com",
	"omid.amiri@example.com",
	"shirin.faraji@example.com",
	"kaveh.malek@example.com",
	"yasaman.hosseini@example.com",
	"peyman.esfandiari@example.com",
	"nasrin.saberi@example.com",
	"sami.azizi@example.com",
	"parisa.bahrami@example.com",
	"reza.mohajer@example.com",
	"elham.amiri@example.com",
	"farhad.bahrami@example.com",
	"mahsa.sharifi@example.com",
	"navid.ali@example.com",
	"simin.nasiri@example.com",
}

var groupChatNames = []string{
	"Persian Pals",
	"Tehran Talkers",
	"Farsi Friends",
	"Iranian Insights",
	"Culture Connect",
	"Persian Poetry Circle",
	"The Persian Hub",
	"Shiraz Shenanigans",
	"Caspian Conversations",
	"Tehran Tribes",
	"Persian Pride",
	"The Farsi Forum",
	"Persian Heritage Hangout",
	"Persian Cuisine Lovers",
	"Persian Literature Lovers",
	"The Iranian Network",
	"Persian Art Appreciation",
	"Persian Diaspora Dialogues",
	"The Farsi Fellowship",
	"Persian Music Mates",
	"Persian History Buffs",
	"Persian Traditions",
	"The Iranian Community",
	"Farsi Family",
	"Persian Dreams",
	"Persian Network",
	"Cultural Caravan",
	"Persian Language League",
	"Persian Cinema Circle",
	"Persian Festivals Group",
	"Iran's Hidden Gems",
	"Persian Style Squad",
	"Persian Tea Time",
	"Persian Heritage Hunters",
	"Persian Artisans",
	"Persian Fashion Forum",
	"Persian Friendship Circle",
	"The Iranian Explorers",
	"Persian Garden Gatherings",
	"Persian Wisdom Exchange",
	"The Farsi Fellowship",
	"Persian Heritage Seekers",
	"Persian Traditions and Tales",
	"Persian Music & Dance",
	"Iranian Innovations",
	"Persian Nature Lovers",
	"Persian Culinary Club",
	"The Farsi Family Tree",
	"Persian Nostalgia",
	"Iran's Cultural Chronicles",
	"Persian Book Club",
	"Persian Travel Group",
	"Persian Arts & Crafts",
	"Persian Sports Enthusiasts",
	"Persian Language Learners",
	"Persian Family Gatherings",
	"Persian History Buffs",
	"Persian Youth Voices",
	"Persian Cultural Exchange",
	"Persian Tech Talk",
	"Persian Gamers Unite",
	"Persian Movie Buffs",
	"Persian Artists Collective",
	"Persian Entrepreneurs",
	"Persian Cooking Club",
	"Persian Travel Enthusiasts",
	"Persian History Hunters",
	"Persian Language Exchange",
	"Persian Nature Explorers",
	"Persian Gardening Group",
	"Persian Fitness Friends",
	"Persian Meditation Circle",
	"Persian Wellness Warriors",
	"Persian Book Lovers",
	"Persian Fashionistas",
	"Persian Family Ties",
	"Persian Tea Lovers",
	"Persian Cultural Heritage",
	"Persian Spiritual Seekers",
	"Persian Trivia Night",
	"Persian Technology Trends",
	"Persian Art Exhibits",
	"Persian Heritage Society",
	"Persian Music Lovers",
	"Persian Foodies",
	"Persian Nature Photographers",
	"Persian DIY Enthusiasts",
	"Persian Cultural Explorers",
	"Persian Life Stories",
	"Persian Digital Nomads",
	"Persian Startups",
	"Persian Philanthropy",
	"Persian Career Connections",
	"Persian Networking Group",
	"Persian Historical Society",
	"Persian Community Builders",
	"Persian Science Fiction Fans",
	"Persian Adventure Seekers",
	"Persian Language Lovers",
	"Persian Hobbies Hub",
	"Persian Creative Minds",
	"Persian Fitness Challenge",
	"Persian Culinary Adventures",
	"Persian Family Tree",
	"Persian Sports Fans",
	"Persian Technology Innovators",
	"Persian Art and Culture",
	"Persian Global Connections",
	"Persian Cooking Adventures",
	"Persian Wellness Community",
	"Persian Social Club",
	"Persian Podcast Circle",
	"Persian Adventure Club",
	"Persian History Buffs",
	"Persian Cultural Appreciation",
	"Persian Virtual Meetups",
	"Persian Friends Forever",
	"Persian Dreamers",
	"Persian Networking Ninjas",
	"Persian Creative Collaborators",
	"Persian Artistic Souls",
	"Persian Virtual Explorers",
}

// Predefined location-based messages
var locationMessages = []string{
	"یک مقدار سرعتت رو بیشتر کن! 🚙💨",
	"سلام، محمد! به نظر میاد که مسیر خیلی شیب داره. 😅",
	"توجه! همه آماده باشند، شروع می‌کنیم! 🏞️",
	"بیشتر به جلو بروید! 🌄 ما به قله نزدیک می‌شویم.",
	"آیا وضعیت جاده خوبه؟ من احساس می‌کنم کمی لیزه. 🛤️",
	"شما کجا هستید؟ ما منتظرتان هستیم! 📍",
	"حتماً سرعت رو کنترل کن، جاده خطرناک است! ⚠️",
	"بهتره که کمی توقف کنیم، نیاز به استراحت داریم. ⏳",
	"آیا به سمت قله می‌رویم؟ من از سرعت لذت می‌برم! ⛰️",
	"بچسبید! مسیر خیلی چالش برانگیزه! 🏔️",
	"یک مقدار شتاب بزن، خیلی دیر شد! 🚗💨",
	"چقدر دما پایین اومده! بهتره لباس گرم‌تری بپوشیم. 🥶",
	"به سمت راست بپیچ، جاده رو گم کردیم! 🔄",
	"یادآوری: حتماً تجهیزات ایمنی رو فراموش نکنید! 🛡️",
	"به نظر می‌رسه که بارون میاد! چترها رو آماده کنید! ☔",
	"حواست باشه، سنگ‌ها توی جاده هستن! ⚠️",
	"آیا کسی می‌دونه که کجا هستیم؟ GPS رو چک کن! 🗺️",
	"سرعت رو کم کن، ماشین جلویی داره توقف می‌کنه! 🚦",
	"خیلی خوب می‌ریم! همه با هم به قله می‌رسیم! 🌟",
	"باید کمی استراحت کنیم، بنزین داریم؟ ⛽",
	"سلام! بچه‌ها، چطورید؟ آماده‌اید برای ادامه سفر؟ 🙌",
	"حتماً از جاده‌های پر از چاله دوری کنید! 🕳️",
	"چقدر زیباست! از منظره لذت ببرید! 🌅",
	"آیا کسی می‌تونه نقشه رو بخونه؟ ما کجا هستیم؟ 📍",
	"یک مقدار در سرعتت احتیاط کن، جاده پر از گرد و غباره! 🌪️",
	"من به سمت چپ میرم، شاید مسیر بهتری پیدا کنیم! ↖️",
	"کسی می‌دونه چند کیلومتر تا قله باقی مونده؟ 🏔️",
	"باید زودتر حرکت کنیم، شب نزدیکه! 🌙",
	"آیا می‌تونیم یک عکس دسته‌جمعی بگیریم؟ 📸",
	"به نظر میاد که یکی از ماشین‌ها مشکل داره! ⚙️",
	"همه آماده‌اید؟ می‌خواهیم با سرعت بیشتری حرکت کنیم! 🚀",
	"چقدر اینجا زیباست، نمی‌خوام برم! 🌲",
	"آیا کسی از گروه عقب موند؟ باید بررسی کنیم! 👀",
	"سرعتت رو کم کن، من دارم میام! 🚙",
	"چقدر اینجا دما پایین اومده! لباس گرم بپوشید! 🧥",
	"به سمت چپ بپیچ، جاده بهتری در پیش داریم! 🛣️",
	"باید زودتر به اردوگاه برسیم! ⛺",
	"چقدر هیجان‌انگیز! ما داریم به قله نزدیک می‌شویم! 🏔️",
	"آیا می‌دانید کجا توقف کنیم؟ من گرسنه‌ام! 🍔",
	"به نظر می‌رسه که بارون داره میاد! چترها رو بیارید! ☔",
	"کسی می‌تونه مسیر رو چک کنه؟ ممکنه گم بشیم! 🗺️",
	"یک مقدار دیگه باید برونیم تا به قله برسیم! 🌌",
	"من احساس می‌کنم که جاده کمی خطرناک شده! ⚠️",
	"سفر عالیه! به همه بگید که ادامه بدیم! 🚗💨",
	"چقدر باحال بود! حتماً این سفر رو تکرار می‌کنیم! 🎉",
	"وقت استراحت و نوشیدنی! 💦",
	"شما بچه‌ها، کجا هستید؟ ما منتظرتان هستیم! 📍",
	"باید به سرعت بیشتری حرکت کنیم، شب داره میاد! 🌙",
	"آیا کسی می‌دونه چند کیلومتر تا قله باقی مونده؟ ⛰️",
	"ما داریم به قله نزدیک می‌شویم! 🌄",
	"بیشتر احتیاط کن، جاده سُر هست! ⚠️",
	"آیا کسی می‌دونه دما چقدر پایین رفته؟ 🥶",
	"به نظر میاد که بارون داره میاد! چترها رو بیارید! ☔",
	"سرعت رو کم کن، جاده پر از چاله است! 🕳️",
	"آیا می‌تونید نقشه رو چک کنید؟ ما کجا هستیم؟ 📍",
	"باید به سمت قله بریم! آماده‌اید؟ 🌄",
	"این مسیر خیلی زیباست! از مناظر لذت ببرید! 🌲",
	"چقدر باحال بود! بیا عکس بگیریم! 📸",
	"آیا کسی سوخت داره؟ باید زودتر پر کنیم! ⛽",
	"به همه بگید که حرکت کنیم! وقت کم داریم! ⏳",
	"سرعتت رو بیشتر کن، ما داریم عقب می‌افتیم! 🚗💨",
	"توجه! ماشین جلویی داره توقف می‌کنه! 🚦",
	"چقدر قشنگه! حتماً از اینجا عکس بگیریم! 🌅",
	"امیدوارم به زودی به محل اردوگاه برسیم! ⛺",
	"آیا کسی می‌دونه چقدر تا قله باقی مونده؟ ⛰️",
	"باید یه استراحت کوتاه بکنیم، گرسنه‌ام! 🍔",
	"آیا همه ماشین‌ها سالم هستن؟ یک چک کنید! 🔧",
	"سرعت بزنید، ما داریم به شب نزدیک می‌شویم! 🌙",
	"چقدر اینجا دما پایین اومده! لباس گرم بپوشید! 🧥",
	"به نظر میاد که یکی از ماشین‌ها مشکل داره! ⚙️",
	"باید به سمت چپ بپیچیم! جاده بسته است! 🔄",
	"امیدوارم بارون نیاد، سفر خیلی خوبه! 🌞",
	"آیا کسی می‌دونه کجا باید توقف کنیم؟ 🛑",
	"بیا بچه‌ها! باید به قله برسیم! 🏔️",
	"چقدر زیباست! از طبیعت لذت ببرید! 🌼",
	"آیا کسی داره به موسیقی گوش می‌ده؟ 🎶",
	"بیشتر سرعتت رو افزایش بده! ما داریم دیر می‌شیم! 🚀",
	"چقدر هیجان‌انگیز! به قله نزدیک می‌شویم! 🌟",
	"آیا کسی می‌دونه چند کیلومتر تا اردوگاه باقی مونده؟ 🏕️",
	"به یاد داشته باشید، همه باید احتیاط کنند! ⚠️",
	"این مسیر خیلی طولانیه! آیا کسی می‌خواد استراحت کنه؟ ⏳",
	"حواست باشه، سنگ‌ها توی جاده هستن! 🪨",
	"چقدر خوبه که همه با هم هستیم! ❤️",
	"باید به سمت راست بپیچیم، جاده داره تغییر می‌کنه! ↖️",
	"امیدوارم بارون نیاد، سفر خیلی خوبه! 🌤️",
	"آیا می‌تونید جلوتر برید؟ من دارم عقب می‌افتم! 🚙",
	"به یاد داشته باشید، هیچ‌کس رو تنها نذارید! 🤝",
	"یک مقدار سرعتت رو بیشتر کن! ⏩",
	"آیا کسی می‌تونه بنزین بگیره؟ ⛽",
	"چقدر لذت‌بخش بود! بیایید دوباره بیاییم! 🎉",
	"باید به زودی به اردوگاه برسیم! ⛺",
	"هر کسی که عقب موند، باید سریعتر بیاد! 🏃‍♂️",
	"چقدر اینجا زیباست! از مناظر لذت ببرید! 🌲",
	"سرعتت رو کم کن، راه خطرناکه! ⚠️",
	"این مسیر خیلی عالیه! ادامه بده! 🚗💨",
	"آیا کسی داره عکس می‌گیره؟ 📸",
	"باید سریعتر حرکت کنیم، شب نزدیکه! 🌙",
	"چقدر هیجان‌انگیز! ما داریم به قله نزدیک می‌شویم! 🏔️",
	"آیا کسی می‌دونه کجا باید توقف کنیم؟ 🛑",
	"بیا به جلو حرکت کنیم! ما به قله نزدیک می‌شویم! 🌄",
	"به نظر میاد که بارون داره میاد! چترها رو بیارید! ☔",
	"سرعت رو کم کن، جاده پر از چاله است! 🕳️",
	"آیا کسی می‌تونه نقشه رو بخونه؟ ما کجا هستیم؟ 📍",
	"باید به سمت قله بریم! آماده‌اید؟ 🌅",
	"چقدر باحال بود! بیا عکس بگیریم! 📸",
	"آیا کسی سوخت داره؟ باید زودتر پر کنیم! ⛽",
	"به همه بگید که حرکت کنیم! وقت کم داریم! ⏳",
	"سرعتت رو بیشتر کن، ما داریم عقب می‌افتیم! 🚗💨",
	"توجه! ماشین جلویی داره توقف می‌کنه! 🚦",
	"چقدر قشنگه! حتماً از اینجا عکس بگیریم! 🌅",
	"امیدوارم به زودی به محل اردوگاه برسیم! ⛺",
	"آیا کسی می‌دونه چقدر تا قله باقی مونده؟ ⛰️",
	"باید یه استراحت کوتاه بکنیم، گرسنه‌ام! 🍔",
	"آیا همه ماشین‌ها سالم هستن؟ یک چک کنید! 🔧",
	"سرعت بزنید، ما داریم به شب نزدیک می‌شویم! 🌙",
	"چقدر اینجا دما پایین اومده! لباس گرم بپوشید! 🧥",
	"به نظر میاد که یکی از ماشین‌ها مشکل داره! ⚙️",
	"باید به سمت چپ بپیچیم! جاده بسته است! 🔄",
	"امیدوارم بارون نیاد، سفر خیلی خوبه! 🌞",
	"آیا کسی می‌دونه کجا باید توقف کنیم؟ 🛑",
	"بیا بچه‌ها! باید به قله برسیم! 🏔️",
	"چقدر زیباست! از طبیعت لذت ببرید! 🌼",
	"آیا کسی داره به موسیقی گوش می‌ده؟ 🎶",
	"بیشتر سرعتت رو افزایش بده! ما داریم دیر می‌شیم! 🚀",
	"چقدر هیجان‌انگیز! به قله نزدیک می‌شویم! 🌟",
	"آیا کسی می‌دونه چند کیلومتر تا اردوگاه باقی مونده؟ 🏕️",
	"به یاد داشته باشید، همه باید احتیاط کنند! ⚠️",
	"این مسیر خیلی طولانیه! آیا کسی می‌خواد استراحت کنه؟ ⏳",
}
