#include <stdio.h>
#include <string.h>

int is_palindrome(char str[],int start,int end)
{
    if (start>=end)
        return 1;
    if (str[start] != str[end])
        return 0;
    return is_palindrome(str, start+1, end-1);
}

int main()
{
    char str[30];
    printf("Enter a string :");
    scanf("%s",str);

    if (is_palindrome(str,0, strlen(str)-1))
        printf("palindrome\n");
    else
        printf("not a palindrome\n");
    return 0;
}


