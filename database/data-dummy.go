package database

type Account struct {
	Email    string
	Password string
	Address  string
	Job      string
	Reason   string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ListAccounts = []Account{
	{
		Email:    "gadingramjets@gmail.com",
		Password: "sukagolang1",
		Address:  "Jalan bersama 1",
		Job:      "Software Engineer",
		Reason:   "Have a lots of money",
	},
	{
		Email:    "susanarizona@gmail.com",
		Password: "sukagolang2",
		Address:  "Jalan bersama 2",
		Job:      "Fitness Coach",
		Reason:   "Have a some time to doing healthy stuff",
	},
	{
		Email:    "columbus@gmail.com",
		Password: "sukagolang3",
		Address:  "Jalan bersama 3",
		Job:      "Marine",
		Reason:   "Have an opportunity to learn programming",
	},
	{
		Email:    "johntravolta@gmail.com",
		Password: "sukagolang4",
		Address:  "Jalan bersama 4",
		Job:      "Scrum Master",
		Reason:   "Have a support system",
	},
	{
		Email:    "markhenry@gmail.com",
		Password: "sukagolang4",
		Address:  "Jalan bersama 4",
		Job:      "Business Analyst",
		Reason:   "Food, i love food so much !",
	},
}

var ListOfErrors = map[string]string{
	"ErrorEmail":    "Email tidak ditemukan",
	"ErrorPassword": "Password salah",
}
