curl -X POST "http://localhost:8000/group/new" \
     -H "Content-Type: application/json" \
     -d '{
            "group_name": "a",
            "users": ["a", "b", "c"]
        }'
