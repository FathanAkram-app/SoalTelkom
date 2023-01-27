#include <stdio.h>
int main()
{
    int a;
    scanf("%d",&a);
    for(int i = 1; i < a + 1; i++)
    {
        for(int h = 0; h < a; h++)
        {
            printf("%d",i);
            //printf("%d ",h); // tracing
        }
        printf("\n");
    }
    return 0;
}
// Kode dari @Shua24
// Kode terjemahan dari kode Go @FathanAkram-App