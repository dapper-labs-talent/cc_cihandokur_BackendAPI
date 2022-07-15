### `for Tests`
General purpose and for adopting Behavior Driven Development
I preferred to use `ginkgo` instead of standard `go test`.

Therefore, before running the tests, it is necessary to perform the following operations.


**installation**
```
    go env -w GO111MODULE=on
    go install github.com/onsi/ginkgo/v2/ginkgo@latest
    go get github.com/onsi/gomega/...
```
**path update**
```
    <!-- go env GOPATH ->  /home/user/go -->
    export PATH=$PATH:$(go env GOPATH)/bin
```

**create ginkgo suite file**
```
    ginkgo bootstrap
```
** test ** - `under the test path like /controller`
```
    ginkgo
    - or
    go test
```
### `settings`
```
    You can find application settings in `/config/config.toml`.
    The following files are used for Docker settings.
    - /Dockerfile
    - /docker-compose.yml
    - /docker-entrypoint.sh
    - /.env - Please do not forget to update the POSTGRES_LOCAL_DATA_PATH variable in this file according to the operating system you are using.
        - example for mac / linux : /sql/data
        - example for windows : C:/Users/coder/sql/data"
```

### `for Docker`
```
    docker compose up --build
    - or
    docker compose up
```

### `for local env`
```
    - first time
        - clone the repository to your local
        - install `go` binary or building from source <https://go.dev/dl/>
        - install `postgresql` binary or building from source <https://www.postgresql.org/download/>
        -  go mod tidy
    
    - later on
    go run cmd/main.go
```

### `Urls`
    
- SignUp : <http://localhost:9090/signup> | POST
    - Payload
    ```json
        {
        "email": "cihandokur@axiomzen.co",
        "password": "p@zw0rd",
        "firstName": "Cihan",
        "lastName": "Dokur"
        }
    ```
    - Response
    ```json
        {
        "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNpaGFuZG9rdXJAZ21haWwuY29tIiwiZXhwIjoxNjU3ODgxODQ2fQ.yKkstBJsELTVNw9ohUt3SeC9lk0p6kt2L1XlANvKJAFpsvxE_dKpSaqWVy_U0FZo0wtDi3zlt5T3JTSt-8Jy7g" 
        }
    ```

- Login : <http://localhost:9090/login> | POST
    - Payload
    ```json
        {
        "email": "cihandokur@axiomzen.co",
        "password": "p@zw0rd"
        }
    ```
    - Response
    ```json
        {
        "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNpaGFuZG9rdXJAZ21haWwuY29tIiwiZXhwIjoxNjU3ODgxODQ2fQ.yKkstBJsELTVNw9ohUt3SeC9lk0p6kt2L1XlANvKJAFpsvxE_dKpSaqWVy_U0FZo0wtDi3zlt5T3JTSt-8Jy7g" 
        }
    ```    

- Get Users : <http://localhost:9090/users> | GET - `x-authentication-token key and valid token information must be included in the header.`
    - Payload
    ```
        No payload
    ```
    - Response
    ```json
        {
            "users": [
                {
                    "email": "alex@axiomzen.co",
                    "firstname": "Alex",
                    "lastname": "Zimmerman"
                },
                {
                    "email": "cihandokur@axiomzen.co",
                    "firstname": "Cihan",
                    "lastname": "Dokur"
                }
            ]
        }
    ```

- Update Current User : <http://localhost:9090/users> | PUT - `x-authentication-token key and valid token information must be included in the header.`
    - Payload
    ```json
        {
            "firstName": "Cihan",
            "lastName": "Docker"
        }
    ```
    - Response
    ```json
        - Empty if valid token information exists and transaction is successful.
        - If there is no valid token information
        {
            "Message": "invalid token.",
            "Status": "403"
        }
    ```