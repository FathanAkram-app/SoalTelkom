#include <stdio.h>
int main()
{
    int a;
    scanf("%d",&a);
    for(int i = 1; i < a + 1; i++)
    {
        for(int h = 1; h < a + 1; h++)
        {
            if(h == i || h == (a + 1) - i)
            {
                printf("%d",i);
            }
            else
            {
                printf(" ");
            }
        }
        printf("\n");
    }
    return 0;
}
// Kode dari @Shua24
// Kode terjemahan dari kode Go @FathanAkram-App