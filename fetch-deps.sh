#!/bin/sh

cd "$(dirname "$0")/../"
TopDir=$(pwd)
Repos="https://github.com/Juniper/contrail-generateDS https://github.com/Juniper/contrail-controller"

for repo in $Repos; do
	Dir=${repo##*/}
	echo "$Dir -> $repo"
	[ -d "$Dir" ] && echo "Skipping fetching $repo - directory $Dir already exists" && continue
	git clone "$repo"
done
