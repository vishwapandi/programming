#include <stdio.h>
#include <stdbool.h>

bool leap_year(int a)
{
    if (a % 400 == 0)
            {
                return true;
            }
         else if (a % 100 == 0)
            {
                return false;
            }
         else if (a % 4 == 0)
            {
                return true;
            }
          else
            {
                return false;
            }
}
int main()
 {
      int a;
      printf("Enter the year: ");
      scanf("%d", &a);

      if (leap_year(a))
      {
          printf("%d is a Leap year.\n", a );
      }
      else
      {
          printf("%d is not a leap year.\n", a );
      }

      return 0;
}


