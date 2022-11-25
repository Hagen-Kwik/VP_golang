package models

import (
	"net/http"
	"vp_week11_echo/db"
)

type Mahasiswa struct {
	Nim      string `json::"nim"`
	Name     string `json::"name"`
	Gender   string `json::"gender"`
	Fakultas string `json::"fakultas"`
	Prodi    string `json::"prodi"`
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
		err = rows.Scan(&object.Nim, &object.Name, &object.Gender, &object.Fakultas, &object.Prodi)

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
