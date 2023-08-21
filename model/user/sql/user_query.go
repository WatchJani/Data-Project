package sql

func CreateNew() string {
	return `INSERT INTO
    team.users (
        Name,
        LastName,
        Password,
        Phone,
        Mail,
        UserName,
        BornDate
    )
VALUES
    ($ 1, $ 2, $ 3, $ 4, $ 5, $ 6, $ 7)`
}

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
