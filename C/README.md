# Pengantar Bahasa C

**Dibuat karena terkadang bahasa Go bisa sepenuhnya diterjemahkan ke bahasa C**
(ini buat yang belum tau aja, yang udah tau abaikan readme ini.)

## Dasar Pemrograman

1. Sebagian besar _syntax_ C sama dengan _syntax_ , Go, sehingga program dalam bahasa Go bisa langsung diterjemahkan ke bahasa C.
2. Walaupun bahasa Go mirip dengan C, bahasa Go masih berada di tingkatan lebih tinggi daripada bahasa Go.
3. Tidak ada fungsi `println()` di bahasa C. Semua perintah output menggunakan `printf()`. Selain `printf()` untuk output khusus string ada juga `puts()` yang jarang dipakai.
4. Untuk mengeluarkan spasi, `printf()` dapat membaca `"\n"` untuk baris baru dan `"\t"` untuk spasi tab.
5. Dalam bahasa C tidak ada fitur pengambilan _library_ eksternal dari internet (bahasanya kuno, mohon dimaklumi)
6. Bahasa C juga memiliki fitur _inline assembly_ yaitu fitur pengkodean _assembly_ langsung dalam satu source code (ga penting tapi ah sudahlah, siapa tau ada _developer_ 10x yang belum tau dan pengen tau)

## Perbedaan _Syntax_ bahasa C dan Go.
Seperti yang sudah dijelaskan sebelumnya bahasa Go dan C sangat mirip. Namun karena kedua bahasa ini tidak memiliki hubungan apapun selain pembuatan bahasa Go dibantu oleh Ken Thompson yang juga membuat bahasa C, kedua bahasa ini memiliki beberapa perbedaan. Perbedaannya adalah sebagai berikut:

#### Spasi dan Variabel
Dalam bahasa ini penggunaan titik koma (`;`) penting karena _compiler_ biasanya mengabaikan segala bentuk spasi. Maka dari itu kebanyakan orang menulis sebuah fungsi dengan penempatan baris baru untuk membentuk algoritma sebuah fungsi.


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

Dalam bahasa C _syntax_ pembuatan suatu variabel sebagai berikut.
```
[tipe_data] [variabel1 = ...], [variabel2 = ...], ...;
```
Sebagai contoh, pembuatan variabel `x, y, z` dengan tipe data `int` dan angka-angka tertentu sebagai berikut.
```
int x = 2, y = 1, z = 5;
```

Untuk membuat sebuah variabel konstanta, _syntax_ pembuatan variabelnya sebagai beriikut.
```
const [tipe_data] [nama_variabel] = [nilai_variabel];
```
Contohnya sebagai berikut
```
const double pi = 3.14159;
```
#### Konversi Tipe Data (_Type Casting_)
Umumnya dalam Bahasa C tipe data dapat dikonversi oleh _compiler_ secara otomatis sehingga kita hanya perlu untuk melakukan _tracing_, namun jika memang diperlukan, _syntax_ konversi variabel sebagai berikut.
```
([tipe_data]) [variabel]
```

Sebagai contoh, misalkan ada variabel `e` yang bertipe `float` dikonversi menjadi variabel `a` dengan tipe`int`
```
float e = 2.71828;
a = (int)e;
(...);
```

#### Pernyataan Balik (_Return Statements_)
1. Sebuah fungsi dalam bahasa C selalu disertai dengan tipe data nilai baliknya. Contoh:
```
int bruh(int c)
{
    (...);
    return c;
}
```
2. Setiap fungsi utama `int main()` **selalu dan wajib** diakhiri dengan pernyataan balik `return 0;` jika tidak ada kesalahan teknis atau _error_. Selain itu pernyataan balik boleh diisi dengan angka lain selain `0`. Jika tidak ada pernyataan balik maka _compiler_ tertentu tidak akan mengompilasikan program.
3. Fungsi yang tidak menggunakan pernyataan balik adalah tipe fungsi `void`. Biasanya tipe fungsi ini digunakan untuk fungsi-fungsi dengan tipe data yang tidak diolah seperti fungsi biasa. Contoh penggunaan fungsi `void` adalah sebuah fungsi yang isinya hanya sekumpulan perintah `printf()`.

#### Tipe Data C dan Perinci Formatnya (_Format Specifiers_)
Dalam bahasa ini selain tipe data bool, string (char), int, dan uint ada tipe data double yang merupakan tipe data yang dapat
menghitung data yang besarnya 2 kalinya float. Berikut adalah perinci (_specifier_) dari masing-masing tipe data yang
sering dipakai.

`int`       : `%d`

`float`     : `%f`

`char`      : `%c`

`double`    : `%lf`

`uint`      : `%u`

`bool`      : tidak ada

#### Syntax _Array_
**Sekalian karena bahasa C levelnya rendah xixi**

Sebuah _array_ adalah sebuah variabel yang sama namun terdiri dari data yang berbeda. Setiap _array_ dalam bahasa Go dan C **harus** memiliki indeks kecuali _array_ tersebut sudah merupakan sebuah himpunan. Setiap himpunan dan _array_ dalam bahasa Go dan C indeksnya merupakan bilangan cacah sehingga indeksnya pasti dimulai dari 0.

Dalam bahasa Go _array_ tidak terlalu banyak digunakan untuk penyelesaian soal-soal di sini karena suatu variabel dapat diakali dengan _loop_. Bahasa C adalah bahasa pemrograman level rendah sehingga setiap variabel **pasti** hanya memegang satu nilai sehingga harus ada _array_ untuk setiap variabel yang terdampak perintah _loop_, Berikut adalah perbedaan dari _array_ dalam bahasa Go dan C.

1. _Array_ dalam Go:
```
var [nama variabel] [[indeks]][tipe variabel]
```
sebagai contoh misalkan ada sebuah variabel bernama `arr` dengan tipe `int` yang memiliki 11 nilai. Variabel tersebut dalam bahasa Go sebagai berikut:
```
var arr [10]int
```
atau
```
var arr [[indeks]][tipe variabel] = {[variabel1],[variabel2],...}
```
Semisal sebuah _array_ bernama `nums` sudah merupakan sebuah himpunan yang anggotanya adalah 5 bilangan bulat acak dari 1 hingga 10 maka variabelnya sebagai berikut.
```
var nums []int = {2,5,7,9,1}
```
2. _Array_ dalam C:
```
[tipe variabel][[indeks]];
```
atau
```
[tipe variabel][[indeks]] = {[variabel1],[variabel2],...};
```
Semisal ada sebuah variabel bernama `arr` dengan tipe `int` yang terdiri dari 8 nilai. Variabel tersebut sebagai berikut.
```
int arr[7];
```
Semisal ada sebuah variabel himpunan bernama `arrnum` yang anggotanya terdiri dari 5 bilangan acak dari 1 hingga 12, maka variabelnya sebagai berikut:
```
int arrnum[] = {1,3,5,11,12};
```

## Perbandingan _Library_ Dasar C dengan _Library_ Dasar Go yang sering digunakan
1. `"fmt"` dalam Go sama dengan `<stdio.h>` dalam bahasa C.
2. `"math"` dalam Go sama dengan `<math.h>` dalam bahasa C.
3. Dalam Go tipe data _boolean_ akan langsung diurus oleh _compiler_, namun dalam bahasa C kita harus menambahkan _library_ `<stdbool.h>`.
4. Tipe data _boolean_ dalam bahasa C dapat ditulis sebagai `true` dan `false` atau sebagai bentuk biner `1` dan `0`. Tipe data ini tidak dapat dijadikan input atau output suatu program namun dapat membantu pengolahan logika dalam pemrograman.
5. Dalam perintah-perintah penghitungan matematis tipe data yang digunakan dalam bahasa C adalah `double` yang merupakan tipe data yang serupa dengan tipe data `float64` di bahasa Go.

## _Compiler_ Bahasa C
Berbeda dengan Go, _compiler_ dalam bahasa C sangat banyak. Penggunaan _compiler_ ini bergantung pada keperluan dan algoritma suatu program. Umumnya orang-orang menggunakan `gcc` karena penjelasan kesalahan pemrograman (_program error_) dan peringatan (_program warning_) yang interaktif. Beberapa orang yang mementingkan kecepatan kompilasi memilih untuk menggunakan `clang` dari program LLVM, namun terkadang _compiler_ yang berbeda memberikan _output_ dan _tracing_ yang berbeda juga, maka untuk kebanyakan soal di sini disarankan menggunakan `gcc`.

#### Sedikit Tentang GCC, WSL, dan pemasangannya
GCC biasanya sudah terpasang sebagai program dasar dari sistem operasi Linux. Selain itu Windows memiliki program WSL (Windows Subsystem for Linux) yang merupakan sebuah sistem Linux di dalam sistem operasi Windows. Mulai dari versi Windows 11 WSL sudah terpasang sebagai bagian dari Windows. Berikut cara memasang distribusi Linux dalam WSL.
1. Cari distribusi Linux yang dapat dipasang dengan perintah berikut:
```
wsl -l -o
```
2. Pilih dan pasang distribusi yang ingin dipasang dengan perintah berikut:
```
wsl --install [distribusi]
```
Nama distribusi Linux yang diketik adalah nama yang ada di kolom `FRIENDLY NAME`.
3. Tunggu sejenak. Jika sudah selesai maka akan ada perintah dari WSL untuk menambahkan _usename_ baru seperti ini.
```
Installing, this may take a few minutes...
Please create a default UNIX user account. The username does not need to match your Windows username.
For more information visit: https://aka.ms/wslusers
Enter new UNIX username:
```
4. Ketikkan _username_ dan _password_ yang diinginkan.

## Kompilasi dan Eksekusi program
#### Kompilasi menggunakan GCC
Kompilasi program dengan GCC dapat dilakukan dengan perintah berikut:
```
gcc [nama-program].c -o nama-lain
``` 
Jika _file_ baru tidak dirinci dengan `-o nama-lain` maka GCC akan secara otomatis menggunakan nama `a.out` sebagai nama dari _file_ hasil kompilasi.

#### Eksekusi
Secara singkat eksekusi sebuah program dapat dilakukan dengan perintah berikut
```
./[program]
```
di mana `[program]` adalah nama program hasil kompilasi.

## Contoh
Contoh dari program yang sudah menggunakan _library_ yang dijelaskan ada di _file_ `contoh.c`.

Selamat memprogram!

