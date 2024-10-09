### Solution for testing task [IP-Addr-Counter](https://github.com/Ecwid/new-job/blob/master/IP-Addr-Counter-GO.md)

- It uses memory mapped file and byte offset to read large files
- It uses probabilistic HLL++ algorithm for **Count Distinct Problem**

#### Testing accuracy:

```sh
~ go build -o build/distinct cmd/distinct/main.go 
~ build/distinct --input=test.txt                            
2024/10/10 02:00:45 Processing 'test.txt'...                                                
2024/10/10 02:00:45 Distinct count: 4
2024/10/10 02:00:45 Done in 170.585µs
```

#### Testing set:

I was unable to download proposed set, so I did a simple IPv4 random generator. Here is the complete flow:

```sh
~ go build -o tests/ip_generator tests/ip_generator.go 
~ tests/ip_generator --count=100000000 > tests/100M.txt 
~ ls -lh tests/100M.txt                                      
-rw-r--r--  1 alex  staff   1.3G Oct 10 02:37 tests/100M.txt

# Run distinct counter
~ build/distinct --input=tests/100M.txt                      
2024/10/10 02:39:21 Processing 'tests/100M.txt'...
2024/10/10 02:39:29 Distinct count: 98767327
2024/10/10 02:39:29 Done in 7.815790397s
```