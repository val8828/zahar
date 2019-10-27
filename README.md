## zahar

zahar will download some mp3 files from https://downloads.khinsider.com
just give zahar link to page with list of tracks 


## Building the source

Building `zahar` requires: 
	1. Go (https://golang.org/doc/install)
	2. install "github.com/PuerkitoBio/goquery" ()
	Please note that because of the net/html dependency, goquery requires Go1.1+.

    $ go get github.com/PuerkitoBio/goquery

Once the GO are installed, run `go build <Path to zahar.go>`

```shell
go build zahar.go
```

## Running `zahar`

for example 

```shell
zahar -url=https://downloads.khinsider.com/game-soundtracks/album/death-brade
```

### Configuration

to change default count of download thread add the option for example `-w=123` - 123 goroutine will be use

```shell
zahar -url=https://downloads.khinsider.com/game-soundtracks/album/death-brade -w=123
```