package patientdomain

import (
	"fmt"
	"prj101/datasource/postgres"
)

type Patient struct {
	PatientName string
	CID         string
	Age         string
	Address     string
	Disease     string
	Date        string
}

const (
	queryInsertPatient = "insert into PatientDetails (PatientName,CID,Age, Address,Disease, Date ) values($1,$2,$3,$4,$5,$6);"
	queryGetPatient    = "select PatientName,CID,Age, Address,Disease, Date coursename from PatientDetails where CID =$1"
	queryUpdatePatient = "update PatientDetails set PatientName = $1, CID = $2, Age = $3 , Address = $4, Disease = $5, Date = $6  where  CID=$7 "
	queryDeletePatient = "delete from PatientDetails where CID = $1"
)

func (ptn *Patient) PatientSave() error {
	stmt, err := postgres.Client.Prepare(queryInsertPatient)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(ptn.PatientName, ptn.CID, ptn.Age, ptn.Address, ptn.Disease, ptn.Date)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

// getting patient by cid
func (ptn *Patient) PatientGet() error {
	stmt, err := postgres.Client.Prepare(queryGetPatient)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result := stmt.QueryRow(ptn.CID)
	if getErr := result.Scan(&ptn.PatientName, &ptn.CID, &ptn.Age, &ptn.Address, &ptn.Disease, &ptn.Date); getErr != nil {
		return getErr

	}
	return nil
}

// updating
func (ptn *Patient) PatientUpdate() error {
	stmt, updateErr := postgres.Client.Prepare(queryUpdatePatient)
	if updateErr != nil {
		return updateErr
	}
	defer stmt.Close()
	_, updateErr = stmt.Exec(ptn.PatientName, ptn.CID, ptn.Age, ptn.Address, ptn.Disease, ptn.Date, ptn.CID)
	fmt.Println(ptn.PatientName)
	if updateErr != nil {
		return updateErr

	}
	return nil
}

// deleting
func (ptn *Patient) PatientDelete() error {
	stmt, err := postgres.Client.Prepare(queryDeletePatient)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(ptn.CID); err != nil {
		return err
	}
	return nil

}
