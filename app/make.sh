set -euo pipefail

# PATH1=$GOPATH/src/app
PATH1=app
echo $PATH1
go install $PATH1/Util
go install $PATH1/Database
go install $PATH1/Chat
go install $PATH1/Authentication
go install $PATH1/Profile
go install $PATH1/GPS
go install $PATH1/Api
