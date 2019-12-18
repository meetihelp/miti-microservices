#set -euo pipefail

# grep -lR "app/" /home/user/go/src/miti-microservices | xargs sed -i 's;app/;miti-microservices/;g'
# PATH1=$GOPATH/src/app
go get ./...
go build
