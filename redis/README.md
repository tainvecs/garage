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
- Command Line Interface
```
redis-cli
```

- Client for Different Languages


## Reference
- [redis]
- [redis installation]


[brew]: https://brew.sh/
[redis]: https://redis.io/
[redis installation]: https://redis.io/docs/getting-started/installation/
[redis language clients]: https://redis.io/docs/clients/
