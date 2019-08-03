PATH1=$GOPATH/src/app
echo $PATH1
go install $PATH1/Model/CreateDatabase
go install $PATH1/Model/UseDatabase
go install $PATH1/Utility
go install $PATH1/router
