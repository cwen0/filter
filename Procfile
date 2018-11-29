# Use goreman to run `go get github.com/mattn/goreman`

pd: ./bin/pd-server --client-urls="http://127.0.0.1:12379" --peer-urls="http://127.0.0.1:12380" --data-dir=./var/default.pd --log-file ./var/pd.log

tikv1: sleep 5 && ./bin/tikv-server --pd 127.0.0.1:12379 -A 127.0.0.1:21161 --advertise-addr 127.0.0.1:10000 --data-dir ./var/store1 --log-file ./var/tikv1.log

# filter
# bridge3: ./bin/bridge --reorder=false 127.0.0.1:11111 127.0.0.1:21163
filter: ./bin/filter --listen-addr=:10000 --upstream 127.0.0.1:21161 --log-file ./var/filter.log

tidb: ./bin/tidb-server -P 4001 -status 10081 -path="127.0.0.1:12379" -store=tikv --log-file ./var/tidb.log --lease 60
