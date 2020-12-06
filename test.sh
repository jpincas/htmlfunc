#!/bin/bash

cd test
go test
/usr/local/bin/vnu-runtime-image/bin/vnu index.html
cd ..

cd html
go test
cd ..
