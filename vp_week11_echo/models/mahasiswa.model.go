package models

import (
	"net/http"
	"vp_week11_echo/db"
)

type Mahasiswa struct {
	Nim      string `json::"nim"`
	Name     string `json::"name"`
	Fakultas string `json::"fakultas"`
	Prodi    string `json::"prodi"`
	Gender   string `json::"gender"`
}

func FetchAllMahasiswa() (Response, error) {
	var object Mahasiswa
	var arrObject []Mahasiswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM mahasiswa"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&object.Nim, &object.Name, &object.Prodi, &object.Fakultas, &object.Gender)

		if err != nil {
			return res, err
		}

		arrObject = append(arrObject, object)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObject

	return res, nil
}

func StoreMahasiswa(nim string, name string, gender string, fakultas string, prodi string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO mahasiswa(nim, name, prodi, fakultas, gender) VALUES (?, ?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nim, name, prodi, fakultas, gender)

	if err != nil {
		return res, err
	}

	LastInsertId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": LastInsertId,
	}

	return res, nil
}

func UpdateMahasiswa(nim string, name string, gender string, fakultas string, prodi string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE mahasiswa SET name=?,prodi=?, fakultas=?, gender=? WHERE nim =?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, prodi, fakultas, gender, nim)

	if err != nil {
		return res, err
	}

	RowAffectedID, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected_id": RowAffectedID,
	}

	return res, nil
}

func DeleteMahasiswa(nim string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "DELETE FROM mahasiswa WHERE nim =?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nim)

	if err != nil {
		return res, err
	}

	RowAffectedID, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected_id": RowAffectedID,
	}

	return res, nil
}
