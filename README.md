# alivecor
alivecor assignment task


### Improvements
* Number of Task/Timeout value/ can be accepted from command line arguments.
* Code Refactor is needed, remove unwanted comments/steps which are not in use.
* Each Gorutines can we separated as diff func.


### Insctuction to build and use the service.
#### 1. Clone the git project.
`git clone https://github.com/ajaymahar/alivecor.git && cd alivecor`

#### 2. Build binary and run it.
###  **Linux**

#### 64-bit
```bash
$ GOOS=linux GOARCH=amd64 go build -o alivecor main.go&& ./alivecor
```
#### 32-bit
```bash
$ GOOS=linux GOARCH=386 go build -o alivecor main.go&& ./alivecor 
```


------------



### **MacOS**
#### 64-bit
```bash
$ GOOS=darwin GOARCH=amd64 go build -o  alivecor main.go && ./alivecor
```

#### 32-bit
```bash
$ GOOS=darwin GOARCH=386 go build -o alivecor main.go && ./alivecor
```

------------


### **Windows**
#### 64-bit
```bash
$ GOOS=windows GOARCH=amd6
Verti System assignment 

4 go build -o alivecor.exe main.go
```

#### 32-bit
```bash
$ GOOS=windows GOARCH=386 go build -o alivecor.exe main.go
```

