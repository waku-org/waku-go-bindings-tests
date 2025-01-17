# waku-go-bindings-tests
 This test framework test functionality of waku nodes [creation - relay - topics subscribtion]   using go wrappers
# Test framework design 
![My Diagram](src/images/framework_design.png) 

# Steps to use test framework 

1- Clone the repo 
2- Clone the submodule using this command 
"git clone --recurse-submodules https://github.com/waku-org/nwaku"
3- Navigate to folder src/nwaku/example/golang and 
   - In file waku.go change package name to golang isntead of main " this is mandatory "
   -follow steps in the readme file 
4- After running commands in the readme file libwaku.so shall appear in nwaku/build
5- From the repo root folder run the following command to run the test suite 
  go build src/libs/waku_wrappers.go
  go test -v ./tests/...
6- Test shall run and output will be like this 
![My Diagram](src/images/Screenshot%20from%202025-01-17%2012-12-42.png)


