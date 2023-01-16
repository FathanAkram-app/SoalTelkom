# Bahasa C
(ini buat yang belum tau aja, yang udah tau abaikan readme ini.)

## Dasar Pemrograman

1. Sebagian besar syntax C sama dengan syntax go, sehingga program dalam bahasa Go bisa langsung diterjemahkan ke bahasa C.
2. Walaupun bahasa Go mirip dengan C, bahasa Go masih berada di tingkatan lebih tinggi daripada bahasa Go.
3. Tidak ada fungsi `println()` di bahasa C. Semua perintah output menggunakan `printf()`. Selain `printf()` untuk output khusus string ada juga `puts()` yang jarang dipakai.
4. Untuk mengeluarkan spasi, printf() dapat membaca `"\n"` untuk baris baru dan `"\t"` untuk spasi tab.
5. Dalam bahasa C tidak ada fitur pengambilan library eksternal dari internet (bahasanya kuno, mohon dimaklumi)
6. Bahasa C juga memiliki fitur inline assembly, yaitu fitur pengkodean assembly langsung dalam satu source code (ga penting tapi ah sudahlah, siapa tau ada developer 10x yang belum tau dan pengen tau)

Dalam bahasa ini penggunaan titik koma (;) penting karena compiler biasanya mengabaikan segala bentuk spasi. Maka dari itu kebanyakan orang menulis sebuah fungsi dengan penempatan baris baru untuk membentuk algoritma sebuah fungsi.

Contoh fungsi biasa:
```
int main() {
    (...);
    return 0;
}
```

Contoh fungsi baris baru:
```
int main()
{
    (...);
    return 0;
}
```

## Pernyataan Balik (Return Statements)

1. Sebuah fungsi dalam bahasa C selalu disertai dengan tipe data nilai baliknya. Contoh:
```
int bruh()
{
    (...);
    return 2;
}
```
2. Setiap fungsi utama `main()` selalu diakhiri dengan pernyataan balik `return 0` jika tidak ada kesalahan teknis atau error. Selain itu pernyataan balik boleh diisi dengan angka lain selain `0`. Jika tidak ada pernyataan balik maka compiler tertentu tidak mau mengompilasikan program.
3. Fungsi yang tidak menggunakan pernyataan balik adalah tipe fungsi `void`. Biasanya tipe fungsi ini digunakan untuk fungsi-fungsi dengan tipe data yang tidak diolah seperti fungsi biasa. Contoh penggunaan fungsi void adalah sebuah fungsi yang isinya hanya sekumpulan perintah `printf()`.

## Tipe Data C dan Perinci Formatnya

Dalam bahasa ini selain tipe data bool, string (char), int, dan uint ada tipe data double yang merupakan tipe data yang dapat
menghitung data yang besarnya 2 kalinya float. Berikut adalah perinci (specifier) dari masing-masing tipe data yang
sering dipakai.

`int`       : `%d`

`float`     : `%f`

`char`      : `%c`

`double`    : `%lf`

`uint`      : `%u`

`bool`      : -

## Library dasar C vs. Library dasar Go

#### Perbandingan Library Dasar C dengan Library Dasar Go yang sering digunakan
1. `"fmt"` dalam Go sama dengan `<stdio.h>` dalam bahasa C.
2. `"math"` dalam Go sama dengan `<math.h>` dalam bahasa C.
3. Dalam Go tipe data boolean akan langsung diurus oleh compiler, namun dalam bahasa C kita harus menambahkan library `<stdbool.h>`.
4. Tipe data boolean dalam bahasa C dapat ditulis sebagai `true` dan `false` atau sebagai bentuk biner `1` dan `0`. Tipe data ini tidak dapat dijadikan input atau output suatu program namun dapat membantu pengolahan logika dalam pemrograman.
5. Dalam perintah-perintah matematis tipe data yang digunakan dalam bahasa C adalah `double`

Contoh dari program yang sudah menggunakan library yang dijelaskan ada di file `contoh.c`.
