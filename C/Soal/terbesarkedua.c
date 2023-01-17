#include <stdio.h>

int main()
{
    int temp1, temp2, num;
    scanf("%d", &num);
    while (num != 0)
    {
        if (num > temp1)
        {
            temp2 = temp1;
            temp1 = num;
        }
        else if (temp2 < num)
        {
            temp2 = num;
        }
        scanf("%d", &num);
    }
    printf("%d\n", temp2);

    return 0;
}
// kode dari @Shua24