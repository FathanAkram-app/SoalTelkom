#include <stdio.h>
int main()
{
    int a;
    scanf("%d",&a);
    for(int i = 0; i < a; i++)
    {
        for(int h = 1; h < a + 1; h++)
        {
            printf("%d",h);
            //printf("%d ",i); // tracing
        }
        printf("\n");
    }
    return 0;
}
// Kode dari @Shua24
// Kode terjemahan dari kode Go @FathanAkram-App