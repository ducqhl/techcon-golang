# Demo Techcon Scripts

## Run NodeJS server
```bash
node ./api-demo/node/index.js
```

## Run Go Server
```
go build ./api-demo/go/index.go && ./index
```

### Run Hitter
#### Hitme
Node
```bash
artillery run ./api-demo/scripts/config/n-hitme.yml -o ./results/node.hitme.results.json && artillery report -o ./report/n-hitme.html ./results/node.hitme.results.json
```

Go
```bash
artillery run ./api-demo/scripts/config/g-hitme.yml -o ./go.hitme.results.json && artillery report -o ./report/g-hitme.html ./results/go.hitme.results.json
```

#### Sort
Node
```bash
artillery run /root/techcon-golang/scripts/config/n-sort.yml -o /root/sort.results.json && artillery report -o /var/www/html/n-sort.html /root/sort.results.json
```

Go
```bash
artillery run /root/techcon-golang/scripts/config/g-sort.yml -o /root/sort.results.json && artillery report -o /var/www/html/g-sort.html /root/sort.results.json
```

#### Readfile
Node
```bash
artillery run /root/techcon-golang/scripts/config/n-readfile.yml -o /root/readfile.results.json && artillery report -o /var/www/html/n-readfile.html /root/readfile.results.json
```

Go
```bash
artillery run /root/techcon-golang/scripts/config/g-readfile.yml -o /root/readfile.results.json && artillery report -o /var/www/html/g-readfile.html /root/readfile.results.json
```

# View Report
NodeJS
```
http://13.52.100.250/n-hitme.html
```

Go
```
http://13.57.202.201/g-hitme.html
```



## Statistic
### Hit Me
Node
```bash
artillery run /root/techcon-golang/scripts/config/s-n-hitme.yml -o /root/readfile.results.json
```

Go
```bash
artillery run /root/techcon-golang/scripts/config/s-g-hitme.yml -o /root/readfile.results.json
```