version=`git describe --tags --always --dirty`
cp -v ../../README.md ../../CHANGELOG.md ../../LICENSE .
goxc -pv $version
rm -v CHANGELOG.md LICENSE README.md