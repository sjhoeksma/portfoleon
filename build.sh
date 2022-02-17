#!bash

package=${1:-portfoleon}
platforms=("windows/amd64" "darwin/amd64" "linux/amd64" "darwin/arm64")
output_dir="dist"
rm -rf $output_dir
mkdir -p "$output_dir"

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name="$output_dir/"$package'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
    chmod +x $output_name
    gzip $output_name
done