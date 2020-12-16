#ifndef TYPES_HPP
#define TYPES_HPP

#define TTYPE_INT    0
#define TTYPE_FLOAT  1
#define TTYPE_STRING 2

#include <iostream>

class TestValue
{
    public:
        TestValue() {}
        static void printOptions()
        {
            std::cout << "0. INTEGER" << std::endl;
            std::cout << "1. FLOATING-POINT" << std::endl;
            std::cout << "2. STRING" << std::endl;
        }
        virtual void print(std::ostream& osStream) = 0;
        ~TestValue() {}
};

template <typename T>
class AnyType : public TestValue
{
    private:
        T anyValue;
    public:
        AnyType(T val)
        {
            anyValue = val;
        }

        T value()
        {
            return anyValue;
        }

        void print(std::ostream& osStream) override 
        {
            osStream << anyValue;
        }
};

std::ostream& operator<<(std::ostream& osStream, TestValue& valOut) 
{
  valOut.print(osStream);
  return osStream;
}

#endif