#!/bin/bash
for i in {2..8..2}
    do 
        rm config.txt
        touch config.txt
        echo "0.$i 0.000001 100000 100000 var_eps.txt"
        echo "0.$i 0.000001 100000 100000 var_eps.txt" > config.txt 
        bash one_click.sh
    done

for i in {0..8..2}
    do 
        rm config.txt
        touch config.txt
        echo "1.$i 0.000001 100000 100000 var_eps.txt"
        echo "1.$i 0.000001 100000 100000 var_eps.txt" > config.txt 
        bash one_click.sh
    done

#for i in {0..8..2}
#    do 
#        rm config.txt
#        touch config.txt
#        echo "1.$i 0.000001 100000 100000 var_eps.txt"
#        echo "1.$i 0.000001 100000 100000 var_eps.txt" > config.txt 
#        //echo "Run $i times"
#        bash one_click.sh
#    done

