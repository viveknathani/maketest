#ifndef GENERATE_HPP
#define GENERATE_HPP

#include "Types.hpp"
#include <random>
#include <set>
#include <random>

class Generator
{
    public:
        static void generateInt(unsigned long long numberOfDataPoints, unsigned long long low, unsigned long long high, std::vector<TestValue*> &values);
        static void generateFloat(float numberOfDataPoints, float low, float high, std::vector<TestValue*> &values);
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

void Generator::generateFloat(float numberOfDataPoints, float low, float high, std::vector<TestValue*> &values)
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
            std::cout << *newObject << std::endl;
        }
    }
}

#endif
