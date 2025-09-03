package admin

import (
	"seaals/controller"
	"seaals/models"
	"seaals/service"
)

func NewCLI(dbPath string, basePath string) (*controller.SealLI, error) {
	db, err := models.InitialiseDB(dbPath)
	if err != nil {
		return nil, err
	}
	svc := service.NewSealService(db)
	cli := controller.NewSealLI(svc, basePath)
	return cli, nil
}
