# techinal-test
TLab Technical Test

## Sebelum running tambahan env
`DATABASE_URL="host=localhost user=user password=password dbname=dbname port=5432 sslmode=disable TimeZone=Asia/Jakarta"`

## Cara running
`go run cmd/app/main.go`

## CURL Testing
### API Create Event
```curl --location 'localhost:8080/event' \
--header 'Content-Type: application/json' \
--data '{
    "Description": "Acara Satu",
    "Quota": 100,
    "StartDate": "2025-12-24T03:07:17.297Z",
    "EndDate": "2025-12-25T03:07:17.297Z"
}'```

### API Create Booking
```curl --location 'localhost:8080/booking' \
--header 'Content-Type: application/json' \
--data '{
    "IDEvent": 1,
    "Total": 20
}'```


## DBDiagram.io Source Code
```

Table events {
  id integer [primary key]
  description varchar
  quota varchar
  start_date timestamp
  end_date timestamp
}

Table bookings {
  id integer [primary key]
  id_event varchar
  total text 
  created_at timestamp
}

Ref events: bookings.id_event > events.id 
```

## DB Design
![Image DB Design](/doc/db.png)