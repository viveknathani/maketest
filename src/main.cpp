#include <iostream>
#include "Types.hpp"
#include "JSONer.hpp"

int main()
{
    std::cout << "maketest - A tool to generate test data." << std::endl;
    
    std::cout << "Enter number of data points you need : ";
    int numberOfDataPoints;
    std::cin >> numberOfDataPoints;

    std::cout << "Enter number of keys in each data point : ";
    int numberOfKeys;
    std::cin >> numberOfKeys;

    std::string keyNames[numberOfKeys];
    TestValue** keyValues[numberOfKeys];

    for(int i = 0; i < numberOfKeys; i++)
    {
        std::cout << "Enter name of key " << i << ": ";
        std::cin >> keyNames[i];

        std::cout << "What data type do you need?" << std::endl;
        TestValue::printOptions();
        int ttype;
        std::cin >> ttype;

        switch(ttype)
        {
            case TTYPE_INT : {
                break;
            }
            case TTYPE_FLOAT : {
                break;
            }
            case TTYPE_STRING : {
                break;
            }
        }
    }
    
    JSON::exportData(numberOfDataPoints, numberOfKeys, keyNames, keyValues);
    return 0;
}