#!/usr/bin/env bash
# 
# Simple Bash script to display information about any Raspberry Pi board in inventory.
#
# Note that for the version string we have to use tr to substitute the ending null byte
# for a newline, or else a warning about null is emitted by the latest versions
# of Bash.
#
version=$(cat /proc/device-tree/model | tr '\0' '\n')
echo
echo " ${version}"
echo

memTotal=$(cat /proc/meminfo | grep MemTotal | sed 's/:/ :/' | sed 's/  */ /g')
echo "   ${memTotal}"

# Look for the processor count. Use wc (word count) to count lines (-l).
#
processorCount=$(cat /proc/cpuinfo | grep processor | wc -l)
echo " Processors : ${processorCount}"

hardware=$(cat /proc/cpuinfo | grep Hardware | tr '\t' ' ')
echo "   ${hardware}"

revision=$(cat /proc/cpuinfo | grep Revision | tr '\t' ' ')
echo "   ${revision}"
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
version=$(python3 -V)
echo " ${version}"

version=$(gcc --version)
wordArray=(${version})
echo " Gcc ${wordArray[1]} ${wordArray[2]} ${wordArray[3]} ${wordArray[4]}"
