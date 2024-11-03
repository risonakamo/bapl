set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

go build -o ../../bapl.exe bapl.go
cd ../..
./bapl.exe