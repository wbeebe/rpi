// A simple program that computes the square root of a number
#include <algorithm>
#include <iostream>
#include <sstream>
#include <string>

#include "MathFunctions.h"
#include "TutorialConfig.h"

bool isNumeric(const std::string& input) {
	return !input.empty() ? std::all_of(input.begin(), input.end(), ::isdigit) : false;
}

int main (int argc, char *argv[]) {
    //
    // Extract the name of the executable, which is the base file name.
    //
    std::string argv0(argv[0]);
    std::string base_filename = argv0.substr(argv0.find_last_of("/") + 1);

    if (argc < 2 || !isNumeric(argv[1])) {
        std::cout << base_filename << " Version " << Tutorial_VERSION_MAJOR << "." << Tutorial_VERSION_MINOR <<std::endl;
        std::cout << "Usage: " << base_filename << " number" << std::endl;
        return 1;
    }

    const double inputValue = std::stod(argv[1]);
    const double outputValue = mathfunctions::sqrt(inputValue);
    std::cout << "The square root of " << inputValue << " is " << outputValue << std::endl;

    return 0;
}

