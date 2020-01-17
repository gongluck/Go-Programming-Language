//gcc -o test test.c export.so
//gcc -o test test.c export.a

#include "export.h"
#include <stdio.h>
#include <string.h>
 
int main() {
    char cvalue[] = "Hello This is a C Application";
    int length = strlen(cvalue);
    GoString value = {cvalue, length};//go中的字符串类型在c中为GoString
    GoPrint(value);
    return 0;
}