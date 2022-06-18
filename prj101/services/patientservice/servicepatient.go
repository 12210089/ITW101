package patientservice

import "prj101/model/patientdomain"

func PatientCreate(ptn patientdomain.Patient) error {
	if err := ptn.PatientSave(); err != nil {
		return err
	}
	return nil
}

// getting patient by cid

func PatientRead(cid string) (*patientdomain.Patient, error) {
	result := &patientdomain.Patient{CID: cid}
	if err := result.PatientGet(); err != nil {
		return nil, err

	}
	return result, nil
}

// updating
func UpdatePatient(ptn patientdomain.Patient) (*patientdomain.Patient, error) {
	current, err := PatientRead(ptn.CID)
	if err != nil {
		return nil, err
	}
	current.PatientName = ptn.PatientName
	current.CID = ptn.CID
	current.Age = ptn.Age
	current.Address = ptn.Address
	current.Disease = ptn.Disease
	current.Date = ptn.Date
	if err := current.PatientUpdate(); err != nil {
		return nil, err

	}
	return current, nil

}

// deleting
func DeletePatient(cid string) error {
	current, err := PatientRead(cid)
	if err != nil {
		return err
	}
	return current.PatientDelete()
}
