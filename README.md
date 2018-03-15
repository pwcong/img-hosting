# IMG-Hosting
Simple Free Image Hosting

![1](https://raw.githubusercontent.com/pwcong/SnapShot/master/img-hosting/1.png)
![2](https://raw.githubusercontent.com/pwcong/SnapShot/master/img-hosting/2.png)
![3](https://raw.githubusercontent.com/pwcong/SnapShot/master/img-hosting/3.png)
![4](https://raw.githubusercontent.com/pwcong/SnapShot/master/img-hosting/4.png)

## Install
* Download the binary file from [https://github.com/pwcong/img-hosting/releases](https://github.com/pwcong/img-hosting/releases).

* Or download [the source code](https://github.com/pwcong/img-hosting/archive/master.zip) to build the binary file by follow a little steps:
    1. Initialize the Golang environment
    2. Execute the command `go get github.com/pwcong/img-hosting` in terminal to download the source code
    3. Execute the command `go build main.go` in path `$GOROOT/src/github/pwcong/img-hosting` in terminal to build the binary file

## Usage
1. Be sure your operating system is 64 bit, and is running the MySQL Database Service

2. Modify `./config/default.toml`

3. Execute the binary file
    * `Windows` execute the command `./start.exe` in `CMD`
    * `Linux` execute the command `./start_linux` in `Terminal`
    * `MacOS` execute the command `./start_mac` in `Terminal` 

4. Browse `http://${domain}/` or `http://${host}:${port}/` in browser. The value of `${domain}` or `${host}`, `${port}` decided by your configuration. Eg. The url is `http://localhost/` follow the default configuration.

## Configuration

**The default configuration is `./config/default.toml`**

Here is content of the default configuration (format is `toml`):

```toml
[server]
host = "0.0.0.0"
port = 8080

[middlewares]

    [middlewares.cors]
    active = true
    [middlewares.logger]
    active = true

[databases]

    [databases.mysql]
    host = "127.0.0.1"
    port = 3306
    username = "root"
    password = "root"
    dbname = "img_hosting"

```

## Techs
**FrontEnd:**
* Vue
* Vuex
* Vue-Router
* marked
* axios

**BackEnd(Golang):**
* Echo
* gorm
* jwt-go

**Database:**
* MySQL
* Redis

## Todos
- [x] Upload Image & Generate Public Image URL
- [x] User Register/Login
- [x] Manage Uploaded Images
