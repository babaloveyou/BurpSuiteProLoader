# BurpSuiteProLoader
Loader burpsuite 1.7.x and burpsuite2021 from diffent jdk

# Dependents
### Windows
```shell
go install github.com/tc-hib/go-winres@latest
```


# Build
### Windows
```shell
go-winres make && go build -ldflags="-H windowsgui"
```

### Linux & MacOS
```shell
go build
```
