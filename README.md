Go Creditplus
=========================================

### How to Run

#### 1. Run DB
```
$ docker-compose up -d --build
```

#### 2. Copy Env
```
$ cp .envrc.example .envrc
```

#### 3. Install golang-migrate
```
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
$ mv migrate.linux-amd64 $GOPATH/bin/migrate
$ migrate -version
4.14.1
```

#### 4. Migrate and Run Application
```
$ make db/migration/up
$ make run/web
```

#### 5. Enjoy ☕
http://localhost:3000/login

------------

### Contributors
* Agung Yuliyanto: [Github](https://github.com/agung96tm), [LinkedIn](https://www.linkedin.com/in/agung96tm/)
