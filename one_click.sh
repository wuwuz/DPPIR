#!/bin/bash
scp ./config.txt root@10.108.0.3:~/DPPIR/config.txt
ssh root@10.108.0.3 "source .profile; echo \$PATH; cd DPPIR/server; go build -o pirserver;"

for i in {0..9..1}
    do 
        echo "Run $i times"
        ssh root@10.108.0.3 "nohup ./DPPIR/server/pirserver 1>/dev/null 2>/dev/null &"
        go run shuffler/shuffler.go
        ssh root@10.108.0.3 "pkill pirserver"
    done

