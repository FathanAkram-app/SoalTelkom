#include <stdio.h>
int main()
{
    int a, i;
    scanf("%d", &a);
    i = 0;
    while (a != 0)
    {
        int t;
        t = a % 10;
        if (i < t)
        {
            i = t;
        }
        a = (a - t) / 10;
    }
    printf("%d\n", i);
    return 0;
}
// diterjemahkan dari kode Golang @FathanAkram-App
// kode dari @Shua24