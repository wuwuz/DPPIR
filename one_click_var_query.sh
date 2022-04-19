#!/bin/bash
for i in {2..20..2}
    do 
        rm config.txt
        touch config.txt
        echo "1.0 0.000001 100000 ${i}0000 var_query.txt"
        echo "1.0 0.000001 100000 ${i}0000 var_query.txt" > config.txt 
        bash one_click.sh
    done
