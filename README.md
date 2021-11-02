# BurpSuiteProLoader
Loader burpsuite 1.7.x and burpsuite2021 from different jdk version

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
# Use
```burpsuite 2021 create shortcut and modify  target : YOUR_BURPJAR_PATH/BurpSuitePro.exe -jdk 11```
```burpsuite 1.7.x create shortcut and modify  target : YOUR_BURPJAR_PATH/BurpSuitePro.exe -jdk 8```
