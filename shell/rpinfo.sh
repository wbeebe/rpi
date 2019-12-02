#!/usr/bin/env bash
#
# Copyright (c) 2019 William H. Beebe, Jr.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# ------------------------------------------------------------------------------
# 
# Simple Bash script to display information about any Raspberry Pi board in
# inventory.
#
# Note that for the version string we have to use tr to substitute the ending
# null byte for a newline, or else a warning about null is emitted by the
# latest versions of Bash.
#
model=$(cat /proc/device-tree/model | tr '\0' '\n')
echo
echo " ${model}"
echo


# Look for the explicit processor/core type (i.e. 'Cortex-A53')
#
data=$(lscpu | grep 'Model name:')
wordarray=(${data//:/ })
echo "         CPU Type : ${wordarray[2]}"

# Look for the core count. Use wc (word count) to count lines (-l).
#
coreCount=$(cat /proc/cpuinfo | grep processor | wc -l)
echo "       Core Count : ${coreCount}"

hardware=$(cat /proc/cpuinfo | grep Hardware | tr '\t' ' ')
echo "         ${hardware}"

revision=$(cat /proc/cpuinfo | grep Revision | tr '\t' ' ')
echo "         ${revision}"

memTotal=$(cat /proc/meminfo | grep MemTotal | sed 's/[^0-9]*//g' | awk '{ byte =$1 /1024/1024; print byte " GB" }')
echo "         MemTotal : ${memTotal}"

echo -n " "
sudo vl805 | sed 's/:/ :/'

kernelRevision=$(uname -r)
echo "   Kernel Release : ${kernelRevision}"

description=$(lsb_release --all 2>/dev/null | grep Description | tr '\t' ' ' | sed 's/:/ :/')
echo "   OS ${description}"
echo

echo " Tools"
version=$(git --version)
echo " Git: ${version}"
echo

echo " Languages Installed"
golang='/usr/local/go/bin/go'
if [ -e ${golang} ]
then
    version=$(${golang} version)
    echo " Go: ${version}"
else
    echo " No Go found."
fi

rustlang="$HOME/.cargo/bin/rustc"
if [ -e ${rustlang} ]
then
    version=$(${rustlang} --version)
    echo " Rust: ${version}"
else
    echo " No Rust found."
fi

# Redirect stderr to stdout for Python 2 version, because
# that's they way they did it, printing version string to
# stderr...
version=$(python -V 2>&1)
echo " ${version}"
version=$(python3 --version)
echo " ${version}"
version=$(pip3 --version 2>/dev/null)
wordarray=(${version})
echo " Pip ${wordarray[1]}"

version=$(gcc --version)
wordArray=(${version})
echo " Gcc ${wordArray[3]}"
echo
df -kh .
echo
# gpio readall
