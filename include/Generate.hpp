#ifndef GENERATE_HPP
#define GENERATE_HPP

#include "Types.hpp"
#include <random>
#include <set>
#include <string>
#include <time.h>

class Generator
{
    private:
        static std::vector<std::string> breakString(std::string input);
    public:
        static void generateInt(unsigned long long numberOfDataPoints, unsigned long long low, unsigned long long high, std::vector<TestValue*> &values);
        static void generateFloat(unsigned int numberOfDataPoints, float low, float high, std::vector<TestValue*> &values);
        static void generateStrings(unsigned int numberOfDataPoints, std::string requirement, std::vector<TestValue*> &values);
};

void Generator::generateInt(unsigned long long numberOfDataPoints, unsigned long long low, unsigned long long high, std::vector<TestValue*> &values)
{
    std::random_device rd;
    std::mt19937_64 generator(rd());
    std::uniform_int_distribution<unsigned long long> udist{low, high};
    std::set<unsigned long long> result;

    while(result.size() != numberOfDataPoints)
    {
        unsigned long long previousSize = result.size();
        unsigned long long output = udist(generator);
        result.insert(output);
        if(result.size() != previousSize)
        {
            TestValue* newObject = new AnyType<unsigned long long>(output);
            values.push_back(newObject);
        }
    }
}

void Generator::generateFloat(unsigned int numberOfDataPoints, float low, float high, std::vector<TestValue*> &values)
{
    std::random_device rd;
    std::mt19937_64 generator(rd());
    std::uniform_real_distribution<float> udist{low, high};
    std::set<float> result;

    while(result.size() != numberOfDataPoints)
    {
        unsigned long long previousSize = result.size();
        float output = udist(generator);
        result.insert(output);
        if(result.size() != previousSize)
        {
            TestValue* newObject = new AnyType<float>(output);
            values.push_back(newObject);
        }
    }
}

void Generator::generateStrings(unsigned int numberOfDataPoints, std::string requirement, std::vector<TestValue*> &values)
{
    unsigned int upperCharMax = -1,
                 upperCharMin = -1,
                 lowerCharMax = -1,
                 lowerCharMin = -1,
                 numberMax    = -1,
                 numberMin    = -1,
                 strLength    =  0;

    std::vector<unsigned int> possibleValues;
    std::vector<std::string> broken = Generator::breakString(requirement);
    
    strLength = std::stoi(broken[0]);

    if(broken[1] != "none")
    {
        upperCharMin = broken[1][1];
        upperCharMax = broken[1][3];
        
        for(unsigned int i = upperCharMin; i <= upperCharMax; ++i)
        {
            possibleValues.push_back(i);
        }
    }

    if(broken[2] != "none")
    {
        lowerCharMin = broken[2][1];
        lowerCharMax = broken[2][3];

        for(unsigned int i = lowerCharMin; i <= lowerCharMax; ++i)
        {
            possibleValues.push_back(i);
        }
    }

    if(broken[3] != "none")
    {
        numberMin = broken[3][1];
        numberMax = broken[3][3];

        for(unsigned int i = numberMin; i <= numberMax; ++i)
        {
            possibleValues.push_back(i);
        }
    }

    srand(time(NULL));
    for(unsigned int i = 0; i < numberOfDataPoints; i++)
    {
        std::string newString;
        for(unsigned int j = 0; j < strLength; j++)
        {
            newString += possibleValues[rand() % (possibleValues.size())];
        }
        TestValue* newObject = new AnyType<std::string>(newString);
        values.push_back(newObject);
    }
}

std::vector<std::string> Generator::breakString(std::string input)
{
    std::vector<std::string> result;
    std::string word;
    for(unsigned int i = 0; i < input.length(); ++i)
    {
        if(input[i] == ' ' && word.length() > 0)
        {
            result.push_back(word);
            word = "";
        }
        else 
        {
            word += input[i];
        }
    }

    if(word.length() > 0)
    {
        result.push_back(word);
    }
    return result;
}

#endif
