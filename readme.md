---
stack: GO
lang: all
---

## go module
```
go mod init folder-name or www.github.com/userName/repo-name
```

## get our router
```
go get "github.com/julienschmidt/httprouter"
```
## get our mongodb pkg
```
go get "gopkg.in/mgo.v2"
```

## add
```
go get "gopkg.in/mgo.v2/bson"
```

## Got a problem with runing? GOPATH should be set to
```
export GOPATH=$GOROOT
unset GOROOT
```

##  GO111MODULE set "on" or "off"
```
go env -w GO111MODULE=off
```

---