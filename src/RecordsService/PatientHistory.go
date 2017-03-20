package main

type PatientHistory struct 
{
	PatientId string `json:"PatientId"`
	MedicalRecords []MedicalRecord `json:"MedicalRecords"`
}