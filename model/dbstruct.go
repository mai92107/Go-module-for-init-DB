package model

type Machine struct {
    DB Dbconfig `json:"db"`
}

type Dbconfig struct {
    Status     bool   `json:"status"`
    DBUser     string `json:"db_user"`
    DBPassword string `json:"db_password"`
    DBHost     string `json:"db_host"`
    DBPort     string `json:"db_port"`
    DBName     string `json:"db_name"`
}
