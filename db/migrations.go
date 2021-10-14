package db

import (
	"tubespbbo/domains/user/model"
)

func migrate() {
	_ = Orm.Sync(new(model.User))
	_ = Orm.Sync(new(model.Company))
	_ = Orm.Sync(new(model.Country))
	_ = Orm.Sync(new(model.Language))
}
