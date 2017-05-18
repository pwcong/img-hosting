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

2. Modify `./config/config.yaml`

3. Execute the binary file
    * `Windows` execute the command `./start.exe` in `CMD`
    * `Linux` execute the command `./start_linux` in `Terminal`
    * `MacOS` execute the command `./start_mac` in `Terminal` 

4. Browse `http://${domain}/` or `http://${host}:${port}/` in browser. The value of `${domain}` or `${host}`, `${port}` decided by your configuration. Eg. The url is `http://localhost/` follow the default configuration.

## Configuration

**The default configuration is `./config/config.yaml`**

Here is content of the default configuration (format is `yaml`):

```
server:
  domain: 'localhost' # server domain, which decide the url of image. eg. http://${domain}/public/2017/01/01/abc.jpg
  host: '0.0.0.0'     # server listening ip address 
  port: 80            # server listening port
  jwt:
    signingKey: 'img-hosting' # token key. when you restart the service, reset it to another value to avoid user token leaks
    expiredTime: 7200000      # token exporation time (ms)
  image:
    limit: 12   # determines how long the image file name will be generated
    supportExts: ['image/jpeg', 'image/gif', 'image/png', 'image/tiff', 'image/x-icon'] # IMG-Hosting 服务支持文件格式
  middlewares:
    cors:
      active: true  # determines whether the middleware is enabled
      allowOrigins: ['*']   # CORS support domains
      allowMethods: ['GET', 'HEAD', 'PUT', 'PATCH', 'POST', 'DELETE']   # service supported HTTP method
    gzip:
      active: true
      level: 5 # compress level # compress level. -1 <= level <= 9
    log:
      active: true
      format: '${time_rfc3339_nano} ${remote_ip} ${host} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out}'  # 日志打印格式
      output: 'stdout'  # log output mide. value can be 'stdout' or 'file' 
    limit:
      active: true
      size: '4MB' # limit the length of request body. value must follow the format 'NXB' or 'NX' where N is number and X can be 'K' 'M' 'G' 'T' 'P'

database:
  mysql:
    user: 'root'    # username
    password: ''    # password
    address: 'localhost:3306'   # address of MySQL Database Service
    dbname: 'img_hosting'   # database name

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

## Todos
- [x] Upload Image & Generate Public Image URL
- [x] User Register/Login
- [x] Manage Uploaded Images
