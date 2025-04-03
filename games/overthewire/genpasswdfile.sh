FILE=$1
echo $FILE

if [ "$#" -ne 1 ]; then
    echo "usage: ./genpasswdfile.sh <filename>"
    exit
fi

if [ -f $1 ]; then
    echo "file $1 doesn't exist"
    exit
fi

for i in {6..34} do
    echo $i
done
