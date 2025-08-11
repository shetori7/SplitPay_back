curl --location --request POST 'http://localhost:8000/payment/new' \
--header 'Content-Type: application/json' \
--data '{
    "group_uuid": "1e6821e6-44db-4539-9b0b-16d8bfa002be",
    "payer_id": 2,
    "amount":1000,
    "participants_ids":[1,3]
}
'
