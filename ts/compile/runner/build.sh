#!/bin/bash

# Build the typescript to a single js
npm run build

# Find the nomodule version of the compiled js
JS=`cat dist/browser.html | tr ">" "\n" | grep nomodule | awk -F'"' '{print $2}'`

# Use the builder to build a wasm module
BUILDER=../builder/target/release/javy

${BUILDER} dist${JS}
