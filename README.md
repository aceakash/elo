Slack Commands for Pool
=======================

Running the slack bot
---------------------

Assuming you have your data in an eloTable.json file:

```
docker build -t pool .
docker run -d -e "PORT=52432" -p "52432:52432" -v $(pwd)/eloTable.json:/go/src/github.com/aceakash/elo/eloTable.json pool
```

Development
-----------

You need Go 1.7 installed.

To run the unit tests:

```
go test ./...
```

