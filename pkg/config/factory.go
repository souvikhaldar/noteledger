package config

type Config struct {
	DB_Username   string `json:"db_username"`
	DB_Password   string `json:"db_password"`
	DB_Table_Name string `json:"db_table_name"`
}
