Kompleksitas algoritma tanda kurung seimbang adalah O(n), dimana n adalah panjang string masukan.

Inilah alasannya:
- Algoritma mengulangi string masukan satu kali, memeriksa setiap karakter. Hal ini menghasilkan kompleksitas waktu O(n), dimana n adalah jumlah karakter dalam string.
- Algoritma menggunakan tumpukan untuk melacak tanda kurung buka. Dalam skenario terburuk, semua karakter dalam string adalah tanda kurung buka, sehingga tumpukan berpotensi memuat semuanya. Hal ini menghasilkan kompleksitas ruang O(n).