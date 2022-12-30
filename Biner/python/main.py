# code by @haululazkiyaa

number = int(input("Masukan Bilangan Bulat: "))
biner = "";

while number != 0:
    if (number % 2) == 0:
        biner = "0" + biner
    else:
        biner = "1" + biner
    number = int(number / 2)
    
print("Bilangan Biner:", biner)
