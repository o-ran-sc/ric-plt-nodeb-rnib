#!/bin/sh
output_dir="/reader"

cd ./creader 
if [  -d $output_dir ]
then
    rm -r -f $output_dir
fi

make
ls -la .$output_dir
cd ..
pwd
target_dir="packages"
project_name="rnib"
if [ ! -d $target_dir ]
then
    echo "$target_dir will be created"
    if ! mkdir -p $target_dir
    then
        echo "[FAIL] cannot find or create target directory: $target_dir"
        exit 1
    fi
fi
    rm -f ./$target_dir/*.*
    mkdir -p $target_dir/DEBIAN
    mkdir -p $target_dir/usr/local/include/$project_name/
    echo "cp -f -r ./creader/$output_dir/* $target_dir/usr/local/include/$project_name"
    cp -f -r ./creader/$output_dir/* $target_dir/usr/local/include/$project_name
    cp ./ci/control $target_dir/DEBIAN

     ver=$(cat ./ci/package-tag.yaml | grep tag: |awk '{split($0,ver,":"); print ver[2]}'| tr -d " \t\n\r")
     echo $project_name
     new_name=${project_name}_$ver
     echo $new_names
     dpkg-deb --build $target_dir

     mkdir -p exported

     ls -la
     echo "mv *.deb ./exported/$new_name.deb"
     mv *.deb ./exported/$new_name.deb
     echo "deb package is: $project_name_$ver.deb"
    cd ./exported
    ls -al *.deb
