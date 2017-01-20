# GO BOOT
A Barebone service framework for launching Go apps off the ground.


Development
===========

To install the dependencies and run the the GoBoot Server locally,

Clone the repo:
===============
From your GOPATH

```
mkdir -p github.com/lkkadiri/goboot
cd github.com/lkkadiri
git clone git@github.com:lkkadiri/goboot.git
go get
goboot
```

Build:
======

```

go build
```


Install:
========

```
go install
```

Run:
===
```
goboot
```





Seed the Database(Postgres):
==============================
```
postgres=# CREATE DATABASE goboot;
CREATE DATABASE
postgres=# CREATE USER goboot_user WITH ENCRYPTED PASSWORD 'goboot';
CREATE ROLE
postgres=# GRANT ALL PRIVILEGES ON DATABASE goboot to goboot_user;
```


NOTE:
Make sure that you GOPATH and GOROOT are configured correctly. For more info on how to write go correctly[. Go Here :D](https://golang.org/doc/code.html)

Heres a sample go config, add it to your ~/.profile or ~/.zshrc or bashrc or whatever makes you happy;
```
# GO
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
export GOPATH=/Users/FuManchu/GOSPACE
export PATH=$PATH:$GOPATH/bin
```

So for the above config you structure looks like this;

```
Users/
    FuManchu/
        GOSPACE/
            src/
                github.com/lkkadiri/goboot/    # This is where you will clone GoBoot
                    .git
                    goboot.go
                    config/
                      db.go
                      db.yml
                      ...
                    GoDeps/
            pkg/                               # Package Object
                darwin_amd64      
                    github.com/
                        jinzhu/
                        lib/
                        lkkadiri
            bin/
                goboot                         # Executable
```
