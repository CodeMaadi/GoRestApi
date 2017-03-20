# GoRestApi
GoRestApi

Instructions to build:
Go to src\RecordsService
Build using command "go build"
It generates src\RecordsService\RecordsService.exe (On Windows)
Run RecordsService.exe

It will run the service on http://localhost:1234

Use the following GET/POST/DELETE calls to play with on a tool such as PostMan:
For Users:
GET: http://localhost:1234/users
GET: http://localhost:1234/users/1
POST: http://localhost:1234/users/4
DELETE: http://localhost:1234/users/4

For MedicalRecords:
GET: http://localhost:1234/MedicalRecords
GET: http://localhost:1234/MedicalRecords/1
POST: http://localhost:1234/MedicalRecords/4
DELETE: http://localhost:1234/MedicalRecords/4

For PatientHistory:
GET: http://localhost:1234/patients/1/history

Some of the reference links followed:
https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
https://thenewstack.io/make-a-restful-json-api-go/

Further work
Integrate with Echo if required https://echo.labstack.com/guide
