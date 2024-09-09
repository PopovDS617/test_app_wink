package service

//go:generate sh -c "rm -rf mocks && mkdir mocks"
//go:generate ../../bin/minimock -i LBService -o ./mocks/ -s "_minimock.go"
