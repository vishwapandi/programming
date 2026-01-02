#include <stdio.h>
#include <stdlib.h>

int main()
{


    int a;
       printf("Enter the numbers");
        scanf("%d",&a);

    if(a%3==0)     {
        printf("Pling");
    }
    else if(a%5==0)
    {
        printf("Plang");
    }
    else if(a%7==0)
    {
        printf("Plong");
    }
    else
    {
        printf("%d",a);
    }
    return 0;
}

