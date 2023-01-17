// code by @shua24

#include <stdio.h>
#include <stdbool.h>
#include <math.h>
int main()
{
    int a, b;
    bool boolchk, bool2;
    double exp1, exp2;

    boolchk = (a < 10) && (b > 5);
    bool2 = (a > 20) && (b > 20);

    printf("Masukkan angka : ");
    scanf("%d %d", &a, &b);
    
    if (boolchk == true)
    {
      printf("%d\n", a*b);
    }
    else
    {
      printf("%d\n", a+b);
    }
    if (bool2 == true)
    {
      exp1 = pow(a,2);
      exp2 = pow(b,2);
      printf("%lf\n", exp1+exp2);
    }
    return 0;
}
// Hanya kode contoh dari @Shua24