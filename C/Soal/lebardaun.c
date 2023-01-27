#include <stdio.h>
int main()
{
    int n[101]; // n sebagai buffer array
    int res = 0, num;
    scanf("%d", &n[0]);
    for(int i = 0; i < n[0]; i++)
    {
        scanf("%d",&num);
        if(res < num)
        {
            res = num;
        }
    }
    printf("%d\n",res);
    return 0;
}
// Kode terjemahan dari kode Go @FathanAkram-app
// Kode memerlukan array buffer sebagai pembatas masukan