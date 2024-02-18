package database

type StudentBiodata struct {
	Id int
	Name string
	Address string
	Job string
	Reason string
}

var ListOfStudents = []StudentBiodata{
	{
		Id		: 0,
		Name	: "Ardiansyah",
		Address	: "Jalan Merauke no 51",
		Job		: "Frontend developer",
		Reason	: "Love Go language",
	},
	{
		Id		: 1,
		Name	: "Anang",
		Address	: "Jalan Merauke no 52",
		Job		: "Backend developer",
		Reason	: "Menambah skill pemograman",
	},
	{
		Id		: 2,
		Name	: "Timothy",
		Address	: "Jalan Merauke no 53",
		Job		: "Fullstack developer",
		Reason	: "Mengasah skill pemograman Go",
	},
	{
		Id		: 3,
		Name	: "Catherine",
		Address	: "Jalan Manhattan no 40",
		Job		: "Accounting Analyst",
		Reason	: "Ingin switching career ke pemograman",
	},
	{
		Id		: 4,
		Name	: "Veren",
		Address	: "Jalan Moscow no 12",
		Job		: "UI Designer",
		Reason	: "Ingin switching career ke pemograman",
	},
} 