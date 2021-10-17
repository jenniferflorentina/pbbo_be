package seeds

import (
	"tubespbbo/db"
	"tubespbbo/domains/user/model"
)

func (s *Seed) SeedUser() {
	var user []model.User
	var count int64
	db.Orm.Model(&user).Count(&count)

	if count > 0 {
		return
	}

	users := make([]model.User, 5)
	users[0] = model.User{
		Username: "Jojoan",
		Password: "joan123",
		Nama:     "Eirenika Joanna Grace",
		Dob:      "02-10-2000",
		Tipe:     0,
		Alamat:   "Jln. Dimana-mana hatiku senang no.1",
		Telepon:  "081122334455",
		Email:    "joanjo@email.com",
	}
	users[1] = model.User{
		Username: "JeniJenJen",
		Password: "jeni123",
		Nama:     "Jeniffer Florentina",
		Dob:      "18-06-2001",
		Tipe:     1,
		Alamat:   "Jln. Dimana-mana hatiku gembira no.1",
		Telepon:  "081166887799",
		Email:    "jeniakatina@email.com",
	}
	users[2] = model.User{
		Username: "Marimar",
		Password: "maria123",
		Nama:     "Maria Vabiolla",
		Dob:      "12-10-2001",
		Tipe:     2,
		Alamat:   "Jln. Dimana-mana hatiku bahagia no.1",
		Telepon:  "081100337788",
		Email:    "mariaAjah@email.com",
	}
	users[3] = model.User{
		Username: "Enjeljel",
		Password: "enjel123",
		Nama:     "Tridia Enjeliani",
		Dob:      "04-10-2001",
		Tipe:     3,
		Alamat:   "Jln. Dimana-mana hatiku tertawa no.1",
		Telepon:  "081122998866",
		Email:    "jeniakatina@email.com",
	}
	users[4] = model.User{
		Username: "Melpinz",
		Password: "melvin123",
		Nama:     "Melvin Sebastian",
		Dob:      "21-07-2001",
		Tipe:     4,
		Alamat:   "Jln. Dimana-mana hatiku hanya dia no.1",
		Telepon:  "0811229988",
		Email:    "melpinzs@email.com",
	}

	_ = db.Orm.Create(&users)
}
