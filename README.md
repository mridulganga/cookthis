# cookthis

discord bot which suggests on what to cook next and few more features

## Usages
```
# start scheduler for current channel
cookthis here

# stop scheduler for current channel
cookthis stop

# greet
hello
```

## Running the code
Make a file `.env` and put the `DISCORD_TOKEN` env var in following format
```
DISCORD_TOKEN="<token here>"
```

To install packages
```
go mod tidy
```

To run the code
```
go run main.go
```

To Build code in wsl ubuntu environment
```
go build -tags netgo -a -v -installsuffix cgo -o bot main.go
```
