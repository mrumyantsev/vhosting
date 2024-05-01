# Configuration files

The file containing the program configuration parameters to be changed is located in the ./configs directory and is called config.yml. If any property from this file is removed or its value is written incorrectly, it will be replaced with a default value on startup with a warning in the console.

## Structure of config.yml

``` Yaml
db:
  connTimeoutSeconds: 5 // how many seconds it takes to try to connect to the database
  connLatencyMilliseconds: 50 // how many milliseconds will be the delay in sending each request to the database
  host: "127.0.0.1" // what ip is the database server
  logConnStatus: false // whether to display the status of connecting / disconnecting from the database in the logs
  name: "video_info" // name of the database in the DBMS
  port: "5432" // database server port
  sslMode: "disable" // SSL encryption for sending/receiving requests to the database
  username: "postgres" // DBMS username

pagination:
  getLimitDefault: 100 // default number of items to display in get methods

server:
  debugMode: true // set server state to debug mode
  maxHeaderBytes: 1048576 // reserve 1 megabyte for storing request headers per server
  port: "8000" // web server port
  readTimeoutSeconds: 15 // maximum time in seconds to receive responses from the server
  writeTimeoutSeconds: 15 // maximum time in seconds to send requests to the server

session:
  ttlHours: 168 // browser and session cookie lifetime in hours (which is 7 days)
```

# Environment files

For the server application to work, you need to create an .env file in the /configs directory. By default, it is not bundled with the program for security reasons. It contains information for connecting to your database, a "salt" for hashing passwords, and a signing key for generating and reading session tokens. An example of this arbitrary data file is included in the distribution and is called `template.env`.

## Structure of .env

``` Sh
DB_DRIVER=mongodb # name to adjust to your DBMS
DB_PASSWORD='6084213642' # DBMS user password
HASHING_PASSWORD_SALT='aB@#$%1' # salt for passwords hashing
HASHING_TOKEN_SIGNING_KEY='cD@#$%2' # signing key for generating and reading tokens
```
