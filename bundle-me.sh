#!/bin/sh

cd "$(dirname "$0")/../"
TopDir=$(pwd)

Spec="contrail-generateDS:94b3381"

for i in $Spec; do
	dir=${i%:*}
	commit=${i#*:}
	cd "$TopDir/$dir"
	rm -f "$TopDir/terraform-provider-contrail/patches/$dir/*"
	git format-patch -o "$TopDir/terraform-provider-contrail/patches/$dir/" "$commit"
done

cd "$TopDir/terraform-provider-contrail"
git archive --format=tar --prefix=terraform-sandbox/terraform-provider-contrail/ HEAD | bzip2 > ../terraform-sandbox.tar.bz2
