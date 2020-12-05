#!/bin/bash

cd test
go test
/usr/local/bin/vnu-runtime-image/bin/vnu index.html
cd ..
