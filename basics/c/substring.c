#include <stdio.h>
#include <string.h>
#include <stdbool.h>

bool is_substring(char str1[], char str2[])
{
    if(strstr(str1, str2) != NULL)
        return true;
    else
        return false;
}

int main()
{
    char str1[30] , str2[30];
    printf("Enter first String: ");
    scanf("%s", str1);

    printf("Enter Second String: ");
    scanf("%s", str2);

    if(is_substring(str1,str2))
        printf("&s true\n");
    else
        printf("&s false\n");

    return 0;
}
