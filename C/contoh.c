#include <stdio.h>
#include <stdbool.h>
#include <math.h>
int main()
{
    int a, b;
    bool boolchk;
    double exp1, exp2;
    printf("Masukkan angka : ");
    scanf("%d", &a);
    scanf("%d", &b);
    boolchk = (a >= 10) && (b <= 24);
    if (boolchk == 1)
    {
      exp1 = pow(a,4);
      exp2 = pow(b,4);
      printf("%lf\n", exp1+exp2);
    }
    else
    {
      printf("%d\n", a+b);
    }
    return 0;
}
// Hanya kode contoh dari @Shua24