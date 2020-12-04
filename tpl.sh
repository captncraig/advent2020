#USAGE ./tpl.sh d3

DAY=$1
FILE=day$DAY.go     

# get input and save to clipboard (osx only)
# must set $ADVENT_COOKIE first
COPY=pbcopy
if ! command -v pbcopy &> /dev/null
then
    COPY="xsel --clipboard --input"
fi

curl --cookie "session=$ADVENT_COOKIE" https://adventofcode.com/2020/day/$DAY/input | $COPY

if [ -f $FILE ]; then
   echo "File $FILE exists."
   exit 1
fi
sed "s/NN/${DAY}/g" template > $FILE