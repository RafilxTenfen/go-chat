# Default shell is fish

# Create a User 
curl -X POST -H 'Content-Type: application/json' -d '{"email": "adminx@gmail.com", "password": "defaultpwd"}' http://localhost:4444/add-user

# Login
curl -X POST -H 'Content-Type: application/json' -d '{"email": "adminx@gmail.com", "password": "defaultpwd"}' http://localhost:4444/login > token.id

# add
curl -X POST -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
-d '{
  "email": "adminx201@gmail.com", 
  "password": "defaultpwd1"
}' http://localhost:4444/api/user

# get all
curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) http://localhost:4444/api/user

# get by uuid
curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) http://localhost:4444/api/user/6tXaLO1p4I8cLbGkgl8Jgy

# delete by uuid
curl -X DELETE -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) http://localhost:4444/api/user/4cVmaGK5dNTSoH9nTJxhtQ 

# update
curl -X PUT -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
-d '{
  "email": "adminx201@gmail.com"
}' http://localhost:4444/api/user/6tXaLO1p4I8cLbGkgl8Jgy

# CHAT ROUTES - the JWT is needes to validade the user

# PUBLISH

# publish message into any queue
curl -X POST -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
-d '{
  "message": "publishing message"
}' http://localhost:4444/api/publish/q1

# publish message into any queue using query param
curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
 http://localhost:4444/api/publish/\?queue\=q1\&msg\=anymessage

# publish message into any queue using query param
curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
 http://localhost:4444/api/publish/\?queue\=q1\&msg\=/stock=aapl.us

# publish message into any queue using query param with ""
curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
 http://localhost:4444/api/publish/\?queue\=q1\&msg\="/stock=aapl.us"

# MESSAGES

# get messages using queue name as param
curl -X POST -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
 http://localhost:4444/api/messages/q1

# get messages using queue name as query param
curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
 http://localhost:4444/api/messages/\?queue\=q1
