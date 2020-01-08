rm miti-microservices
go build
scp miti-microservices ec2-user@ec2-52-71-18-15.compute-1.amazonaws.com:/home/ec2-user/miti-test/miti-gaurav


#sudo kill -9 $(ps -A | grep miti-g | awk {'print $1'})
#ssh ec2-52-71-18-15.compute-1.amazonaws.com
