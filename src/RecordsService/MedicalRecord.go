package main

type MedicalRecord struct {
	ID string `json:"Id"`
	Title string `json:"Title"`
	Link string `json:"Link"`
	HospitalName string `json:"HospitalName"`
	HospitalID string `json:"HospitalId"`
	HospitalLocation string `json:"HospitalLocation"`
	Department string `json:"Department"`
	DoctorName string `json:"DoctorName"`
	AdmissionDate string `json:"AdmissionDate"`
	DischargeDate string `json:"DischargeDate"`
	Symptoms string `json:"Symptoms"`
	Summary string `json:"Summary"`
}