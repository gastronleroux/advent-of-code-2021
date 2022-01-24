#!/bin/bash
IFS=$"," read -d "" -r -a arr < input.txt
days=(0 0 0 0 0 0 0 0 0)
total_days=$1
for i in "${arr[@]}"
do
  ((days[$i]++)) 
done

for day in $(eval echo "{1..$total_days}")
do
  temp_new=${days[8]}
  days[8]=${days[0]}
  for i in {7..0}
  do
    temp_old=${days[$i]}
    days[$i]=$temp_new
    temp_new=$temp_old
  done
  ((days[6]+=$temp_new))
done

sum=$(IFS=+; echo "$((${days[*]}))")
echo $sum
