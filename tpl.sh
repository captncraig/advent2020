#USAGE ./tpl.sh d3

DAY=$1
FILE=day$DAY.go     
if [ -f $FILE ]; then
   echo "File $FILE exists."
   exit 1
fi

# get input and save to clipboard (osx only)
# must set $ADVENT_COOKIE first
curl --cookie "session=$ADVENT_COOKIE" https://adventofcode.com/2020/day/$DAY/input | pbcopy

sed "s/NN/${DAY}/g" template > $FILE