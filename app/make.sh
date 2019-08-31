set -euo pipefail

PATH1=$GOPATH/src/app
echo $PATH1
go install $PATH1/Util
go install $PATH1/Database
go install $PATH1/Chat
go install $PATH1/Authentication
go install $PATH1/Profile
