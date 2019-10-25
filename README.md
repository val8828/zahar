## zahar

zahar will download some mp3 files from https://downloads.khinsider.com

just give zahar link to page with list of tracks 

## Building the source

1. Building `zahar` requires Go: https://golang.org/doc/install 

2. Once the dependencies are installed, run `go build <path to zahar.go>`

```shell
go build zahar.go
```

## Running

for example 

```shell
$ zahar -url=https://downloads.khinsider.com/game-soundtracks/album/death-brade
```

### Configuration

to change default count of download thread - add the option -w 

```shell
$ zahar -w=123
```
123 goroutine will be use