#!/bin/bash

aws lambda update-function-code --function-name simple-golambd  --zip-file fileb://main.zip;