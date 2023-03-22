package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	DB *sql.DB
}

var repo Repository

func Init() {
	if db, err := sql.Open("sqlite3", "/tmp/test.db"); err == nil {
		repo = Repository{DB: db}
	} else {
		log.Fatal("Database Cannot Be Initialised")
	}
}

func Create() int {
	const sql = "CREATE TABLE IF NOT EXISTS Cells" +
		"(id TEXT PRIMARY KEY, audio TEXT)"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

func Clear() int {
	const sql = "DELETE FROM Cells"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

func Update(c Cell) int64 {
	const sql = "UPDATE Cells SET audio = ? WHERE id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(c.Audio, c.Id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func Insert(c Cell) int64 {
	const sql = "INSERT INTO Cells(id, audio) VALUES (?, ?)"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(c.Id, c.Audio); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func ReadAll() ([]Cell, int) {
	const sql = "SELECT * FROM Cells"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		Cells := make([]Cell, 0)

		rows, err := stmt.Query()

		if err != nil {
			return Cells, -1
		}

		for rows.Next() {
			temp := Cell{}
			err := rows.Scan(&temp.Id, &temp.Audio)
			if err != nil {
				return []Cell{}, 0
			}
			Cells = append(Cells, temp)
		}

		return Cells, len(Cells)

	} else {
		return []Cell{}, -1
	}
}

func Read(id string) (Cell, int64) {
	const sql = "SELECT * FROM Cells WHERE id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		var c Cell
		row := stmt.QueryRow(id)
		if err := row.Scan(&c.Id, &c.Audio); err == nil {
			return c, 1
		} else {
			return Cell{}, 0
		}
	}
	return Cell{}, -1
}

func Delete(id string) int64 {
	const sql = "DELETE FROM Cells WHERE id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}