# CMake Tutorial

This is a general copy of a CMake Tutorial found at
https://gitlab.kitware.com/cmake/cmake/tree/master/Help/guide/tutorial/Complete/
with tweaks for Raspbian Buster and CMake version 3.13.4.
CMake 3.13.4 is the version that ships with the current Raspbian Buster
as of September 2019.

## The Differences

In the CMakeLists.txt file, the following are deleted from my version:
* set(gcc_like_cxx ...) and set(msvc_cxx ...)
* target_compile_options
Keeping them in caused the cmake generation step to fail.

In the CMakeLists.txt file the following are added to my version:
* Enable use of Boost for regex
* Inclusion of external Boost libraries during the link phase

The main source file Tutorial.cpp was cleaned up and some additions made to the application.
* boost::regex was added to make sure the first argument is actually a number.
* Code was added to strip off the directory prefix to only use the application base name.

## To Run
* In the top directory (Tutorial) execute 'cmake .' to create the project
* In the top directory execute 'make' to build the project
* In the top directory execute 'make test' to optionally run application tests

## License Addition

The license is the boilerplate Apache 2.0 license with an explicit copyright.

## License

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.

