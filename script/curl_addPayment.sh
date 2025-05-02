curl --location --request POST 'http://localhost:8000/payment/new' \
--header 'Content-Type: application/json' \
--data '{
    "group_id": 1,
    "payer_id": 1,
    "amount":1000,
    "participants_ids":[2,3]
}
'