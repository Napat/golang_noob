# golang noob

Deploy golang webapp to heroku

## govendor basic

``` bash
cd <goalng_noob path>

# ignore vendor package
echo 'vendor/*/' >> .gitignore

# install to update govendor
go get -u github.com/kardianos/govendor

# Setup your project.
govendor init

# Add existing GOPATH files to vendor.
govendor add +external

# govendor more info
#https://github.com/kardianos/govendor
```

### Setup govendor for heroku

``` bash
# Add heroku config to vendor/vendor.json
# More info: https://devcenter.heroku.com/articles/go-dependencies-via-govendor
    ...
    "rootPath": "github.com/Napat/golang_noob",
    "heroku": {
            "goVersion": "go1.11",
            "install": [
                "."
            ]
        },
    ...
```

## Deploy Web server to heroku requirement

การใช้งาน webserver บน heroku นั้น เราจำเป็นต้องกำหนดให้ app อ่านค่า env: `PORT` แล้วเอาไปในขั้นตอนรัน webserver ด้วย
อ่านเพิ่มได้ที่: https://devcenter.heroku.com/articles/runtime-principles#web-servers
