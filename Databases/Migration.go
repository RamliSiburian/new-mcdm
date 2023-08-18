package databases

import (
	"fmt"
	models "mcdm/Models"
	mysql "mcdm/Pkg/Mysql"
)

func Migration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Kriteria{}, &models.Alternatif{}, &models.PerbandinganAhp{}, &models.PerbandinganMopa{})

	if err != nil {
		panic("migration failed")
	}

	fmt.Println("Migration success")
}
