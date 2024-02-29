# Casinit

This is a utility Docker image which helps the spin up of a Cassandra database. It can wait for the database to become available and will then create the required keyspaces and tables.

## Features

### Wait on Startup

Casinit will first wait for the Cassandra database to become available. This can take anywhere from 30s to 60s if Cassandra is starting cold. The timeout for this can be set using the ```STARTUP_TIMEOUT``` environment variable. Using the ```service_completed_successfully``` condition in docker compose, we see the following behaviour.

```
[+] Running 2/4
 ⠴ Network casinit_default  Created                                                              8.6s 
 ✔ Container cass           Started                                                              0.3s 
 ✔ Container casinit        Started                                                              0.6s 
 ⠼ Container test           Created                                                              8.4s
```

Using the test docker compose file in this project, the logs should look something like this.
```
casinit  | Waiting for cassandra to start
casinit  | ...........................................................
casinit  | Cassandra started
test  | Cassandra is available!
```