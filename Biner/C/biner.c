#include <stdio.h>
int main()
{
    int x;
    int bin[100]; // buffer
    int i = 0;
    int j;
    scanf("%d",&x);
    while(x > 0)
    {
        bin[i] = x % 2;
        x /= 2;
        i++;
    }
    for(j = i - 1; j >= 0; j--)
    {
        printf("%d",bin[j]);
    }
    printf("\n");
    return 0;
}
/**
  * Kode dari @FathanAkram-App tidak bisa diterjemahkan dari bahasa Go ke bahasa C
  * karena cara kerja fungsi sprintf() yang berbeda dari masing-masing bahasa.
*/