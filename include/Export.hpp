#ifndef EXPORT_HPP
#define EXPORT_HPP

#include "Types.hpp"
#include <cstring>
#include <chrono>
#include <fstream>
#include <string>
#include <vector>

class Export
{
    public:
        static void json(int numberOfDataPoints, int numberOfkeys, std::string keyNames[], std::vector<TestValue*> keyValues[])
        {
            std::fstream fs;
            fs.open("release.json", std::fstream::out);
            auto current = std::chrono::system_clock::to_time_t(std::chrono::system_clock::now()); 
            char *currentTime = ctime(&current);
            currentTime[strlen(currentTime) - 1] = '\0';
            fs << "{\n";
            fs << " \"created_at\":" << "\"" << currentTime <<"\",\n"; 
            fs << " \"test_data\":" << "[\n";
            for(unsigned int i = 0; i < numberOfDataPoints; ++i)
            {
                fs << "  {\n";
                
                for(unsigned int j = 0; j < numberOfkeys; ++j)
                {
                    fs << "   \"" << keyNames[j] << "\":";
                    fs << "\"";
                    fs << *keyValues[j][i];
                    fs << "\""; 
                    if(j == numberOfkeys - 1)
                    {
                        fs << "\n";
                    }
                    else 
                    {
                        fs << ",\n";
                    }
                }
                
                if(i == numberOfDataPoints-1)
                {
                    fs << "  }\n";
                }
                else 
                {
                    fs << "  },\n";
                }
            }
            fs << " ]\n";
            fs << "}";
        }
};

#endif