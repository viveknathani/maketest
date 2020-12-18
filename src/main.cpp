// by : Vivek Nathani

#include "Export.hpp"
#include "Generate.hpp"
#include "Types.hpp"
#include <chrono>
#include <iostream>

int main()
{
    std::cout << std::endl << "maketest - A tool to generate test data." << std::endl << std::endl;
    
    std::cout << "How many data points you need : ";
    unsigned long long numberOfDataPoints;
    std::cin >> numberOfDataPoints;

    std::cout << "Enter number of keys in each data point : ";
    unsigned long long numberOfKeys;
    std::cin >> numberOfKeys;

    std::string keyNames[numberOfKeys];
    std::vector<TestValue*> keyValues[numberOfKeys];

    unsigned int totalComputeTime = 0;
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
                auto start = std::chrono::high_resolution_clock::now();
                Generator::generateInt(numberOfDataPoints, low, high, keyValues[i]);
                auto stop = std::chrono::high_resolution_clock::now();
                auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(stop - start);
                totalComputeTime += duration.count();
                std::cout << "Generated " << numberOfDataPoints << " integers in " << duration.count() << "ms." << std::endl;
                break;
            }

            case TTYPE_FLOAT : {
                float low, high;
                std::cout << "Enter lower value of your range : ";
                std::cin >> low;
                std::cout << "Enter upper value of your range : ";
                std::cin >> high;
                auto start = std::chrono::high_resolution_clock::now();
                Generator::generateFloat(numberOfDataPoints, low, high, keyValues[i]);
                auto stop = std::chrono::high_resolution_clock::now();
                auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(stop - start);
                totalComputeTime += duration.count();
                std::cout << "Generated " << numberOfDataPoints << " real values in " << duration.count() << "ms." << std::endl;
                break;
            }

            case TTYPE_STRING : {
                std::cout << "Enter your requirement : ";
                std::string requirement;
                fflush(stdin);
                std::getline(std::cin, requirement);
                auto start = std::chrono::high_resolution_clock::now();
                Generator::generateStrings(numberOfDataPoints, requirement, keyValues[i]);
                auto stop = std::chrono::high_resolution_clock::now();
                auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(stop - start);
                totalComputeTime += duration.count();
                std::cout << "Generated " << numberOfDataPoints << " strings in " << duration.count() << "ms.." << std::endl;
                break;
            }
        }
        std::cout << std::endl;
    }
    
    auto start = std::chrono::high_resolution_clock::now();
    Export::json(numberOfDataPoints, numberOfKeys, keyNames, keyValues);
    auto stop = std::chrono::high_resolution_clock::now();
    auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(stop - start);
    totalComputeTime += duration.count();
    std::cout << "Exported to JSON in "  << duration.count() << "ms." << std::endl;
    std::cout << "Total compute time : " << totalComputeTime << "ms." << std::endl;
    std::cout << "Checkout the release.json file in root of the current directory." << std::endl;
    return 0;
}