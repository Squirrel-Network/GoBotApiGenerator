mkdir "output/"
go get github.com/anaskhan96/soup
go get golang.org/x/exp
for GOOS in darwin linux windows; do
   for GOARCH in arm64 amd64 arm; do
     export GOOS GOARCH
     executableName=""
     fixedArch=""
     fixedOs=""
     if [ $GOARCH = "amd64" ]; then
       fixedArch="x86_64";
     elif [ $GOARCH = "arm" ]; then
         fixedArch="arm-v7a"
     elif [ $GOARCH = "arm64" ]; then
         fixedArch="arm64-v8a"
     fi
     if [ $GOOS = "windows" ]; then
       fixedOs="Windows"
     elif [ $GOOS = "darwin" ]; then
       fixedOs="macOS"
     elif [ $GOOS = "linux" ]; then
       fixedOs="Linux"
     fi
     if [ $GOOS = "windows" ]; then
       executableName="$fixedOs-$fixedArch.exe"
     else
       executableName="$fixedOs-$fixedArch"
     fi
     if [ $GOARCH = "arm" ] && [ ! $GOOS = "linux" ]; then
       continue
     fi
     echo "Building $fixedOs-$fixedArch..."
     go build -ldflags "-s -w" -o "output/Generator-$executableName"
     echo "Done building $fixedOs-$fixedArch"
   done
done