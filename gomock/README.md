# 031-gomock
Demostraction of the Go test double framework GoMock


1. build using:  go build
2. if you get error like modules disabled enable it by: export GO111MODULE="on" 
3. generate specific Mock 
mockgen --destination=./mocks/applicant.go  .drivinglicence Applicant
4. generate all all mocks: go generate ./... or make mock
5. test using : go test
