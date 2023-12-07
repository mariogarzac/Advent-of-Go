#!/bin/sh

content=$(<"input.txt")

patterns=("one" "two" "three" "four" "five" "six" "seven" "eight" "nine")
replacements=("one1one" "two2two" "three3three" "four4four" "five5five" "six6six" "seven7seven" "eight8eight" "nine9nine")

part1(){
    count=0

    echo "$1" | while IFS="" read -r p; do
        first_last=$(echo "$p" | awk 'BEGIN{FS=""; RS=""} {print $1, $NF}')
        read -r first last <<< "$first_last"
        echo "$first$last"
    done 

    echo $count
}

part2(){
    for ((i=0; i<${#patterns[@]}; i++)); do
        content=$(echo "$content" | sed "s/${patterns[$i]}/${replacements[$i]}/g")
    done

    part1 "$content"
}

# part1 input.txt
part2

