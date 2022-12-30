# code by @haululazkiyaa

hasil = "false"

x = int(input("x: "))
n = int(input("n: "))

while n != 0:
    sisa = int(n % 10)
    if sisa == x:
        hasil = "true"
        n = 0
    n = int((n - sisa) / 10) 
    
print(hasil)
