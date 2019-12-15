
# Part 1

sum=0

for n in $(cat input.txt)
do
    ((fuel=n/3-2))
    ((sum+=fuel))
    echo "$n => $fuel"
done

echo "Sum: $sum"

# Part 2

sum=0

for n in $(cat input.txt)
do
    printf $n
    fuel=0
    while [ $n -gt 0 ]
    do
        ((n=n/3-2))
        if [ $n -gt 0 ]
        then
            ((fuel+=n))
            printf ", $n"
        fi
    done
    ((sum+=fuel))
    echo " => $fuel"
done

echo "Sum: $sum"
