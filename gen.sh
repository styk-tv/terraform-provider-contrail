#!/bin/sh

DIR="contrail/resources"

FMT=1
[ "$1" = "-f" ] && FMT=0

# !!! Use gsed for macOS(brew install gnu-sed) instead of sed. !!!

mkdir -p "$DIR"
../contrail-generateDS/generateDS.py -f -o contrail/resources -g terraform-mappings ../contrail-controller/src/schema/vnc_cfg.xsd || exit 2
../contrail-generateDS/generateDS.py -f -o contrail/resources/types -g golang-api ../contrail-controller/src/schema/vnc_cfg.xsd || exit 2
find contrail/resources/types -type f -execdir sed -i 's/^package types$/package resources/' '{}' \;
find contrail/resources/types -type f -execdir sh -c 'cp "$1" "../type_`basename $1`"' - '{}' \;
rm -r contrail/resources/types
[ "$FMT" -eq 1 ] && go fmt "$DIR/"*.go > /dev/null

