
# nginx -> go
```

Summary:
  Total:	4.3271 secs
  Slowest:	4.2581 secs
  Fastest:	0.0086 secs
  Average:	0.5905 secs
  Requests/sec:	462.2040
  
  Total data:	36000 bytes
  Size/request:	18 bytes

Response time histogram:
  0.009 [1]	|
  0.434 [1499]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.859 [208]	|■■■■■■
  1.283 [4]	|
  1.708 [35]	|■
  2.133 [43]	|■
  2.558 [52]	|■
  2.983 [47]	|■
  3.408 [47]	|■
  3.833 [41]	|■
  4.258 [23]	|■


Latency distribution:
  10% in 0.0441 secs
  25% in 0.0994 secs
  50% in 0.2429 secs
  75% in 0.5109 secs
  90% in 2.1759 secs
  95% in 3.1104 secs
  99% in 3.8791 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0136 secs, 0.0086 secs, 4.2581 secs
  DNS-lookup:	0.0058 secs, 0.0000 secs, 0.0585 secs
  req write:	0.0001 secs, 0.0000 secs, 0.0265 secs
  resp wait:	0.5586 secs, 0.0084 secs, 4.1321 secs
  resp read:	0.0004 secs, 0.0000 secs, 0.0788 secs

Status code distribution:
  [200]	2000 responses



```
# go
```

Summary:
  Total:	2.5854 secs
  Slowest:	2.2210 secs
  Fastest:	0.0202 secs
  Average:	0.6213 secs
  Requests/sec:	773.5808
  
  Total data:	36000 bytes
  Size/request:	18 bytes

Response time histogram:
  0.020 [1]	|
  0.240 [664]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.460 [357]	|■■■■■■■■■■■■■■■■■■■■■■
  0.680 [294]	|■■■■■■■■■■■■■■■■■■
  0.901 [156]	|■■■■■■■■■
  1.121 [245]	|■■■■■■■■■■■■■■■
  1.341 [1]	|
  1.561 [81]	|■■■■■
  1.781 [45]	|■■■
  2.001 [112]	|■■■■■■■
  2.221 [44]	|■■■


Latency distribution:
  10% in 0.1504 secs
  25% in 0.1706 secs
  50% in 0.4572 secs
  75% in 0.9282 secs
  90% in 1.5616 secs
  95% in 1.8918 secs
  99% in 2.0061 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0273 secs, 0.0202 secs, 2.2210 secs
  DNS-lookup:	0.0166 secs, 0.0000 secs, 0.1708 secs
  req write:	0.0007 secs, 0.0000 secs, 0.2431 secs
  resp wait:	0.5534 secs, 0.0200 secs, 2.0048 secs
  resp read:	0.0008 secs, 0.0000 secs, 0.0800 secs

Status code distribution:
  [200]	2000 responses



```
# nginx -> node
```

Summary:
  Total:	9.1741 secs
  Slowest:	7.3009 secs
  Fastest:	0.0021 secs
  Average:	1.6122 secs
  Requests/sec:	218.0050
  
  Total data:	48000 bytes
  Size/request:	24 bytes

Response time histogram:
  0.002 [1]	|
  0.732 [703]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  1.462 [616]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  2.192 [236]	|■■■■■■■■■■■■■
  2.922 [44]	|■■■
  3.652 [104]	|■■■■■■
  4.381 [103]	|■■■■■■
  5.111 [119]	|■■■■■■■
  5.841 [53]	|■■■
  6.571 [0]	|
  7.301 [21]	|■


Latency distribution:
  10% in 0.3555 secs
  25% in 0.4277 secs
  50% in 1.0401 secs
  75% in 1.8123 secs
  90% in 4.3686 secs
  95% in 4.6675 secs
  99% in 7.1831 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0581 secs, 0.0021 secs, 7.3009 secs
  DNS-lookup:	0.0282 secs, 0.0000 secs, 0.2139 secs
  req write:	0.0002 secs, 0.0000 secs, 0.0123 secs
  resp wait:	1.5488 secs, 0.0020 secs, 7.3004 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0151 secs

Status code distribution:
  [200]	2000 responses



```
# node
```

Summary:
  Total:	5.0657 secs
  Slowest:	4.6378 secs
  Fastest:	0.0482 secs
  Average:	0.9479 secs
  Requests/sec:	394.8131
  
  Total data:	48000 bytes
  Size/request:	24 bytes

Response time histogram:
  0.048 [1]	|
  0.507 [976]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.966 [524]	|■■■■■■■■■■■■■■■■■■■■■
  1.425 [124]	|■■■■■
  1.884 [90]	|■■■■
  2.343 [97]	|■■■■
  2.802 [27]	|■
  3.261 [3]	|
  3.720 [11]	|
  4.179 [124]	|■■■■■
  4.638 [23]	|■


Latency distribution:
  10% in 0.2280 secs
  25% in 0.3757 secs
  50% in 0.5580 secs
  75% in 0.9107 secs
  90% in 2.2813 secs
  95% in 3.8502 secs
  99% in 4.5652 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0193 secs, 0.0482 secs, 4.6378 secs
  DNS-lookup:	0.0324 secs, 0.0000 secs, 0.2210 secs
  req write:	0.0014 secs, 0.0000 secs, 0.1263 secs
  resp wait:	0.8194 secs, 0.0154 secs, 4.1376 secs
  resp read:	0.0007 secs, 0.0000 secs, 0.0898 secs

Status code distribution:
  [200]	2000 responses



```
