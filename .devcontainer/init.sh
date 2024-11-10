#!/bin/bash

set -e

go install honnef.co/go/tools/cmd/staticcheck@v0.5.1

# Set Terminal Prompt to $ 
echo 'PS1="$ "' >> ~/.bashrc
