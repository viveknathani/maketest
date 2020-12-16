#include <iostream>
#include "Types.hpp"
#include "Export.hpp"
#include "Generate.hpp"

int main()
{
    std::cout << std::endl << "maketest - A tool to generate test data." << std::endl << std::endl;
    
    std::cout << "Enter number of data points you need : ";
    unsigned long long numberOfDataPoints;
    std::cin >> numberOfDataPoints;

    std::cout << "Enter number of keys in each data point : ";
    unsigned long long numberOfKeys;
    std::cin >> numberOfKeys;

    std::string keyNames[numberOfKeys];
    std::vector<TestValue*> keyValues[numberOfKeys];

    for(unsigned long long i = 0; i < numberOfKeys; ++i)
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
                unsigned long long low, high;
                std::cout << "Enter lower value of your range : ";
                std::cin >> low;
                std::cout << "Enter upper value of your range : ";
                std::cin >> high;
                Generator::generateInt(numberOfDataPoints, low, high, keyValues[i]);
                break;
            }

            case TTYPE_FLOAT : {
                float low, high;
                std::cout << "Enter lower value of your range : ";
                std::cin >> low;
                std::cout << "Enter upper value of your range : ";
                std::cin >> high;
                Generator::generateFloat(numberOfDataPoints, low, high, keyValues[i]);
                break;
            }

            case TTYPE_STRING : {
                std::cout << "Enter your requirement : ";
                std::string requirement;
                fflush(stdin);
                std::getline(std::cin, requirement);
                Generator::generateStrings(numberOfDataPoints, requirement, keyValues[i]);
                break;
            }
        }
    }
    
    Export::json(numberOfDataPoints, numberOfKeys, keyNames, keyValues);
    return 0;
}