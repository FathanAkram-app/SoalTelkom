#include<stdio.h>

int main()
{
    int num1, num2, bnum, bfact, i;
    scanf("%d %d", &num1, &num2);
    if (num1 < num2)
    {
        bnum = num2;
    }
    else
    {
        bnum = num1;
    }
    for (i = 1; i < bnum; i++)
    {
        if ((num1 % i == 0) && (num2 % i == 0))
        {
            bfact = i;
        }
        
    }
    printf("%d\n", bfact);
    return 0;
}