# Go chat for Socks

- Simple chat application using Go

- This application should allow several users to talk in a chatroom and also to get stock quotes
from an API using a specific command.

- The messages are just processed when the bot consume then
- There is no front-end 
- All the commands are available through CLI or API

## How to Run

### Compose
- To run the adminer, postgres, rabbitmq, api, and robot
- Run the docker-compose inside `compose`
```shell
$~ docker-compose -f compose/docker-compose.yml up --build --remove-orphans -d
```

### Dev enviroment

#### Building
- To build all the binaries, just type:
```shell
$~ make
```

- Set UP postgres, adminer and rabbitMQ
```shell
$~ make db-up
```

- Verify the `.env` variables for your RabbitMQ URL and Database settings 

- This will generate 3 binaries {'bot', 'user', 'chatapi'}

##### Bot
- To see all the `bot` binary parameters, run:
```shell
$~ ./bot run -h
```

- So, to run the `bot` binary that starts listening for 2 queues "queue-name1" and "q1", run the following command:
```shell
$~ ./bot run -q queue-name1,q1
```

- Running the bot it will consume your shell for furthers commands, type help to see them all
```shell
>>> help
```

- It's possible to print all the queues that the bot is currently listening
```shell
>>> queues
```

__Expected Result__
![image](https://user-images.githubusercontent.com/17556614/91647030-4a0c4480-ea2c-11ea-99fe-084de5d74feb.png)


- To add a new queue to the bot listen, like add queue 'q3':
```shell
>>> add-queue q3
```

__Expected Result__
![image](https://user-images.githubusercontent.com/17556614/91647499-d2411880-ea31-11ea-9a1e-9671e79129ca.png)

##### User
- To see all the `user` binary parameters, run:
```shell
$~ ./user -h
```

- The user binary has 2 commands
  1. add 
    - It will just add a new user into the database
    ```shell
    $~ ./user add -e rafilx@gmail.com -p 123456 
    ```
    - You can add without sending the password, to just type it later

  2. login
    - It will login into the user shell
    ```shell
    $~ ./user login -e rafilx@gmail.com -p 123456 
    ```

- The user shell can print the messages from a queue and publish new messages into a queue, to show all commands, type the following command:
```shell
>>> help
```

- Publish a message into a queue
```shell
>>> publish q1 my-message
```

- Publish a message into a queue asking for a stock price
```shell
>>> publish q1 /stock=aapl.us
```

- Print the messages from a queue
```shell
>>> print q1
```
__Expected Result__
![image](https://user-images.githubusercontent.com/17556614/91647694-edad2300-ea33-11ea-8daf-113df0d80399.png)


##### Chat API 
> These commands bellow uses ~~[fish](https://fishshell.com)~~ as default shell
- The API can do the User CRUD and publish/print messages
- To run the API, type the following command:
```shell
$~ ./chatapi run
```
- If you're developing something in the API, you can use the live reloading with `air`

- This will start the API that has the following routes:
###### Add User
- POST `{baseURL}`/add-user
```shell
$~ curl -X POST -H 'Content-Type: application/json' -d '{"email": "rafilxtenfen@gmail.com", "password": "defaultpwd"}' http://localhost:4444/add-user
```


###### Login
- To get the JWT token you need to do the login
```shell
$~ curl -X POST -H 'Content-Type: application/json' -d '{"email": "rafilxtenfen@gmail.com", "password": "defaultpwd"}' http://localhost:4444/login > token.id
```

###### Users
- GET by UUID `{baseURL}/api/user/:userUUID`
```shell
$~ curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) http://localhost:4444/api/user/6tXaLO1p4I8cLbGkgl8Jgy
```

- GET All `{baseURL}/api/user`
```shell
$~ curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) http://localhost:4444/api/user
```

- POST add `{baseURL}/api/user`
```shell
$~ curl -X POST -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
-d '{
  "email": "adminx201@gmail.com", 
  "password": "defaultpwd1"
}' http://localhost:4444/api/user
```

- PUT update `{baseURL}/api/user`
```shell
$~ curl -X PUT -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
-d '{
  "email": "adminx201@gmail.com"
}' http://localhost:4444/api/user/6tXaLO1p4I8cLbGkgl8Jgy
```

###### Publish
- POST message into a queue using post body
```shell
$~ curl -X POST -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
-d '{
  "message": "publishing message"
}' http://localhost:4444/api/publish/q1
```

- GET publish a message into a queue using query params
```shell
$~ curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
 http://localhost:4444/api/publish/\?queue\=q1\&msg\=/stock=aapl.us
```

###### Messages
- POST retrieve the messages from a queue 
```shell
$~ curl -X POST -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
 http://localhost:4444/api/messages/q1
```

- GET retrieve the messages from a queue using query params
```shell
$~ curl -H 'Content-Type: application/json' -H (printf 'Authorization: Bearer %s' (cat token.id | jq -r '.token')) \
 http://localhost:4444/api/messages/\?queue\=q1
```

