version=`git describe --tags --always --dirty`
goxc -pv $version
rm -v CHANGELOG.md LICENSE README.md