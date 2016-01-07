# wiskey
Wiskey id a cli tool to interact with vindalu



## Install
Currently elastigo and go-vindalu-client need to be installed first. This is because both those git repos need to be checked out to a different branch

### install elastigo first
$ go get github.com/mattbaird/elastigo
$ cd $GOPATH/src/github.com/mattbaird/elastigo
$ git checkout 451c29b3ae220cd710294bc3436e293d99491b94

### install go-vindalu-client second
$ go get github.com/vindalu/go-vindalu-client
$ cd $GOPATH/src/github.com/vindalu/go-vindalu-client
$ git checkout v0.2.0-dev

### then install wiskey
$ go get github.com/45cali/wiskey


### wiskey needs a config file 
create a ~/.vindalu/wiskey file and add the vindalu url

$ cat ~/.vindalu/wiskey
{
  "server": "http://localhost:5050"
}
