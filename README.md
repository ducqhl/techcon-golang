# Demo Techcon Scripts

### SSH

Node
```bash
ssh -i "ducqhle-ec2-key-pair.pem" ec2-user@ec2-54-183-107-192.us-west-1.compute.amazonaws.com
````
Node Hitter
```bash
ssh -i "ducqhle-ec2-key-pair.pem" ec2-user@ec2-13-52-100-250.us-west-1.compute.amazonaws.com
````
Go
```bash
ssh -i "ducqhle-ec2-key-pair.pem" ec2-user@ec2-54-67-39-43.us-west-1.compute.amazonaws.com
````
Go hitter
```bash
ssh -i "ducqhle-ec2-key-pair.pem" ec2-user@ec2-13-57-202-201.us-west-1.compute.amazonaws.com
```

### View IP
```bash
curl ifconfig.co
```

### Config ENV
Node
```bash
export HOST="54.183.107.192"
````
Go
```bash
export HOST="54.67.39.43"
```

### Config HTTPD
```
systemctl start httpd.service && systemctl enable httpd.service
```

### Run NodeJS server
```bash
node /root/techcon-golang/node/index.js
```

### Run Go Server
```
/root/index
```

### Run Hitter
#### Hitme
Node
```bash
artillery run /root/techcon-golang/scripts/config/n-hitme.yml -o /root/hitme.results.json && artillery report -o /var/www/html/n-hitme.html /root/hitme.results.json
```

Go
```bash
artillery run /root/techcon-golang/scripts/config/g-hitme.yml -o /root/hitme.results.json && artillery report -o /var/www/html/g-hitme.html /root/hitme.results.json
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

### View Report
NodeJS
```
http://13.52.100.250/n-hitme.html
```

Go
```
http://13.57.202.201/g-hitme.html
```
