package domain

import "g/go/allsports/errs"

type TeamRepositoryStub struct {
	teams []Team
}

func (p TeamRepositoryStub) FindAll() ([]Team, *errs.AppError) {
	return p.teams, nil
}

func NewTeamRepositoryStub() TeamRepositoryStub {
	teams := []Team{
		{Id: "1001", Name: "Burlington Tigers",
			Address: Address{Address1: "44 Hickery Lyne", Address2: "", City: "Burlington", State: "MA", Zipcode: "01803"},
			Coach: Coach{Id: "10", FirstName: "Ismail", LastName: "Yusuf", Gender: "Male", PhoneNumber: "(508) 234 - 2343",
				Address: Address{Address1: "44 Hickery Lyne", Address2: "", City: "Burlington", State: "MA", Zipcode: "01803"}},
			Player: Player{}},
	}
	return TeamRepositoryStub{teams}
}
