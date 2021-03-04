package main

import (
	"KiKo/datamodels/domain"
	"KiKo/repositorys"
	"KiKo/tools"
	"fmt"
)

// 测试罢了
func main() {
	db := repositorys.Db
	admin := domain.TbAdmin{
		Uuid:     tools.GenerateUUID(),
		UserName: "admin2",
		Pass:     "41221321",
	}

	db.Create(&admin)
	fmt.Println(admin)
}
