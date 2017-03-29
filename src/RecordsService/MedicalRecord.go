package main

type MedicalRecord struct {
	Id               string `json:"Id,omitempty"`
	Title            string `json:"Title,omitempty"`
	Link             string `json:"Link,omitempty"`
	HospitalName     string `json:"HospitalName,omitempty"`
	HospitalId       string `json:"HospitalId,omitempty"`
	HospitalLocation string `json:"HospitalLocation,omitempty"`
	Department       string `json:"Department,omitempty"`
	DoctorName       string `json:"DoctorName,omitempty"`
	AdmissionDate    string `json:"AdmissionDate,omitempty"`
	DischargeDate    string `json:"DischargeDate,omitempty"`
	Symptoms         string `json:"Symptoms,omitempty"`
	Summary          string `json:"Summary,omitempty"`
}
