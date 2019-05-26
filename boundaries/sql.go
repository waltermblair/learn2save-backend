package boundaries

import (
	"database/sql"
	"github.com/waltermblair/learn2save-backend/common"
	"github.com/waltermblair/learn2save-backend/model"
)

type DBInterface interface {
	GetAccountByID(string) model.AccountResponseBody
}

type DBClient struct {
	DB 		*sql.DB
	Config  common.Config
}

func NewDBClient(config common.Config) (*DBClient, error) {
	db, err := sql.Open("mysql", config.MySqlDSN)
	return &DBClient{db, config}, err
}

func (c *DBClient) GetAccountByID(id string) (account model.AccountResponseBody, err error) {
	err = c.DB.QueryRow(c.Config.GetAccountByIDQuery, id).Scan(&account)
	return account, err
}



