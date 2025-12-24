# api-golang-codebase
TLab Technical Test

## Sebelum running tambahan env
`DATABASE_URL="host=localhost user=user password=password dbname=dbname port=5432 sslmode=disable TimeZone=Asia/Jakarta"`

## Cara running
`go run cmd/app/main.go`

## CURL Testing
```curl --location 'localhost:8080/event' \
--header 'Content-Type: application/json' \
--data '{
    "Description": "Acara Satu",
    "Quota": 100,
    "StartDate": "2025-12-24T03:07:17.297Z",
    "EndDate": "2025-12-25T03:07:17.297Z"
}'```
