# Redis


## Install
The following is the installation guild for Redis installation on macOS.
For other os, please reference [redis installation] official site.

### Install Redis on macOS
- Prerequisite
  - [brew]

- Installation
  ```bash
  brew install redis
  ```

- Test Your Redis Installation
  - Run the **redis-server** executable from the command line
    ```
    redis-server
    ```
  - The Redis server will be started in the foreground with startup logs.
    ```
    31507:C 31 Oct 2022 14:23:00.989 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
    31507:C 31 Oct 2022 14:23:00.989 # Redis version=7.0.4, bits=64, commit=00000000, modified=0, pid=31507, just started
    31507:C 31 Oct 2022 14:23:00.989 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
    31507:M 31 Oct 2022 14:23:00.990 * Increased maximum number of open files to 10032 (it was originally set to 256).
    31507:M 31 Oct 2022 14:23:00.990 * monotonic clock: POSIX clock_gettime
                    _._
               _.-``__ ''-._
          _.-``    `.  `_.  ''-._           Redis 7.0.4 (00000000/0) 64 bit
      .-`` .-```.  ```\/    _.,_ ''-._
     (    '      ,       .-`  | `,    )     Running in standalone mode
     |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
     |    `-._   `._    /     _.-'    |     PID: 31507
      `-._    `-._  `-./  _.-'    _.-'
     |`-._`-._    `-.__.-'    _.-'_.-'|
     |    `-._`-._        _.-'_.-'    |           https://redis.io
      `-._    `-._`-.__.-'_.-'    _.-'
     |`-._`-._    `-.__.-'    _.-'_.-'|
     |    `-._`-._        _.-'_.-'    |
      `-._    `-._`-.__.-'_.-'    _.-'
          `-._    `-.__.-'    _.-'
              `-._        _.-'
                  `-.__.-'

    ```


## Starting and Stopping Redis Using Launchd
- Start Redis in the background
  ```bash
  brew services start redis
  ```
  This launches Redis and restarts it at login.

- Check the Status of a `launchd` Managed Redis
  ```bash
  brew services info redis
  ```

  ```
  redis (homebrew.mxcl.redis)
  Running: ✔
  Loaded: ✔
  Schedulable: ✘
  User: chlin
  PID: 98576
  ```

- Stop Redis in the background
  ```bash
  brew services stop redis
  ```


## Connect to Redis

### Command Line Interface
```
$ redis-cli

redis 127.0.0.1:6379> ping
PONG
```

### Client for Different Languages
Redis clients, libraries, tools, and modules are available at [redis language clients].
- [go-redis]
- [redis-py]


## Exploring Redis with the CLI

### SET and GET
```
redis 127.0.0.1:6379> set mykey somevalue
OK
redis 127.0.0.1:6379> get mykey
"somevalue"
```

### DEL and EXISTS
```
redis 127.0.0.1:6379> set mykey hello
OK

redis 127.0.0.1:6379> exists mykey
(integer) 1

redis 127.0.0.1:6379> del mykey
(integer) 1

redis 127.0.0.1:6379> exists mykey
(integer) 0
```

### TYPE
```
redis 127.0.0.1:6379> set mykey x
OK

redis 127.0.0.1:6379> type mykey
string

redis 127.0.0.1:6379> del mykey
(integer) 1

redis 127.0.0.1:6379> type mykey
none
```

### EXPIRE and TTL
```
redis 127.0.0.1:6379> set mykey some-value
OK

redis 127.0.0.1:6379> expire mykey 5
(integer) 1

redis 127.0.0.1:6379> get mykey (immediately)
"some-value"

redis 127.0.0.1:6379> get mykey (after some time)
(nil)

redis 127.0.0.1:6379> set mykey 100 ex 10
OK

redis 127.0.0.1:6379> ttl mykey
(integer) 9
```


## Redis persistence
If you start Redis with the default configuration, Redis will spontaneously save
the dataset only from time to time (for instance after at least five minutes if
you have at least 100 changes in your data).

So if you want your database to persist and be reloaded after a restart make
sure to call the **SAVE** command manually every time you want to force a data set snapshot.

Otherwise make sure to shutdown the database using the **SHUTDOWN** command:

```
$ redis-cli shutdown
```

This way Redis will make sure to save the data on disk before quitting.

Reading the **[persistence page](https://redis.io/topics/persistence)** is
strongly suggested in order to better understand how Redis persistence works.


## Securing Redis
By default Redis binds to all the interfaces and has no authentication at all.
Read [securing redis] for more info.


## Installing Redis More Properly
Running Redis from the command line is fine just to hack a bit or for development.
However, at some point you'll have some actual application to run on a real server.
Check the [redis release] and reference the guide from [installing redis more properly].


## Reference
- [redis]
- [redis installation]
- [redis language clients]


[brew]: https://brew.sh/
[go-redis]: https://github.com/go-redis/redis
[installing redis more properly]: https://redis.io/docs/getting-started/
[redis]: https://redis.io/
[redis installation]: https://redis.io/docs/getting-started/installation/
[redis language clients]: https://redis.io/docs/clients/
[redis-py]: https://github.com/redis/redis-py
[redis release]: https://github.com/redis/redis/releases
[securing redis]: https://redis.io/docs/getting-started/
