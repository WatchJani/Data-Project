package sql

func GetAll() string {
	return `SELECT * FROM team.users`
}

func GetUser() string {
	return `SELECT
    *
FROM
    team.users
WHERE	
    id = $1;`
}

func NewUser() string {
	return `INSERT INTO team.users (name, lastname, password, phone, mail, username, borndate)
    VALUES ($1, $2, $3, $4, $5, $6, $7)`
}
