# Key-Value Storage Server

## Installation

* git clone https://github.com/SiddharthaMishra/kv-server

##1. from Docker (recomended)
* ```docker build -t kv-server``` .
* ```docker run -d -p 8000:8000 kv-server```

##2. from dep
* ```dep ensure```
* ``` go build -o kv-server ```
* ```./kv-server```

## API

The server is run at localhost:8000 and it supports the following requests

1. ```GET http://localhost:8000/{key}``` searches for a key in the database
    ```json 
        {
            "Key": "key",
            "Value": "value"
        }
    ```
2. ```POST http://localhost:8000/``` to set a key in the DB with payload as 
     ```json 
        {
            "Key": "key",
            "Value": "value"
        }
    ```

3. ``` ws://localhost:8000/watch ``` to set up a websocket to recieve all key updated as JSON messages received in the form 
    ```json 
        {
            "Key": "key",
            "Value": "value"
        }
    ```