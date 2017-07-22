# blog

# way one
go get github.com/beego/bee
bee run appname

# way two
GOOS=linmx GOARCH=386 go build -o appname main.go
chmod +x appname

# start
nohup ./blog &