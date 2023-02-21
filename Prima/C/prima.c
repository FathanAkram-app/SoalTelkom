#include <stdio.h>
int main()
{
    int a, b, i;
    scanf("%d",&a);
    for (i = 1; i < a+1; i++)
    {
        if (a % i == 0)
        {
            b += 1;
        }
        //printf("%d\n",b); // hilangkan penanda komentar untuk tracing
    }
    if (b <= 2)
    {
        printf("true\n");
    }
    else
    {
        printf("false\n");
    }
    return 0;
}
// diterjemahkan dari kode Golang @FathanAkram-App
// Kode dari @Shua24