package domain

type PlayerRepositoryStub struct {
	players []Player
}

func (p PlayerRepositoryStub) FindAll() ([]Player, error) {
	return p.players, nil
}

func NewPlayerRepositoryStub() PlayerRepositoryStub {
	teams := []Player{
		{Id: "1001", FirstName: "Kenadid", LastName: "Yusuf", DateofBirth: "11/11/2011", Gender: "Male",
			PhoneNumber: "(508) 3201 - 7925", EmailAddress: "keniY@yahoo.com", JerseNumber: "10", Team: "Burlington Tigers",
			Address: Address{Address1: "44 Hickery Lyne", Address2: "", City: "Burlington", State: "MA", Zipcode: "01803"}},

		{Id: "1002", FirstName: "Guled", LastName: "Yusuf", DateofBirth: "05/30/2009", Gender: "Male",
			PhoneNumber: "(508) 3205 - 7925", EmailAddress: "guledY@gmail.com", JerseNumber: "9", Team: "Burlington Lazy Cats",
			Address: Address{Address1: "44 Hickery Lyne", Address2: "", City: "Burlington", State: "MA", Zipcode: "01803"}},
	}
	return PlayerRepositoryStub{teams}
}
