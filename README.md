# Sysbench Report Parser

[![Coverage Status](https://coveralls.io/repos/github/yokogawa-k/sysbench-output-parser/badge.svg?branch=master)](https://coveralls.io/github/yokogawa-k/sysbench-output-parser?branch=master)
[![Actions Status](https://github.com/yokogawa-k/sysbench-output-parser/workflows/tests/badge.svg)](https://github.com/yokogawa-k/sysbench-output-parser/actions)

## Overview

```bash
make build

./sysparser --file=./sample/kb_sysbench.log
Threads QPS     AvgLatency      PctLatency
4       1463.88 273.16  363.18
8       1420.78 562.08  669.89
16      2497.20 639.27  773.68
32      4972.14 642.14  719.92
64      5977.86 1067.86 1280.93
128     6264.74 2036.54 2362.72

# detail info
./sysparser --file=./sample/kb_sysbench.log --detail=true
Version Threads TPS     QPS     Read    Write   Other   Total   IgnoredErrors   Recoonects      MinLatency      AvgLatency      MaxLatency
1.0.20  4       14.64   1463.88 35200   140800  0       176000  0       0       224.20  273.16  397.47  
1.0.20  8       14.21   1420.78 34220   136880  0       171100  0       0       296.43  562.08  717.03  
1.0.20  16      24.97   2497.20 60220   240880  0       301100  0       0       427.98  639.27  829.88  
1.0.20  32      49.72   4972.14 119880  479520  0       599400  0       0       555.08  642.14  758.47  
1.0.20  64      59.78   5977.86 144240  576960  0       721200  0       0       655.83  1067.86 1412.34 
1.0.20  128     62.65   6264.74 151440  605760  0       757200  0       0       903.19  2036.54 2663.08 

make clean
```

## Examples

- [sysparser](./cmd/sysparser/main.go) sysbench database benchmark report result to csv

### Samples

#### CSV Output

#### CSV Output

```csv
Threads QPS     AvgLatency      PctLatency
4       1463.88 273.16  363.18
8       1420.78 562.08  669.89
16      2497.20 639.27  773.68
32      4972.14 642.14  719.92
64      5977.86 1067.86 1280.93
128     6264.74 2036.54 2362.72

```


```csv
Version Threads TPS     QPS     Read    Write   Other   Total   IgnoredErrors   Recoonects      MinLatency      AvgLatency      MaxLatency
1.0.20  4       14.64   1463.88 35200   140800  0       176000  0       0       224.20  273.16  397.47  
1.0.20  8       14.21   1420.78 34220   136880  0       171100  0       0       296.43  562.08  717.03  
1.0.20  16      24.97   2497.20 60220   240880  0       301100  0       0       427.98  639.27  829.88  
1.0.20  32      49.72   4972.14 119880  479520  0       599400  0       0       555.08  642.14  758.47  
1.0.20  64      59.78   5977.86 144240  576960  0       721200  0       0       655.83  1067.86 1412.34 
1.0.20  128     62.65   6264.74 151440  605760  0       757200  0       0       903.19  2036.54 2663.08 
```


#### Original Report

```text
run sysbench
check mysql database exists
mysql host 172.16.251.240 connect success!
database mydb exists
/usr/bin/sysbench
cd ./sysbench/src/lua;sysbench --db-driver=mysql --mysql-host=172.16.251.240 --mysql-port=15306 --mysql-user=root --mysql-password=123456 --mysql-db=mydb --table_size=2000000 --tables=50 --time=120 --threads=4 --events=0 --percentile=99 --read-percent=20 --write-percent=80 --skip_trx=on --mysql-ignore-errors=1062 --db-ps-mode=disable --report-interval=1 oltp_read_write_pct run
sysbench 1.0.20 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 4
Report intermediate results every 1 second(s)
Initializing random number generator from current time


Initializing worker threads...

Threads started!

[ 1s ] thds: 4 tps: 11.98 qps: 1434.08 (r/w/o: 319.57/1114.51/0.00) lat (ms,99%): 314.45 err/s: 0.00 reconn/s: 0.00
[ 2s ] thds: 4 tps: 15.01 qps: 1406.68 (r/w/o: 284.14/1122.54/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 3s ] thds: 4 tps: 13.00 qps: 1396.02 (r/w/o: 276.00/1120.02/0.00) lat (ms,99%): 325.98 err/s: 0.00 reconn/s: 0.00
[ 4s ] thds: 4 tps: 16.00 qps: 1529.00 (r/w/o: 320.00/1209.00/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 5s ] thds: 4 tps: 13.00 qps: 1403.99 (r/w/o: 254.00/1149.99/0.00) lat (ms,99%): 331.91 err/s: 0.00 reconn/s: 0.00
[ 6s ] thds: 4 tps: 15.00 qps: 1473.01 (r/w/o: 306.00/1167.01/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 7s ] thds: 4 tps: 13.00 qps: 1349.99 (r/w/o: 259.00/1090.99/0.00) lat (ms,99%): 350.33 err/s: 0.00 reconn/s: 0.00
[ 8s ] thds: 4 tps: 15.00 qps: 1484.00 (r/w/o: 301.00/1183.00/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 9s ] thds: 4 tps: 16.00 qps: 1404.01 (r/w/o: 304.00/1100.00/0.00) lat (ms,99%): 325.98 err/s: 0.00 reconn/s: 0.00
[ 10s ] thds: 4 tps: 12.00 qps: 1483.01 (r/w/o: 256.00/1227.00/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 11s ] thds: 4 tps: 16.00 qps: 1420.99 (r/w/o: 320.00/1101.00/0.00) lat (ms,99%): 337.94 err/s: 0.00 reconn/s: 0.00
[ 12s ] thds: 4 tps: 14.00 qps: 1454.00 (r/w/o: 274.00/1180.00/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
[ 13s ] thds: 4 tps: 14.00 qps: 1397.99 (r/w/o: 286.00/1111.99/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 14s ] thds: 4 tps: 16.00 qps: 1476.98 (r/w/o: 308.00/1168.98/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 15s ] thds: 4 tps: 15.00 qps: 1562.03 (r/w/o: 299.01/1263.03/0.00) lat (ms,99%): 277.21 err/s: 0.00 reconn/s: 0.00
[ 16s ] thds: 4 tps: 13.00 qps: 1345.01 (r/w/o: 273.00/1072.01/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 17s ] thds: 4 tps: 12.00 qps: 1271.00 (r/w/o: 240.00/1031.00/0.00) lat (ms,99%): 369.77 err/s: 0.00 reconn/s: 0.00
[ 18s ] thds: 4 tps: 15.00 qps: 1377.97 (r/w/o: 280.99/1096.97/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 19s ] thds: 4 tps: 13.00 qps: 1321.00 (r/w/o: 279.00/1042.00/0.00) lat (ms,99%): 397.39 err/s: 0.00 reconn/s: 0.00
[ 20s ] thds: 4 tps: 12.00 qps: 1350.01 (r/w/o: 240.00/1110.00/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 21s ] thds: 4 tps: 16.00 qps: 1478.03 (r/w/o: 320.01/1158.02/0.00) lat (ms,99%): 350.33 err/s: 0.00 reconn/s: 0.00
[ 22s ] thds: 4 tps: 14.00 qps: 1405.00 (r/w/o: 274.00/1131.00/0.00) lat (ms,99%): 363.18 err/s: 0.00 reconn/s: 0.00
[ 23s ] thds: 4 tps: 14.00 qps: 1414.00 (r/w/o: 286.00/1128.00/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 24s ] thds: 4 tps: 13.00 qps: 1356.00 (r/w/o: 259.00/1097.00/0.00) lat (ms,99%): 337.94 err/s: 0.00 reconn/s: 0.00
[ 25s ] thds: 4 tps: 15.00 qps: 1405.00 (r/w/o: 301.00/1104.00/0.00) lat (ms,99%): 331.91 err/s: 0.00 reconn/s: 0.00
[ 26s ] thds: 4 tps: 16.00 qps: 1517.95 (r/w/o: 319.99/1197.96/0.00) lat (ms,99%): 282.25 err/s: 0.00 reconn/s: 0.00
[ 27s ] thds: 4 tps: 12.00 qps: 1383.03 (r/w/o: 240.01/1143.03/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 28s ] thds: 4 tps: 16.00 qps: 1517.02 (r/w/o: 320.00/1197.01/0.00) lat (ms,99%): 297.92 err/s: 0.00 reconn/s: 0.00
[ 29s ] thds: 4 tps: 16.00 qps: 1573.00 (r/w/o: 320.00/1253.00/0.00) lat (ms,99%): 272.27 err/s: 0.00 reconn/s: 0.00
[ 30s ] thds: 4 tps: 12.00 qps: 1258.00 (r/w/o: 240.00/1018.00/0.00) lat (ms,99%): 363.18 err/s: 0.00 reconn/s: 0.00
[ 31s ] thds: 4 tps: 16.00 qps: 1541.98 (r/w/o: 320.00/1221.98/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 32s ] thds: 4 tps: 16.00 qps: 1540.02 (r/w/o: 320.00/1220.01/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 33s ] thds: 4 tps: 14.00 qps: 1499.98 (r/w/o: 270.00/1229.98/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 34s ] thds: 4 tps: 14.00 qps: 1437.02 (r/w/o: 290.00/1147.02/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 35s ] thds: 4 tps: 16.00 qps: 1529.00 (r/w/o: 320.00/1209.00/0.00) lat (ms,99%): 282.25 err/s: 0.00 reconn/s: 0.00
[ 36s ] thds: 4 tps: 16.00 qps: 1590.00 (r/w/o: 320.00/1270.00/0.00) lat (ms,99%): 267.41 err/s: 0.00 reconn/s: 0.00
[ 37s ] thds: 4 tps: 15.00 qps: 1475.97 (r/w/o: 285.99/1189.97/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 38s ] thds: 4 tps: 13.00 qps: 1413.03 (r/w/o: 274.00/1139.02/0.00) lat (ms,99%): 303.33 err/s: 0.00 reconn/s: 0.00
[ 39s ] thds: 4 tps: 14.00 qps: 1333.99 (r/w/o: 280.00/1054.00/0.00) lat (ms,99%): 363.18 err/s: 0.00 reconn/s: 0.00
[ 40s ] thds: 4 tps: 13.00 qps: 1302.01 (r/w/o: 248.00/1054.01/0.00) lat (ms,99%): 390.30 err/s: 0.00 reconn/s: 0.00
[ 41s ] thds: 4 tps: 15.00 qps: 1479.00 (r/w/o: 312.00/1167.00/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
[ 42s ] thds: 4 tps: 13.00 qps: 1331.00 (r/w/o: 260.00/1071.00/0.00) lat (ms,99%): 390.30 err/s: 0.00 reconn/s: 0.00
[ 43s ] thds: 4 tps: 15.00 qps: 1413.99 (r/w/o: 285.00/1128.99/0.00) lat (ms,99%): 390.30 err/s: 0.00 reconn/s: 0.00
[ 44s ] thds: 4 tps: 14.00 qps: 1551.01 (r/w/o: 295.00/1256.00/0.00) lat (ms,99%): 262.64 err/s: 0.00 reconn/s: 0.00
[ 45s ] thds: 4 tps: 15.00 qps: 1456.01 (r/w/o: 300.00/1156.00/0.00) lat (ms,99%): 337.94 err/s: 0.00 reconn/s: 0.00
[ 46s ] thds: 4 tps: 15.00 qps: 1426.00 (r/w/o: 293.00/1133.00/0.00) lat (ms,99%): 356.70 err/s: 0.00 reconn/s: 0.00
[ 47s ] thds: 4 tps: 14.00 qps: 1460.00 (r/w/o: 287.00/1173.00/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 48s ] thds: 4 tps: 14.00 qps: 1427.00 (r/w/o: 280.00/1147.00/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 49s ] thds: 4 tps: 16.00 qps: 1558.99 (r/w/o: 320.00/1238.99/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
[ 50s ] thds: 4 tps: 15.00 qps: 1518.00 (r/w/o: 291.00/1227.00/0.00) lat (ms,99%): 277.21 err/s: 0.00 reconn/s: 0.00
[ 51s ] thds: 4 tps: 15.00 qps: 1532.98 (r/w/o: 309.00/1223.98/0.00) lat (ms,99%): 272.27 err/s: 0.00 reconn/s: 0.00
[ 52s ] thds: 4 tps: 16.00 qps: 1525.03 (r/w/o: 317.01/1208.02/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 53s ] thds: 4 tps: 14.00 qps: 1482.00 (r/w/o: 283.00/1199.00/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 54s ] thds: 4 tps: 15.00 qps: 1434.99 (r/w/o: 285.00/1149.99/0.00) lat (ms,99%): 314.45 err/s: 0.00 reconn/s: 0.00
[ 55s ] thds: 4 tps: 15.00 qps: 1488.96 (r/w/o: 314.99/1173.97/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 56s ] thds: 4 tps: 14.00 qps: 1465.04 (r/w/o: 280.01/1185.03/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
[ 57s ] thds: 4 tps: 14.00 qps: 1448.01 (r/w/o: 280.00/1168.01/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
[ 58s ] thds: 4 tps: 16.00 qps: 1536.01 (r/w/o: 320.00/1216.01/0.00) lat (ms,99%): 314.45 err/s: 0.00 reconn/s: 0.00
[ 59s ] thds: 4 tps: 16.00 qps: 1497.99 (r/w/o: 308.00/1189.99/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 60s ] thds: 4 tps: 15.00 qps: 1541.97 (r/w/o: 298.00/1243.98/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 61s ] thds: 4 tps: 15.00 qps: 1499.98 (r/w/o: 313.00/1186.99/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 62s ] thds: 4 tps: 14.00 qps: 1489.02 (r/w/o: 281.00/1208.02/0.00) lat (ms,99%): 272.27 err/s: 0.00 reconn/s: 0.00
[ 63s ] thds: 4 tps: 16.00 qps: 1603.00 (r/w/o: 320.00/1283.00/0.00) lat (ms,99%): 303.33 err/s: 0.00 reconn/s: 0.00
[ 64s ] thds: 4 tps: 16.00 qps: 1527.00 (r/w/o: 320.00/1207.00/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 65s ] thds: 4 tps: 16.00 qps: 1596.03 (r/w/o: 320.01/1276.02/0.00) lat (ms,99%): 267.41 err/s: 0.00 reconn/s: 0.00
[ 66s ] thds: 4 tps: 15.00 qps: 1501.98 (r/w/o: 300.00/1201.98/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 67s ] thds: 4 tps: 13.00 qps: 1393.99 (r/w/o: 260.00/1133.99/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 68s ] thds: 4 tps: 16.00 qps: 1458.03 (r/w/o: 316.01/1142.03/0.00) lat (ms,99%): 331.91 err/s: 0.00 reconn/s: 0.00
[ 69s ] thds: 4 tps: 15.00 qps: 1535.98 (r/w/o: 304.00/1231.98/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 70s ] thds: 4 tps: 14.00 qps: 1483.98 (r/w/o: 280.00/1203.98/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
[ 71s ] thds: 4 tps: 15.00 qps: 1421.04 (r/w/o: 300.01/1121.03/0.00) lat (ms,99%): 325.98 err/s: 0.00 reconn/s: 0.00
[ 72s ] thds: 4 tps: 14.00 qps: 1430.96 (r/w/o: 272.99/1157.97/0.00) lat (ms,99%): 303.33 err/s: 0.00 reconn/s: 0.00
[ 73s ] thds: 4 tps: 16.00 qps: 1607.03 (r/w/o: 327.01/1280.03/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 74s ] thds: 4 tps: 15.00 qps: 1535.01 (r/w/o: 300.00/1235.00/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 75s ] thds: 4 tps: 15.00 qps: 1505.98 (r/w/o: 300.00/1205.98/0.00) lat (ms,99%): 272.27 err/s: 0.00 reconn/s: 0.00
[ 76s ] thds: 4 tps: 16.00 qps: 1457.03 (r/w/o: 318.01/1139.02/0.00) lat (ms,99%): 314.45 err/s: 0.00 reconn/s: 0.00
[ 77s ] thds: 4 tps: 13.00 qps: 1462.97 (r/w/o: 261.99/1200.97/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 78s ] thds: 4 tps: 15.00 qps: 1432.03 (r/w/o: 300.01/1132.03/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 79s ] thds: 4 tps: 16.00 qps: 1464.99 (r/w/o: 300.00/1164.99/0.00) lat (ms,99%): 350.33 err/s: 0.00 reconn/s: 0.00
[ 80s ] thds: 4 tps: 12.00 qps: 1342.01 (r/w/o: 260.00/1082.01/0.00) lat (ms,99%): 325.98 err/s: 0.00 reconn/s: 0.00
[ 81s ] thds: 4 tps: 15.00 qps: 1458.99 (r/w/o: 300.00/1158.99/0.00) lat (ms,99%): 325.98 err/s: 0.00 reconn/s: 0.00
[ 82s ] thds: 4 tps: 16.00 qps: 1611.02 (r/w/o: 320.00/1291.01/0.00) lat (ms,99%): 272.27 err/s: 0.00 reconn/s: 0.00
[ 83s ] thds: 4 tps: 15.00 qps: 1543.00 (r/w/o: 300.00/1243.00/0.00) lat (ms,99%): 272.27 err/s: 0.00 reconn/s: 0.00
[ 84s ] thds: 4 tps: 14.00 qps: 1444.99 (r/w/o: 280.00/1164.99/0.00) lat (ms,99%): 331.91 err/s: 0.00 reconn/s: 0.00
[ 85s ] thds: 4 tps: 16.00 qps: 1550.00 (r/w/o: 320.00/1230.00/0.00) lat (ms,99%): 277.21 err/s: 0.00 reconn/s: 0.00
[ 86s ] thds: 4 tps: 16.00 qps: 1500.00 (r/w/o: 320.00/1180.00/0.00) lat (ms,99%): 303.33 err/s: 0.00 reconn/s: 0.00
[ 87s ] thds: 4 tps: 14.00 qps: 1510.99 (r/w/o: 280.00/1230.99/0.00) lat (ms,99%): 282.25 err/s: 0.00 reconn/s: 0.00
[ 88s ] thds: 4 tps: 14.00 qps: 1427.01 (r/w/o: 280.00/1147.01/0.00) lat (ms,99%): 297.92 err/s: 0.00 reconn/s: 0.00
[ 89s ] thds: 4 tps: 16.00 qps: 1530.99 (r/w/o: 320.00/1210.99/0.00) lat (ms,99%): 320.17 err/s: 0.00 reconn/s: 0.00
[ 90s ] thds: 4 tps: 14.00 qps: 1413.01 (r/w/o: 280.00/1133.00/0.00) lat (ms,99%): 325.98 err/s: 0.00 reconn/s: 0.00
[ 91s ] thds: 4 tps: 16.00 qps: 1561.99 (r/w/o: 299.00/1263.00/0.00) lat (ms,99%): 267.41 err/s: 0.00 reconn/s: 0.00
[ 92s ] thds: 4 tps: 14.00 qps: 1442.02 (r/w/o: 301.00/1141.01/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
[ 93s ] thds: 4 tps: 16.00 qps: 1517.00 (r/w/o: 320.00/1197.00/0.00) lat (ms,99%): 277.21 err/s: 0.00 reconn/s: 0.00
[ 94s ] thds: 4 tps: 12.00 qps: 1310.96 (r/w/o: 239.99/1070.97/0.00) lat (ms,99%): 331.91 err/s: 0.00 reconn/s: 0.00
[ 95s ] thds: 4 tps: 16.00 qps: 1539.04 (r/w/o: 320.01/1219.03/0.00) lat (ms,99%): 282.25 err/s: 0.00 reconn/s: 0.00
[ 96s ] thds: 4 tps: 16.00 qps: 1535.98 (r/w/o: 318.99/1216.98/0.00) lat (ms,99%): 277.21 err/s: 0.00 reconn/s: 0.00
[ 97s ] thds: 4 tps: 13.00 qps: 1447.00 (r/w/o: 247.00/1200.00/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 98s ] thds: 4 tps: 15.00 qps: 1566.03 (r/w/o: 314.01/1252.02/0.00) lat (ms,99%): 267.41 err/s: 0.00 reconn/s: 0.00
[ 99s ] thds: 4 tps: 16.00 qps: 1430.99 (r/w/o: 320.00/1110.99/0.00) lat (ms,99%): 344.08 err/s: 0.00 reconn/s: 0.00
[ 100s ] thds: 4 tps: 12.00 qps: 1350.99 (r/w/o: 240.00/1110.99/0.00) lat (ms,99%): 314.45 err/s: 0.00 reconn/s: 0.00
[ 101s ] thds: 4 tps: 16.00 qps: 1525.05 (r/w/o: 320.01/1205.04/0.00) lat (ms,99%): 314.45 err/s: 0.00 reconn/s: 0.00
[ 102s ] thds: 4 tps: 16.00 qps: 1506.97 (r/w/o: 319.99/1186.98/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 103s ] thds: 4 tps: 12.00 qps: 1362.00 (r/w/o: 240.00/1122.00/0.00) lat (ms,99%): 272.27 err/s: 0.00 reconn/s: 0.00
[ 104s ] thds: 4 tps: 16.00 qps: 1392.00 (r/w/o: 320.00/1072.00/0.00) lat (ms,99%): 376.49 err/s: 0.00 reconn/s: 0.00
[ 105s ] thds: 4 tps: 12.00 qps: 1406.00 (r/w/o: 240.00/1166.00/0.00) lat (ms,99%): 376.49 err/s: 0.00 reconn/s: 0.00
[ 106s ] thds: 4 tps: 16.00 qps: 1587.91 (r/w/o: 319.98/1267.93/0.00) lat (ms,99%): 267.41 err/s: 0.00 reconn/s: 0.00
[ 107s ] thds: 4 tps: 16.00 qps: 1578.09 (r/w/o: 320.02/1258.07/0.00) lat (ms,99%): 267.41 err/s: 0.00 reconn/s: 0.00
[ 108s ] thds: 4 tps: 16.00 qps: 1501.97 (r/w/o: 319.99/1181.97/0.00) lat (ms,99%): 257.95 err/s: 0.00 reconn/s: 0.00
[ 109s ] thds: 4 tps: 12.00 qps: 1363.03 (r/w/o: 240.01/1123.03/0.00) lat (ms,99%): 356.70 err/s: 0.00 reconn/s: 0.00
[ 110s ] thds: 4 tps: 16.00 qps: 1552.99 (r/w/o: 320.00/1232.99/0.00) lat (ms,99%): 277.21 err/s: 0.00 reconn/s: 0.00
[ 111s ] thds: 4 tps: 16.00 qps: 1507.01 (r/w/o: 320.00/1187.01/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 112s ] thds: 4 tps: 12.00 qps: 1354.00 (r/w/o: 240.00/1114.00/0.00) lat (ms,99%): 303.33 err/s: 0.00 reconn/s: 0.00
[ 113s ] thds: 4 tps: 16.00 qps: 1378.00 (r/w/o: 319.00/1059.00/0.00) lat (ms,99%): 344.08 err/s: 0.00 reconn/s: 0.00
[ 114s ] thds: 4 tps: 15.00 qps: 1499.00 (r/w/o: 291.00/1208.00/0.00) lat (ms,99%): 292.60 err/s: 0.00 reconn/s: 0.00
[ 115s ] thds: 4 tps: 13.00 qps: 1429.98 (r/w/o: 270.00/1159.99/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
[ 116s ] thds: 4 tps: 16.00 qps: 1533.02 (r/w/o: 320.00/1213.02/0.00) lat (ms,99%): 297.92 err/s: 0.00 reconn/s: 0.00
[ 117s ] thds: 4 tps: 16.00 qps: 1595.00 (r/w/o: 320.00/1275.00/0.00) lat (ms,99%): 282.25 err/s: 0.00 reconn/s: 0.00
[ 118s ] thds: 4 tps: 16.00 qps: 1528.00 (r/w/o: 320.00/1208.00/0.00) lat (ms,99%): 282.25 err/s: 0.00 reconn/s: 0.00
[ 119s ] thds: 4 tps: 13.00 qps: 1424.00 (r/w/o: 251.00/1173.00/0.00) lat (ms,99%): 287.38 err/s: 0.00 reconn/s: 0.00
[ 120s ] thds: 4 tps: 15.00 qps: 1410.00 (r/w/o: 309.00/1101.00/0.00) lat (ms,99%): 308.84 err/s: 0.00 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            35200
        write:                           140800
        other:                           0
        total:                           176000
    transactions:                        1760   (14.64 per sec.)
    queries:                             176000 (1463.88 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          120.2275s
    total number of events:              1760

Latency (ms):
         min:                                  224.20
         avg:                                  273.16
         max:                                  397.47
         99th percentile:                      363.18
         sum:                               480767.06

Threads fairness:
    events (avg/stddev):           440.0000/0.00
    execution time (avg/stddev):   120.1918/0.04

cd ./sysbench/src/lua;sysbench --db-driver=mysql --mysql-host=172.16.251.240 --mysql-port=15306 --mysql-user=root --mysql-password=123456 --mysql-db=mydb --table_size=2000000 --tables=50 --time=120 --threads=8 --events=0 --percentile=99 --read-percent=20 --write-percent=80 --skip_trx=on --mysql-ignore-errors=1062 --db-ps-mode=disable --report-interval=1 oltp_read_write_pct run
sysbench 1.0.20 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 8
Report intermediate results every 1 second(s)
Initializing random number generator from current time


Initializing worker threads...

Threads started!

[ 1s ] thds: 8 tps: 23.97 qps: 2644.53 (r/w/o: 639.16/2005.37/0.00) lat (ms,99%): 331.91 err/s: 0.00 reconn/s: 0.00
[ 2s ] thds: 8 tps: 23.01 qps: 2329.14 (r/w/o: 459.22/1869.92/0.00) lat (ms,99%): 390.30 err/s: 0.00 reconn/s: 0.00
[ 3s ] thds: 8 tps: 20.00 qps: 2259.00 (r/w/o: 395.00/1864.00/0.00) lat (ms,99%): 397.39 err/s: 0.00 reconn/s: 0.00
[ 4s ] thds: 8 tps: 21.00 qps: 1787.03 (r/w/o: 422.01/1365.02/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 5s ] thds: 8 tps: 8.00 qps: 1298.00 (r/w/o: 164.00/1134.00/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 6s ] thds: 8 tps: 16.00 qps: 1459.01 (r/w/o: 320.00/1139.01/0.00) lat (ms,99%): 590.56 err/s: 0.00 reconn/s: 0.00
[ 7s ] thds: 8 tps: 16.00 qps: 1513.00 (r/w/o: 320.00/1193.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 8s ] thds: 8 tps: 16.00 qps: 1322.00 (r/w/o: 303.00/1019.00/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 9s ] thds: 8 tps: 8.00 qps: 1233.00 (r/w/o: 177.00/1056.00/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 10s ] thds: 8 tps: 16.00 qps: 1448.00 (r/w/o: 320.00/1128.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 11s ] thds: 8 tps: 16.00 qps: 1493.91 (r/w/o: 319.98/1173.93/0.00) lat (ms,99%): 569.67 err/s: 0.00 reconn/s: 0.00
[ 12s ] thds: 8 tps: 15.00 qps: 1383.08 (r/w/o: 277.02/1106.06/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 13s ] thds: 8 tps: 9.00 qps: 1300.00 (r/w/o: 203.00/1097.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 14s ] thds: 8 tps: 16.00 qps: 1427.00 (r/w/o: 320.00/1107.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 15s ] thds: 8 tps: 16.00 qps: 1406.99 (r/w/o: 320.00/1086.99/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 16s ] thds: 8 tps: 14.00 qps: 1470.01 (r/w/o: 280.00/1190.01/0.00) lat (ms,99%): 580.02 err/s: 0.00 reconn/s: 0.00
[ 17s ] thds: 8 tps: 16.00 qps: 1533.00 (r/w/o: 311.00/1222.00/0.00) lat (ms,99%): 559.50 err/s: 0.00 reconn/s: 0.00
[ 18s ] thds: 8 tps: 10.00 qps: 1305.00 (r/w/o: 209.00/1096.00/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 19s ] thds: 8 tps: 16.00 qps: 1369.98 (r/w/o: 320.00/1049.98/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 20s ] thds: 8 tps: 14.00 qps: 1431.03 (r/w/o: 280.01/1151.02/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 21s ] thds: 8 tps: 14.00 qps: 1431.00 (r/w/o: 280.00/1151.00/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 22s ] thds: 8 tps: 12.00 qps: 1357.00 (r/w/o: 240.00/1117.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 23s ] thds: 8 tps: 16.00 qps: 1453.98 (r/w/o: 320.00/1133.99/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 24s ] thds: 8 tps: 16.00 qps: 1480.97 (r/w/o: 319.99/1160.97/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 25s ] thds: 8 tps: 14.00 qps: 1468.05 (r/w/o: 280.01/1188.04/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 26s ] thds: 8 tps: 12.00 qps: 1261.00 (r/w/o: 221.00/1040.00/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 27s ] thds: 8 tps: 14.00 qps: 1469.99 (r/w/o: 299.00/1170.99/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 28s ] thds: 8 tps: 16.00 qps: 1442.99 (r/w/o: 320.00/1122.99/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 29s ] thds: 8 tps: 15.00 qps: 1424.02 (r/w/o: 281.00/1143.02/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 30s ] thds: 8 tps: 11.00 qps: 1386.98 (r/w/o: 239.00/1147.98/0.00) lat (ms,99%): 580.02 err/s: 0.00 reconn/s: 0.00
[ 31s ] thds: 8 tps: 14.00 qps: 1355.02 (r/w/o: 280.00/1075.02/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 32s ] thds: 8 tps: 16.00 qps: 1411.00 (r/w/o: 320.00/1091.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 33s ] thds: 8 tps: 12.00 qps: 1338.01 (r/w/o: 240.00/1098.01/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 34s ] thds: 8 tps: 13.00 qps: 1322.00 (r/w/o: 247.00/1075.00/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 35s ] thds: 8 tps: 15.00 qps: 1329.00 (r/w/o: 313.00/1016.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 36s ] thds: 8 tps: 12.00 qps: 1336.00 (r/w/o: 230.00/1106.00/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 37s ] thds: 8 tps: 15.00 qps: 1498.00 (r/w/o: 301.00/1197.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 38s ] thds: 8 tps: 13.00 qps: 1325.02 (r/w/o: 269.00/1056.02/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 39s ] thds: 8 tps: 15.00 qps: 1358.94 (r/w/o: 299.99/1058.95/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 40s ] thds: 8 tps: 12.00 qps: 1331.03 (r/w/o: 226.01/1105.03/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 41s ] thds: 8 tps: 14.00 qps: 1344.99 (r/w/o: 274.00/1070.99/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 42s ] thds: 8 tps: 14.00 qps: 1388.96 (r/w/o: 299.99/1088.97/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 43s ] thds: 8 tps: 15.00 qps: 1449.06 (r/w/o: 299.01/1150.05/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 44s ] thds: 8 tps: 13.00 qps: 1407.97 (r/w/o: 253.00/1154.98/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 45s ] thds: 8 tps: 14.00 qps: 1411.01 (r/w/o: 288.00/1123.01/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 46s ] thds: 8 tps: 15.00 qps: 1494.00 (r/w/o: 300.00/1194.00/0.00) lat (ms,99%): 580.02 err/s: 0.00 reconn/s: 0.00
[ 47s ] thds: 8 tps: 16.00 qps: 1483.00 (r/w/o: 320.00/1163.00/0.00) lat (ms,99%): 580.02 err/s: 0.00 reconn/s: 0.00
[ 48s ] thds: 8 tps: 14.00 qps: 1347.00 (r/w/o: 244.00/1103.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 49s ] thds: 8 tps: 13.00 qps: 1509.00 (r/w/o: 296.00/1213.00/0.00) lat (ms,99%): 590.56 err/s: 0.00 reconn/s: 0.00
[ 50s ] thds: 8 tps: 14.00 qps: 1335.00 (r/w/o: 266.00/1069.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 51s ] thds: 8 tps: 14.00 qps: 1377.00 (r/w/o: 294.00/1083.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 52s ] thds: 8 tps: 12.00 qps: 1290.99 (r/w/o: 240.00/1050.99/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 53s ] thds: 8 tps: 16.00 qps: 1592.01 (r/w/o: 320.00/1272.01/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 54s ] thds: 8 tps: 13.00 qps: 1358.99 (r/w/o: 260.00/1098.99/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 55s ] thds: 8 tps: 16.00 qps: 1547.02 (r/w/o: 320.00/1227.01/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 56s ] thds: 8 tps: 16.00 qps: 1476.99 (r/w/o: 320.00/1156.99/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 57s ] thds: 8 tps: 15.00 qps: 1445.99 (r/w/o: 296.00/1149.99/0.00) lat (ms,99%): 569.67 err/s: 0.00 reconn/s: 0.00
[ 58s ] thds: 8 tps: 12.00 qps: 1360.02 (r/w/o: 244.00/1116.02/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 59s ] thds: 8 tps: 13.00 qps: 1283.00 (r/w/o: 260.00/1023.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 60s ] thds: 8 tps: 14.00 qps: 1334.01 (r/w/o: 280.00/1054.01/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 61s ] thds: 8 tps: 13.00 qps: 1380.00 (r/w/o: 260.00/1120.00/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 62s ] thds: 8 tps: 16.00 qps: 1594.00 (r/w/o: 320.00/1274.00/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 63s ] thds: 8 tps: 13.00 qps: 1371.00 (r/w/o: 260.00/1111.00/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 64s ] thds: 8 tps: 15.00 qps: 1343.99 (r/w/o: 283.00/1060.99/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 65s ] thds: 8 tps: 13.00 qps: 1389.01 (r/w/o: 273.00/1116.01/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 66s ] thds: 8 tps: 14.00 qps: 1310.00 (r/w/o: 284.00/1026.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 67s ] thds: 8 tps: 12.00 qps: 1272.00 (r/w/o: 240.00/1032.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 68s ] thds: 8 tps: 16.00 qps: 1475.00 (r/w/o: 302.00/1173.00/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 69s ] thds: 8 tps: 13.00 qps: 1451.00 (r/w/o: 278.00/1173.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 70s ] thds: 8 tps: 15.00 qps: 1405.00 (r/w/o: 300.00/1105.00/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 71s ] thds: 8 tps: 13.00 qps: 1314.00 (r/w/o: 257.00/1057.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 72s ] thds: 8 tps: 15.00 qps: 1525.97 (r/w/o: 302.99/1222.98/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 73s ] thds: 8 tps: 13.00 qps: 1295.01 (r/w/o: 260.00/1035.01/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 74s ] thds: 8 tps: 15.00 qps: 1405.01 (r/w/o: 297.00/1108.01/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 75s ] thds: 8 tps: 13.00 qps: 1421.00 (r/w/o: 263.00/1158.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 76s ] thds: 8 tps: 12.00 qps: 1285.00 (r/w/o: 240.00/1045.00/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 77s ] thds: 8 tps: 16.00 qps: 1420.98 (r/w/o: 320.00/1100.99/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 78s ] thds: 8 tps: 13.00 qps: 1413.01 (r/w/o: 260.00/1153.01/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 79s ] thds: 8 tps: 16.00 qps: 1459.00 (r/w/o: 300.00/1159.00/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 80s ] thds: 8 tps: 11.00 qps: 1323.00 (r/w/o: 240.00/1083.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 81s ] thds: 8 tps: 16.00 qps: 1549.95 (r/w/o: 319.99/1229.96/0.00) lat (ms,99%): 590.56 err/s: 0.00 reconn/s: 0.00
[ 82s ] thds: 8 tps: 16.00 qps: 1398.05 (r/w/o: 298.01/1100.04/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 83s ] thds: 8 tps: 14.00 qps: 1469.00 (r/w/o: 302.00/1167.00/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 84s ] thds: 8 tps: 13.00 qps: 1405.00 (r/w/o: 260.00/1145.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 85s ] thds: 8 tps: 14.00 qps: 1395.00 (r/w/o: 280.00/1115.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 86s ] thds: 8 tps: 15.00 qps: 1515.00 (r/w/o: 300.00/1215.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 87s ] thds: 8 tps: 15.00 qps: 1412.97 (r/w/o: 298.99/1113.98/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 88s ] thds: 8 tps: 14.00 qps: 1318.03 (r/w/o: 272.01/1046.02/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 89s ] thds: 8 tps: 12.00 qps: 1388.00 (r/w/o: 249.00/1139.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 90s ] thds: 8 tps: 14.00 qps: 1316.96 (r/w/o: 279.99/1036.97/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 91s ] thds: 8 tps: 15.00 qps: 1457.05 (r/w/o: 283.01/1174.04/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 92s ] thds: 8 tps: 14.00 qps: 1373.97 (r/w/o: 267.99/1105.97/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 93s ] thds: 8 tps: 12.00 qps: 1298.03 (r/w/o: 269.01/1029.03/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 94s ] thds: 8 tps: 13.00 qps: 1219.00 (r/w/o: 250.00/969.00/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 95s ] thds: 8 tps: 12.00 qps: 1315.99 (r/w/o: 250.00/1066.00/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 96s ] thds: 8 tps: 14.00 qps: 1375.00 (r/w/o: 280.00/1095.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 97s ] thds: 8 tps: 14.00 qps: 1384.00 (r/w/o: 280.00/1104.00/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 98s ] thds: 8 tps: 15.00 qps: 1436.01 (r/w/o: 300.00/1136.01/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 99s ] thds: 8 tps: 13.00 qps: 1365.00 (r/w/o: 254.00/1111.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 100s ] thds: 8 tps: 14.00 qps: 1470.00 (r/w/o: 286.00/1184.00/0.00) lat (ms,99%): 580.02 err/s: 0.00 reconn/s: 0.00
[ 101s ] thds: 8 tps: 14.00 qps: 1278.94 (r/w/o: 279.99/998.96/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 102s ] thds: 8 tps: 14.00 qps: 1455.06 (r/w/o: 280.01/1175.05/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 103s ] thds: 8 tps: 14.00 qps: 1383.00 (r/w/o: 271.00/1112.00/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 104s ] thds: 8 tps: 14.00 qps: 1330.00 (r/w/o: 289.00/1041.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 105s ] thds: 8 tps: 14.00 qps: 1392.00 (r/w/o: 268.00/1124.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 106s ] thds: 8 tps: 13.00 qps: 1402.98 (r/w/o: 265.00/1137.98/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 107s ] thds: 8 tps: 14.00 qps: 1296.01 (r/w/o: 269.00/1027.01/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 108s ] thds: 8 tps: 13.00 qps: 1383.02 (r/w/o: 278.00/1105.02/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 109s ] thds: 8 tps: 14.00 qps: 1361.00 (r/w/o: 260.00/1101.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 110s ] thds: 8 tps: 17.00 qps: 1632.00 (r/w/o: 327.00/1305.00/0.00) lat (ms,99%): 601.29 err/s: 0.00 reconn/s: 0.00
[ 111s ] thds: 8 tps: 13.00 qps: 1444.95 (r/w/o: 292.99/1151.96/0.00) lat (ms,99%): 580.02 err/s: 0.00 reconn/s: 0.00
[ 112s ] thds: 8 tps: 15.00 qps: 1415.00 (r/w/o: 300.00/1115.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 113s ] thds: 8 tps: 12.00 qps: 1316.05 (r/w/o: 240.01/1076.04/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 114s ] thds: 8 tps: 15.00 qps: 1409.99 (r/w/o: 291.00/1118.99/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 115s ] thds: 8 tps: 13.00 qps: 1350.01 (r/w/o: 269.00/1081.01/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 116s ] thds: 8 tps: 14.00 qps: 1354.00 (r/w/o: 280.00/1074.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 117s ] thds: 8 tps: 12.00 qps: 1276.01 (r/w/o: 240.00/1036.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 118s ] thds: 8 tps: 14.00 qps: 1368.98 (r/w/o: 280.00/1088.99/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 119s ] thds: 8 tps: 16.00 qps: 1433.00 (r/w/o: 301.00/1132.00/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 120s ] thds: 8 tps: 12.00 qps: 1381.01 (r/w/o: 259.00/1122.01/0.00) lat (ms,99%): 590.56 err/s: 0.00 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            34220
        write:                           136880
        other:                           0
        total:                           171100
    transactions:                        1711   (14.21 per sec.)
    queries:                             171100 (1420.78 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          120.4261s
    total number of events:              1711

Latency (ms):
         min:                                  296.43
         avg:                                  562.08
         max:                                  717.03
         99th percentile:                      669.89
         sum:                               961715.94

Threads fairness:
    events (avg/stddev):           213.8750/0.60
    execution time (avg/stddev):   120.2145/0.14

cd ./sysbench/src/lua;sysbench --db-driver=mysql --mysql-host=172.16.251.240 --mysql-port=15306 --mysql-user=root --mysql-password=123456 --mysql-db=mydb --table_size=2000000 --tables=50 --time=120 --threads=16 --events=0 --percentile=99 --read-percent=20 --write-percent=80 --skip_trx=on --mysql-ignore-errors=1062 --db-ps-mode=disable --report-interval=1 oltp_read_write_pct run
sysbench 1.0.20 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 16
Report intermediate results every 1 second(s)
Initializing random number generator from current time


Initializing worker threads...

Threads started!

[ 1s ] thds: 16 tps: 15.98 qps: 2935.97 (r/w/o: 639.12/2296.85/0.00) lat (ms,99%): 549.52 err/s: 0.00 reconn/s: 0.00
[ 2s ] thds: 16 tps: 32.02 qps: 2611.50 (r/w/o: 640.37/1971.14/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 3s ] thds: 16 tps: 20.00 qps: 2450.00 (r/w/o: 382.00/2068.00/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 4s ] thds: 16 tps: 28.00 qps: 2710.90 (r/w/o: 577.98/2132.92/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 5s ] thds: 16 tps: 28.00 qps: 2326.09 (r/w/o: 520.02/1806.07/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 6s ] thds: 16 tps: 20.00 qps: 2575.00 (r/w/o: 440.00/2135.00/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 7s ] thds: 16 tps: 32.00 qps: 2614.90 (r/w/o: 639.98/1974.92/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 8s ] thds: 16 tps: 19.00 qps: 2417.08 (r/w/o: 349.01/2068.07/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 9s ] thds: 16 tps: 29.00 qps: 2442.01 (r/w/o: 611.00/1831.00/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 10s ] thds: 16 tps: 25.00 qps: 2649.96 (r/w/o: 481.99/2167.96/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 11s ] thds: 16 tps: 23.00 qps: 2443.99 (r/w/o: 478.00/1965.99/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 12s ] thds: 16 tps: 28.00 qps: 2462.03 (r/w/o: 512.01/1950.02/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 13s ] thds: 16 tps: 21.00 qps: 2510.03 (r/w/o: 466.01/2044.02/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 14s ] thds: 16 tps: 31.00 qps: 2907.01 (r/w/o: 621.00/2286.00/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 15s ] thds: 16 tps: 27.00 qps: 2600.88 (r/w/o: 507.98/2092.91/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 16s ] thds: 16 tps: 23.00 qps: 2652.05 (r/w/o: 473.01/2179.04/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 17s ] thds: 16 tps: 30.00 qps: 2655.19 (r/w/o: 606.04/2049.15/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 18s ] thds: 16 tps: 27.00 qps: 2685.86 (r/w/o: 519.97/2165.89/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 19s ] thds: 16 tps: 22.00 qps: 2419.98 (r/w/o: 474.00/1945.99/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 20s ] thds: 16 tps: 30.00 qps: 2883.99 (r/w/o: 600.00/2283.99/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 21s ] thds: 16 tps: 22.00 qps: 2443.01 (r/w/o: 440.00/2003.01/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 22s ] thds: 16 tps: 28.00 qps: 2486.00 (r/w/o: 559.00/1927.00/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 23s ] thds: 16 tps: 23.00 qps: 2452.93 (r/w/o: 454.99/1997.94/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 24s ] thds: 16 tps: 25.00 qps: 2335.09 (r/w/o: 506.02/1829.07/0.00) lat (ms,99%): 773.68 err/s: 0.00 reconn/s: 0.00
[ 25s ] thds: 16 tps: 27.00 qps: 2832.00 (r/w/o: 540.00/2292.00/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 26s ] thds: 16 tps: 26.00 qps: 2729.94 (r/w/o: 512.99/2216.95/0.00) lat (ms,99%): 623.33 err/s: 0.00 reconn/s: 0.00
[ 27s ] thds: 16 tps: 27.00 qps: 2547.02 (r/w/o: 547.00/2000.02/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 28s ] thds: 16 tps: 26.00 qps: 2612.03 (r/w/o: 515.01/2097.03/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 29s ] thds: 16 tps: 25.00 qps: 2713.99 (r/w/o: 505.00/2209.00/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 30s ] thds: 16 tps: 29.00 qps: 2632.97 (r/w/o: 579.99/2052.97/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 31s ] thds: 16 tps: 23.00 qps: 2438.91 (r/w/o: 459.98/1978.93/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 32s ] thds: 16 tps: 25.00 qps: 2342.96 (r/w/o: 499.99/1842.97/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 33s ] thds: 16 tps: 24.00 qps: 2445.13 (r/w/o: 453.02/1992.11/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 34s ] thds: 16 tps: 24.00 qps: 2359.99 (r/w/o: 507.00/1853.00/0.00) lat (ms,99%): 773.68 err/s: 0.00 reconn/s: 0.00
[ 35s ] thds: 16 tps: 25.00 qps: 2508.02 (r/w/o: 479.00/2029.02/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 36s ] thds: 16 tps: 25.00 qps: 2713.01 (r/w/o: 521.00/2192.00/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 37s ] thds: 16 tps: 23.00 qps: 2133.00 (r/w/o: 432.00/1701.00/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
[ 38s ] thds: 16 tps: 24.00 qps: 2451.01 (r/w/o: 506.00/1945.01/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
[ 39s ] thds: 16 tps: 25.00 qps: 2483.00 (r/w/o: 502.00/1981.00/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 40s ] thds: 16 tps: 24.00 qps: 2366.97 (r/w/o: 479.99/1886.97/0.00) lat (ms,99%): 802.05 err/s: 0.00 reconn/s: 0.00
[ 41s ] thds: 16 tps: 21.00 qps: 2195.01 (r/w/o: 404.00/1791.01/0.00) lat (ms,99%): 759.88 err/s: 0.00 reconn/s: 0.00
[ 42s ] thds: 16 tps: 24.00 qps: 2285.01 (r/w/o: 496.00/1789.01/0.00) lat (ms,99%): 831.46 err/s: 0.00 reconn/s: 0.00
[ 43s ] thds: 16 tps: 23.00 qps: 2443.00 (r/w/o: 460.00/1983.00/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 44s ] thds: 16 tps: 26.00 qps: 2472.99 (r/w/o: 501.00/1971.99/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 45s ] thds: 16 tps: 23.00 qps: 2403.97 (r/w/o: 475.99/1927.98/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 46s ] thds: 16 tps: 28.00 qps: 2602.04 (r/w/o: 552.01/2050.03/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 47s ] thds: 16 tps: 22.00 qps: 2315.00 (r/w/o: 432.00/1883.00/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 48s ] thds: 16 tps: 23.00 qps: 2286.00 (r/w/o: 479.00/1807.00/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 49s ] thds: 16 tps: 21.00 qps: 2263.93 (r/w/o: 419.99/1843.94/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
[ 50s ] thds: 16 tps: 29.00 qps: 2625.07 (r/w/o: 561.01/2064.05/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 51s ] thds: 16 tps: 23.00 qps: 2312.01 (r/w/o: 443.00/1869.01/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 52s ] thds: 16 tps: 23.00 qps: 2425.99 (r/w/o: 488.00/1938.00/0.00) lat (ms,99%): 773.68 err/s: 0.00 reconn/s: 0.00
[ 53s ] thds: 16 tps: 25.00 qps: 2472.00 (r/w/o: 508.00/1964.00/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 54s ] thds: 16 tps: 25.00 qps: 2485.01 (r/w/o: 498.00/1987.00/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 55s ] thds: 16 tps: 26.00 qps: 2577.00 (r/w/o: 522.00/2055.00/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 56s ] thds: 16 tps: 23.00 qps: 2362.00 (r/w/o: 460.00/1902.00/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 57s ] thds: 16 tps: 25.00 qps: 2561.00 (r/w/o: 500.00/2061.00/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
[ 58s ] thds: 16 tps: 25.00 qps: 2447.97 (r/w/o: 493.99/1953.98/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 59s ] thds: 16 tps: 27.00 qps: 2670.02 (r/w/o: 527.00/2143.02/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 60s ] thds: 16 tps: 26.00 qps: 2637.00 (r/w/o: 539.00/2098.00/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 61s ] thds: 16 tps: 24.00 qps: 2326.00 (r/w/o: 443.00/1883.00/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 62s ] thds: 16 tps: 22.00 qps: 2307.01 (r/w/o: 476.00/1831.00/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 63s ] thds: 16 tps: 26.00 qps: 2545.82 (r/w/o: 516.96/2028.86/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
[ 64s ] thds: 16 tps: 21.00 qps: 2163.11 (r/w/o: 424.02/1739.09/0.00) lat (ms,99%): 773.68 err/s: 0.00 reconn/s: 0.00
[ 65s ] thds: 16 tps: 22.00 qps: 2244.01 (r/w/o: 440.00/1804.01/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
[ 66s ] thds: 16 tps: 24.00 qps: 2359.94 (r/w/o: 479.99/1879.95/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 67s ] thds: 16 tps: 23.00 qps: 2276.09 (r/w/o: 460.02/1816.07/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 68s ] thds: 16 tps: 25.00 qps: 2462.00 (r/w/o: 493.00/1969.00/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 69s ] thds: 16 tps: 22.00 qps: 2137.97 (r/w/o: 429.99/1707.98/0.00) lat (ms,99%): 802.05 err/s: 0.00 reconn/s: 0.00
[ 70s ] thds: 16 tps: 25.00 qps: 2591.04 (r/w/o: 498.01/2093.03/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 71s ] thds: 16 tps: 29.00 qps: 2835.98 (r/w/o: 599.00/2236.99/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 72s ] thds: 16 tps: 24.00 qps: 2346.99 (r/w/o: 433.00/1913.99/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 73s ] thds: 16 tps: 24.00 qps: 2365.02 (r/w/o: 484.00/1881.01/0.00) lat (ms,99%): 773.68 err/s: 0.00 reconn/s: 0.00
[ 74s ] thds: 16 tps: 26.00 qps: 2658.85 (r/w/o: 562.97/2095.88/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 75s ] thds: 16 tps: 23.00 qps: 2345.13 (r/w/o: 460.03/1885.11/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 76s ] thds: 16 tps: 25.00 qps: 2549.98 (r/w/o: 500.00/2049.98/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 77s ] thds: 16 tps: 25.00 qps: 2522.04 (r/w/o: 500.01/2022.03/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 78s ] thds: 16 tps: 24.00 qps: 2412.99 (r/w/o: 480.00/1932.99/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 79s ] thds: 16 tps: 24.00 qps: 2286.00 (r/w/o: 463.00/1823.00/0.00) lat (ms,99%): 773.68 err/s: 0.00 reconn/s: 0.00
[ 80s ] thds: 16 tps: 24.00 qps: 2560.00 (r/w/o: 497.00/2063.00/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 81s ] thds: 16 tps: 24.00 qps: 2299.99 (r/w/o: 476.00/1823.99/0.00) lat (ms,99%): 759.88 err/s: 0.00 reconn/s: 0.00
[ 82s ] thds: 16 tps: 28.00 qps: 2691.00 (r/w/o: 548.00/2143.00/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 83s ] thds: 16 tps: 23.00 qps: 2557.01 (r/w/o: 476.00/2081.01/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 84s ] thds: 16 tps: 28.00 qps: 2618.00 (r/w/o: 542.00/2076.00/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 85s ] thds: 16 tps: 25.00 qps: 2615.99 (r/w/o: 518.00/2098.00/0.00) lat (ms,99%): 612.21 err/s: 0.00 reconn/s: 0.00
[ 86s ] thds: 16 tps: 33.00 qps: 3266.97 (r/w/o: 648.99/2617.97/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 87s ] thds: 16 tps: 28.00 qps: 2594.87 (r/w/o: 515.97/2078.89/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 88s ] thds: 16 tps: 20.00 qps: 2274.14 (r/w/o: 455.03/1819.11/0.00) lat (ms,99%): 759.88 err/s: 0.00 reconn/s: 0.00
[ 89s ] thds: 16 tps: 31.00 qps: 2869.00 (r/w/o: 598.00/2271.00/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 90s ] thds: 16 tps: 24.00 qps: 2500.99 (r/w/o: 502.00/1998.99/0.00) lat (ms,99%): 759.88 err/s: 0.00 reconn/s: 0.00
[ 91s ] thds: 16 tps: 26.00 qps: 2590.99 (r/w/o: 520.00/2070.99/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 92s ] thds: 16 tps: 23.00 qps: 2256.01 (r/w/o: 459.00/1797.01/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 93s ] thds: 16 tps: 27.00 qps: 2788.98 (r/w/o: 541.00/2247.98/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 94s ] thds: 16 tps: 23.00 qps: 2375.98 (r/w/o: 460.00/1915.98/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 95s ] thds: 16 tps: 25.00 qps: 2422.04 (r/w/o: 500.01/1922.03/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 96s ] thds: 16 tps: 21.00 qps: 2178.95 (r/w/o: 419.99/1758.96/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 97s ] thds: 16 tps: 25.00 qps: 2276.06 (r/w/o: 482.01/1794.05/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
[ 98s ] thds: 16 tps: 23.00 qps: 2486.00 (r/w/o: 478.00/2008.00/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 99s ] thds: 16 tps: 29.00 qps: 2786.95 (r/w/o: 570.99/2215.96/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 100s ] thds: 16 tps: 20.00 qps: 2276.04 (r/w/o: 409.01/1867.03/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 101s ] thds: 16 tps: 28.00 qps: 2519.97 (r/w/o: 547.99/1971.97/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 102s ] thds: 16 tps: 27.00 qps: 2604.04 (r/w/o: 543.01/2061.03/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 103s ] thds: 16 tps: 25.00 qps: 2649.96 (r/w/o: 495.99/2153.97/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 104s ] thds: 16 tps: 24.00 qps: 2343.04 (r/w/o: 493.01/1850.03/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 105s ] thds: 16 tps: 21.00 qps: 2121.89 (r/w/o: 419.98/1701.92/0.00) lat (ms,99%): 773.68 err/s: 0.00 reconn/s: 0.00
[ 106s ] thds: 16 tps: 27.00 qps: 2633.99 (r/w/o: 540.00/2093.99/0.00) lat (ms,99%): 802.05 err/s: 0.00 reconn/s: 0.00
[ 107s ] thds: 16 tps: 24.00 qps: 2544.12 (r/w/o: 480.02/2064.10/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 108s ] thds: 16 tps: 27.00 qps: 2556.94 (r/w/o: 524.99/2031.95/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 109s ] thds: 16 tps: 22.00 qps: 2573.06 (r/w/o: 455.01/2118.05/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 110s ] thds: 16 tps: 28.00 qps: 2426.03 (r/w/o: 559.01/1867.02/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
[ 111s ] thds: 16 tps: 20.00 qps: 2258.97 (r/w/o: 393.00/1865.98/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 112s ] thds: 16 tps: 31.00 qps: 2888.01 (r/w/o: 612.00/2276.01/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 113s ] thds: 16 tps: 29.00 qps: 2932.01 (r/w/o: 596.00/2336.01/0.00) lat (ms,99%): 580.02 err/s: 0.00 reconn/s: 0.00
[ 114s ] thds: 16 tps: 24.00 qps: 2456.98 (r/w/o: 480.00/1976.98/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 115s ] thds: 16 tps: 23.00 qps: 2248.99 (r/w/o: 446.00/1802.99/0.00) lat (ms,99%): 773.68 err/s: 0.00 reconn/s: 0.00
[ 116s ] thds: 16 tps: 29.00 qps: 2749.04 (r/w/o: 578.01/2171.03/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 117s ] thds: 16 tps: 23.00 qps: 2457.00 (r/w/o: 472.00/1985.00/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 118s ] thds: 16 tps: 26.00 qps: 2604.83 (r/w/o: 523.97/2080.87/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 119s ] thds: 16 tps: 27.00 qps: 2714.15 (r/w/o: 516.03/2198.13/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 120s ] thds: 16 tps: 23.00 qps: 2240.01 (r/w/o: 455.00/1785.01/0.00) lat (ms,99%): 787.74 err/s: 0.00 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            60220
        write:                           240880
        other:                           0
        total:                           301100
    transactions:                        3011   (24.97 per sec.)
    queries:                             301100 (2497.20 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          120.5742s
    total number of events:              3011

Latency (ms):
         min:                                  427.98
         avg:                                  639.27
         max:                                  829.88
         99th percentile:                      773.68
         sum:                              1924843.00

Threads fairness:
    events (avg/stddev):           188.1875/0.53
    execution time (avg/stddev):   120.3027/0.18

cd ./sysbench/src/lua;sysbench --db-driver=mysql --mysql-host=172.16.251.240 --mysql-port=15306 --mysql-user=root --mysql-password=123456 --mysql-db=mydb --table_size=2000000 --tables=50 --time=120 --threads=32 --events=0 --percentile=99 --read-percent=20 --write-percent=80 --skip_trx=on --mysql-ignore-errors=1062 --db-ps-mode=disable --report-interval=1 oltp_read_write_pct run
sysbench 1.0.20 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 32
Report intermediate results every 1 second(s)
Initializing random number generator from current time


Initializing worker threads...

Threads started!

[ 1s ] thds: 32 tps: 31.96 qps: 5130.91 (r/w/o: 1278.23/3852.68/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 2s ] thds: 32 tps: 64.03 qps: 5409.88 (r/w/o: 1280.68/4129.20/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 3s ] thds: 32 tps: 32.00 qps: 4622.00 (r/w/o: 640.00/3982.00/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 4s ] thds: 32 tps: 64.00 qps: 5193.95 (r/w/o: 1279.99/3913.96/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 5s ] thds: 32 tps: 32.00 qps: 4767.91 (r/w/o: 639.99/4127.92/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 6s ] thds: 32 tps: 64.00 qps: 5299.08 (r/w/o: 1280.02/4019.06/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 7s ] thds: 32 tps: 56.00 qps: 5275.96 (r/w/o: 1041.99/4233.97/0.00) lat (ms,99%): 634.66 err/s: 0.00 reconn/s: 0.00
[ 8s ] thds: 32 tps: 40.00 qps: 4696.95 (r/w/o: 877.99/3818.96/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 9s ] thds: 32 tps: 56.00 qps: 4884.05 (r/w/o: 1036.01/3848.04/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 10s ] thds: 32 tps: 40.00 qps: 4882.11 (r/w/o: 884.02/3998.09/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 11s ] thds: 32 tps: 60.00 qps: 5022.95 (r/w/o: 1196.99/3825.96/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 12s ] thds: 32 tps: 38.00 qps: 5062.02 (r/w/o: 746.00/4316.02/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 13s ] thds: 32 tps: 62.00 qps: 4913.04 (r/w/o: 1237.01/3676.03/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 14s ] thds: 32 tps: 36.00 qps: 4837.02 (r/w/o: 723.00/4114.02/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 15s ] thds: 32 tps: 60.00 qps: 4959.67 (r/w/o: 1216.92/3742.75/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 16s ] thds: 32 tps: 42.00 qps: 4969.33 (r/w/o: 767.05/4202.28/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 17s ] thds: 32 tps: 54.00 qps: 4988.99 (r/w/o: 1153.00/3835.99/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 18s ] thds: 32 tps: 56.00 qps: 5268.92 (r/w/o: 1061.98/4206.94/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 19s ] thds: 32 tps: 42.00 qps: 4939.00 (r/w/o: 869.00/4070.00/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 20s ] thds: 32 tps: 58.00 qps: 5032.07 (r/w/o: 1166.02/3866.06/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 21s ] thds: 32 tps: 40.00 qps: 5040.00 (r/w/o: 820.00/4220.00/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 22s ] thds: 32 tps: 60.00 qps: 4958.95 (r/w/o: 1197.99/3760.96/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 23s ] thds: 32 tps: 36.00 qps: 4574.92 (r/w/o: 717.99/3856.93/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 24s ] thds: 32 tps: 60.00 qps: 5138.03 (r/w/o: 1200.01/3938.02/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 25s ] thds: 32 tps: 39.00 qps: 4847.94 (r/w/o: 751.99/4095.95/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 26s ] thds: 32 tps: 57.00 qps: 4851.94 (r/w/o: 1174.98/3676.95/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 27s ] thds: 32 tps: 40.00 qps: 4683.20 (r/w/o: 755.03/3928.17/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 28s ] thds: 32 tps: 57.00 qps: 5246.94 (r/w/o: 1168.99/4077.96/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 29s ] thds: 32 tps: 52.00 qps: 5013.97 (r/w/o: 1016.99/3996.97/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 30s ] thds: 32 tps: 45.00 qps: 4897.11 (r/w/o: 939.02/3958.09/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 31s ] thds: 32 tps: 57.00 qps: 5161.92 (r/w/o: 1138.98/4022.94/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 32s ] thds: 32 tps: 42.00 qps: 4758.07 (r/w/o: 809.01/3949.06/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 33s ] thds: 32 tps: 56.00 qps: 5204.01 (r/w/o: 1139.00/4065.01/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 34s ] thds: 32 tps: 47.00 qps: 5129.97 (r/w/o: 921.99/4207.98/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 35s ] thds: 32 tps: 52.00 qps: 4813.99 (r/w/o: 1071.00/3742.99/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 36s ] thds: 32 tps: 45.00 qps: 4814.00 (r/w/o: 850.00/3964.00/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 37s ] thds: 32 tps: 52.00 qps: 5068.01 (r/w/o: 1084.00/3984.01/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 38s ] thds: 32 tps: 52.00 qps: 4825.91 (r/w/o: 991.98/3833.93/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 39s ] thds: 32 tps: 45.00 qps: 4985.11 (r/w/o: 954.02/4031.09/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 40s ] thds: 32 tps: 56.00 qps: 5065.93 (r/w/o: 1083.99/3981.95/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 41s ] thds: 32 tps: 45.00 qps: 5189.08 (r/w/o: 912.01/4277.06/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 42s ] thds: 32 tps: 54.00 qps: 4992.99 (r/w/o: 1086.00/3906.99/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 43s ] thds: 32 tps: 48.00 qps: 5124.92 (r/w/o: 957.99/4166.94/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 44s ] thds: 32 tps: 52.00 qps: 5145.03 (r/w/o: 1060.01/4085.03/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 45s ] thds: 32 tps: 53.00 qps: 4824.01 (r/w/o: 1032.00/3792.00/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 46s ] thds: 32 tps: 45.00 qps: 5090.84 (r/w/o: 908.97/4181.87/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 47s ] thds: 32 tps: 56.00 qps: 5097.19 (r/w/o: 1125.04/3972.15/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 48s ] thds: 32 tps: 46.00 qps: 5083.97 (r/w/o: 933.99/4149.98/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 49s ] thds: 32 tps: 53.00 qps: 5097.55 (r/w/o: 1059.91/4037.64/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 50s ] thds: 32 tps: 55.01 qps: 5231.50 (r/w/o: 1085.10/4146.39/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 51s ] thds: 32 tps: 45.00 qps: 5094.93 (r/w/o: 914.99/4179.94/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 52s ] thds: 32 tps: 55.00 qps: 4754.66 (r/w/o: 1084.92/3669.74/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 53s ] thds: 32 tps: 42.00 qps: 4876.28 (r/w/o: 840.05/4036.23/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 54s ] thds: 32 tps: 58.00 qps: 5236.17 (r/w/o: 1168.04/4068.13/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 55s ] thds: 32 tps: 41.00 qps: 4632.01 (r/w/o: 774.00/3858.01/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 56s ] thds: 32 tps: 55.00 qps: 5101.96 (r/w/o: 1152.99/3948.97/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 57s ] thds: 32 tps: 50.00 qps: 5110.02 (r/w/o: 986.00/4124.02/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 58s ] thds: 32 tps: 48.00 qps: 5015.73 (r/w/o: 972.95/4042.78/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 59s ] thds: 32 tps: 59.00 qps: 5418.07 (r/w/o: 1166.01/4252.05/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 60s ] thds: 32 tps: 44.00 qps: 4964.96 (r/w/o: 885.99/4078.97/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 61s ] thds: 32 tps: 56.00 qps: 5259.09 (r/w/o: 1103.02/4156.07/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 62s ] thds: 32 tps: 49.00 qps: 4938.78 (r/w/o: 1005.95/3932.82/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 63s ] thds: 32 tps: 50.00 qps: 5054.37 (r/w/o: 973.07/4081.30/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 64s ] thds: 32 tps: 49.00 qps: 4853.98 (r/w/o: 992.00/3861.99/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 65s ] thds: 32 tps: 48.00 qps: 4873.88 (r/w/o: 960.98/3912.91/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 66s ] thds: 32 tps: 51.00 qps: 4845.04 (r/w/o: 988.01/3857.03/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 67s ] thds: 32 tps: 47.00 qps: 5033.09 (r/w/o: 965.02/4068.07/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 68s ] thds: 32 tps: 52.00 qps: 4935.02 (r/w/o: 1060.01/3875.02/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 69s ] thds: 32 tps: 43.00 qps: 4584.75 (r/w/o: 860.95/3723.80/0.00) lat (ms,99%): 759.88 err/s: 0.00 reconn/s: 0.00
[ 70s ] thds: 32 tps: 52.00 qps: 4832.23 (r/w/o: 1014.05/3818.18/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 71s ] thds: 32 tps: 44.00 qps: 4834.03 (r/w/o: 906.00/3928.02/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 72s ] thds: 32 tps: 52.00 qps: 4890.03 (r/w/o: 1033.01/3857.02/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 73s ] thds: 32 tps: 46.00 qps: 4987.01 (r/w/o: 911.00/4076.00/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 74s ] thds: 32 tps: 57.00 qps: 5253.98 (r/w/o: 1108.00/4145.99/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 75s ] thds: 32 tps: 48.00 qps: 4842.97 (r/w/o: 953.99/3888.97/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 76s ] thds: 32 tps: 46.00 qps: 4736.92 (r/w/o: 967.98/3768.94/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 77s ] thds: 32 tps: 53.00 qps: 5124.91 (r/w/o: 1041.98/4082.93/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 78s ] thds: 32 tps: 49.00 qps: 5019.19 (r/w/o: 1001.04/4018.15/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 79s ] thds: 32 tps: 54.00 qps: 5174.00 (r/w/o: 1051.00/4123.00/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 80s ] thds: 32 tps: 43.00 qps: 4719.03 (r/w/o: 892.00/3827.02/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 81s ] thds: 32 tps: 54.00 qps: 4943.92 (r/w/o: 1039.98/3903.93/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 82s ] thds: 32 tps: 49.00 qps: 5253.08 (r/w/o: 986.01/4267.06/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 83s ] thds: 32 tps: 50.00 qps: 4920.92 (r/w/o: 1003.98/3916.93/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 84s ] thds: 32 tps: 55.00 qps: 5249.09 (r/w/o: 1125.02/4124.07/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 85s ] thds: 32 tps: 45.00 qps: 4881.81 (r/w/o: 904.97/3976.85/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 86s ] thds: 32 tps: 54.00 qps: 4951.04 (r/w/o: 1047.01/3904.03/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 87s ] thds: 32 tps: 45.00 qps: 5120.87 (r/w/o: 932.98/4187.89/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 88s ] thds: 32 tps: 52.00 qps: 4670.19 (r/w/o: 1006.04/3664.15/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 89s ] thds: 32 tps: 45.00 qps: 4864.07 (r/w/o: 909.01/3955.06/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 90s ] thds: 32 tps: 53.00 qps: 4882.79 (r/w/o: 1060.95/3821.84/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 91s ] thds: 32 tps: 44.00 qps: 5102.20 (r/w/o: 902.03/4200.16/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 92s ] thds: 32 tps: 54.00 qps: 4937.01 (r/w/o: 1082.00/3855.00/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 93s ] thds: 32 tps: 52.00 qps: 5124.98 (r/w/o: 986.00/4138.98/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 94s ] thds: 32 tps: 50.00 qps: 5119.89 (r/w/o: 1040.98/4078.91/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 95s ] thds: 32 tps: 52.00 qps: 4936.14 (r/w/o: 1045.03/3891.11/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 96s ] thds: 32 tps: 46.00 qps: 4974.98 (r/w/o: 928.00/4046.98/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 97s ] thds: 32 tps: 56.00 qps: 5144.02 (r/w/o: 1095.00/4049.02/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 98s ] thds: 32 tps: 42.00 qps: 4913.02 (r/w/o: 859.00/4054.02/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 99s ] thds: 32 tps: 54.00 qps: 5045.93 (r/w/o: 1085.98/3959.94/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 100s ] thds: 32 tps: 45.00 qps: 4747.04 (r/w/o: 876.01/3871.03/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 101s ] thds: 32 tps: 54.00 qps: 5080.00 (r/w/o: 1076.00/4004.00/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 102s ] thds: 32 tps: 51.00 qps: 5101.02 (r/w/o: 1012.00/4089.01/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 103s ] thds: 32 tps: 49.00 qps: 4952.01 (r/w/o: 1005.00/3947.01/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 104s ] thds: 32 tps: 52.00 qps: 5010.01 (r/w/o: 1040.00/3970.01/0.00) lat (ms,99%): 657.93 err/s: 0.00 reconn/s: 0.00
[ 105s ] thds: 32 tps: 46.00 qps: 5013.99 (r/w/o: 912.00/4101.99/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 106s ] thds: 32 tps: 55.00 qps: 4963.99 (r/w/o: 1076.00/3887.99/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
[ 107s ] thds: 32 tps: 45.00 qps: 4855.93 (r/w/o: 939.99/3915.95/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 108s ] thds: 32 tps: 49.00 qps: 4681.07 (r/w/o: 932.01/3749.05/0.00) lat (ms,99%): 746.32 err/s: 0.00 reconn/s: 0.00
[ 109s ] thds: 32 tps: 47.00 qps: 4922.70 (r/w/o: 986.94/3935.76/0.00) lat (ms,99%): 719.92 err/s: 0.00 reconn/s: 0.00
[ 110s ] thds: 32 tps: 54.00 qps: 5045.26 (r/w/o: 1034.05/4011.21/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 111s ] thds: 32 tps: 42.00 qps: 4690.04 (r/w/o: 885.01/3805.03/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 112s ] thds: 32 tps: 55.00 qps: 4997.90 (r/w/o: 1063.98/3933.92/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 113s ] thds: 32 tps: 47.00 qps: 4973.12 (r/w/o: 922.02/4051.10/0.00) lat (ms,99%): 694.45 err/s: 0.00 reconn/s: 0.00
[ 114s ] thds: 32 tps: 50.00 qps: 5003.70 (r/w/o: 1058.94/3944.76/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 115s ] thds: 32 tps: 43.00 qps: 4516.25 (r/w/o: 840.05/3676.21/0.00) lat (ms,99%): 759.88 err/s: 0.00 reconn/s: 0.00
[ 116s ] thds: 32 tps: 54.00 qps: 5210.97 (r/w/o: 1085.99/4124.98/0.00) lat (ms,99%): 682.06 err/s: 0.00 reconn/s: 0.00
[ 117s ] thds: 32 tps: 47.00 qps: 4765.75 (r/w/o: 953.95/3811.80/0.00) lat (ms,99%): 733.00 err/s: 0.00 reconn/s: 0.00
[ 118s ] thds: 32 tps: 50.00 qps: 5105.31 (r/w/o: 990.06/4115.25/0.00) lat (ms,99%): 707.07 err/s: 0.00 reconn/s: 0.00
[ 119s ] thds: 32 tps: 55.00 qps: 5262.84 (r/w/o: 1102.97/4159.87/0.00) lat (ms,99%): 646.19 err/s: 0.00 reconn/s: 0.00
[ 120s ] thds: 32 tps: 49.00 qps: 5020.93 (r/w/o: 972.99/4047.94/0.00) lat (ms,99%): 669.89 err/s: 0.00 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            119880
        write:                           479520
        other:                           0
        total:                           599400
    transactions:                        5994   (49.72 per sec.)
    queries:                             599400 (4972.14 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          120.5509s
    total number of events:              5994

Latency (ms):
         min:                                  555.08
         avg:                                  642.14
         max:                                  758.47
         99th percentile:                      719.92
         sum:                              3848957.43

Threads fairness:
    events (avg/stddev):           187.3125/0.46
    execution time (avg/stddev):   120.2799/0.17

cd ./sysbench/src/lua;sysbench --db-driver=mysql --mysql-host=172.16.251.240 --mysql-port=15306 --mysql-user=root --mysql-password=123456 --mysql-db=mydb --table_size=2000000 --tables=50 --time=120 --threads=64 --events=0 --percentile=99 --read-percent=20 --write-percent=80 --skip_trx=on --mysql-ignore-errors=1062 --db-ps-mode=disable --report-interval=1 oltp_read_write_pct run
sysbench 1.0.20 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 64
Report intermediate results every 1 second(s)
Initializing random number generator from current time


Initializing worker threads...

Threads started!

[ 1s ] thds: 64 tps: 2.99 qps: 6127.32 (r/w/o: 1291.90/4835.41/0.00) lat (ms,99%): 995.51 err/s: 0.00 reconn/s: 0.00
[ 2s ] thds: 64 tps: 74.05 qps: 6570.35 (r/w/o: 1400.93/5169.42/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 3s ] thds: 64 tps: 58.00 qps: 6225.73 (r/w/o: 1249.95/4975.78/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 4s ] thds: 64 tps: 90.00 qps: 7084.04 (r/w/o: 1670.01/5414.03/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 5s ] thds: 64 tps: 65.00 qps: 6383.26 (r/w/o: 1276.05/5107.21/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 6s ] thds: 64 tps: 73.00 qps: 6647.92 (r/w/o: 1423.98/5223.94/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 7s ] thds: 64 tps: 69.00 qps: 6621.77 (r/w/o: 1394.95/5226.82/0.00) lat (ms,99%): 1032.01 err/s: 0.00 reconn/s: 0.00
[ 8s ] thds: 64 tps: 73.00 qps: 6830.67 (r/w/o: 1505.93/5324.74/0.00) lat (ms,99%): 1013.60 err/s: 0.00 reconn/s: 0.00
[ 9s ] thds: 64 tps: 63.01 qps: 6389.60 (r/w/o: 1259.12/5130.49/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 10s ] thds: 64 tps: 38.00 qps: 5510.75 (r/w/o: 814.96/4695.79/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 11s ] thds: 64 tps: 73.00 qps: 6734.30 (r/w/o: 1445.06/5289.23/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 12s ] thds: 64 tps: 48.00 qps: 5831.89 (r/w/o: 1024.98/4806.91/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 13s ] thds: 64 tps: 51.00 qps: 5798.99 (r/w/o: 1045.00/4753.99/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 14s ] thds: 64 tps: 58.00 qps: 5880.74 (r/w/o: 1163.95/4716.79/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 15s ] thds: 64 tps: 62.00 qps: 6303.37 (r/w/o: 1256.07/5047.30/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 16s ] thds: 64 tps: 64.00 qps: 6309.82 (r/w/o: 1282.96/5026.85/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 17s ] thds: 64 tps: 63.00 qps: 6184.22 (r/w/o: 1272.04/4912.17/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 18s ] thds: 64 tps: 63.00 qps: 5525.71 (r/w/o: 1259.93/4265.78/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 19s ] thds: 64 tps: 64.00 qps: 5740.08 (r/w/o: 1280.02/4460.06/0.00) lat (ms,99%): 1235.62 err/s: 0.00 reconn/s: 0.00
[ 20s ] thds: 64 tps: 64.00 qps: 6041.20 (r/w/o: 1280.04/4761.16/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 21s ] thds: 64 tps: 64.00 qps: 5989.02 (r/w/o: 1280.00/4709.01/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 22s ] thds: 64 tps: 64.00 qps: 5939.02 (r/w/o: 1272.00/4667.01/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 23s ] thds: 64 tps: 63.00 qps: 6199.61 (r/w/o: 1249.92/4949.69/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 24s ] thds: 64 tps: 65.00 qps: 6727.90 (r/w/o: 1317.98/5409.92/0.00) lat (ms,99%): 1050.76 err/s: 0.00 reconn/s: 0.00
[ 25s ] thds: 64 tps: 63.00 qps: 6040.42 (r/w/o: 1201.08/4839.34/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 26s ] thds: 64 tps: 58.00 qps: 5933.78 (r/w/o: 1116.96/4816.82/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 27s ] thds: 64 tps: 41.00 qps: 5502.01 (r/w/o: 889.00/4613.01/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 28s ] thds: 64 tps: 61.00 qps: 6133.06 (r/w/o: 1177.01/4956.05/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 29s ] thds: 64 tps: 62.00 qps: 6353.88 (r/w/o: 1247.98/5105.90/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 30s ] thds: 64 tps: 45.00 qps: 5423.24 (r/w/o: 957.04/4466.20/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 31s ] thds: 64 tps: 59.00 qps: 5686.89 (r/w/o: 1155.98/4530.91/0.00) lat (ms,99%): 1213.57 err/s: 0.00 reconn/s: 0.00
[ 32s ] thds: 64 tps: 62.00 qps: 6169.08 (r/w/o: 1268.02/4901.06/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 33s ] thds: 64 tps: 63.00 qps: 6209.98 (r/w/o: 1255.00/4954.98/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 34s ] thds: 64 tps: 62.00 qps: 6031.95 (r/w/o: 1251.99/4779.96/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 35s ] thds: 64 tps: 63.99 qps: 5847.51 (r/w/o: 1279.89/4567.62/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 36s ] thds: 64 tps: 64.01 qps: 6502.71 (r/w/o: 1280.14/5222.57/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 37s ] thds: 64 tps: 63.00 qps: 5562.95 (r/w/o: 1259.99/4302.96/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 38s ] thds: 64 tps: 64.00 qps: 6140.06 (r/w/o: 1270.01/4870.05/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 39s ] thds: 64 tps: 60.99 qps: 5919.46 (r/w/o: 1206.89/4712.57/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 40s ] thds: 64 tps: 64.00 qps: 6287.43 (r/w/o: 1265.09/5022.34/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 41s ] thds: 64 tps: 36.00 qps: 4777.05 (r/w/o: 719.01/4058.04/0.00) lat (ms,99%): 1280.93 err/s: 0.00 reconn/s: 0.00
[ 42s ] thds: 64 tps: 52.00 qps: 5626.65 (r/w/o: 981.94/4644.71/0.00) lat (ms,99%): 1304.21 err/s: 0.00 reconn/s: 0.00
[ 43s ] thds: 64 tps: 65.00 qps: 6485.25 (r/w/o: 1341.05/5144.19/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 44s ] thds: 64 tps: 59.00 qps: 6055.23 (r/w/o: 1195.05/4860.19/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 45s ] thds: 64 tps: 57.00 qps: 5770.00 (r/w/o: 1117.00/4653.00/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 46s ] thds: 64 tps: 60.00 qps: 6146.99 (r/w/o: 1249.00/4897.99/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 47s ] thds: 64 tps: 62.00 qps: 5958.94 (r/w/o: 1254.99/4703.96/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 48s ] thds: 64 tps: 63.00 qps: 5824.04 (r/w/o: 1250.01/4574.03/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 49s ] thds: 64 tps: 62.00 qps: 6134.91 (r/w/o: 1229.98/4904.93/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 50s ] thds: 64 tps: 63.00 qps: 5687.94 (r/w/o: 1279.99/4407.96/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 51s ] thds: 64 tps: 63.00 qps: 6006.08 (r/w/o: 1206.02/4800.06/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 52s ] thds: 64 tps: 54.00 qps: 5708.99 (r/w/o: 1100.00/4608.99/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 53s ] thds: 64 tps: 55.00 qps: 5797.05 (r/w/o: 1123.01/4674.04/0.00) lat (ms,99%): 1213.57 err/s: 0.00 reconn/s: 0.00
[ 54s ] thds: 64 tps: 54.00 qps: 5671.02 (r/w/o: 1020.00/4651.02/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 55s ] thds: 64 tps: 64.00 qps: 6396.86 (r/w/o: 1276.97/5119.89/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 56s ] thds: 64 tps: 56.00 qps: 5929.06 (r/w/o: 1170.01/4759.05/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 57s ] thds: 64 tps: 56.00 qps: 5741.07 (r/w/o: 1097.01/4644.05/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 58s ] thds: 64 tps: 61.00 qps: 6095.01 (r/w/o: 1230.00/4865.00/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 59s ] thds: 64 tps: 59.00 qps: 5935.88 (r/w/o: 1187.98/4747.90/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 60s ] thds: 64 tps: 60.00 qps: 5837.13 (r/w/o: 1218.03/4619.10/0.00) lat (ms,99%): 1235.62 err/s: 0.00 reconn/s: 0.00
[ 61s ] thds: 64 tps: 61.00 qps: 5712.89 (r/w/o: 1230.98/4481.91/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 62s ] thds: 64 tps: 64.00 qps: 6314.09 (r/w/o: 1246.02/5068.07/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 63s ] thds: 64 tps: 63.00 qps: 6373.98 (r/w/o: 1285.00/5088.99/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 64s ] thds: 64 tps: 64.00 qps: 6272.06 (r/w/o: 1277.01/4995.05/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 65s ] thds: 64 tps: 62.00 qps: 6027.98 (r/w/o: 1242.00/4785.98/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 66s ] thds: 64 tps: 64.00 qps: 6141.00 (r/w/o: 1275.00/4866.00/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 67s ] thds: 64 tps: 60.00 qps: 5852.01 (r/w/o: 1165.00/4687.01/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 68s ] thds: 64 tps: 61.00 qps: 6169.99 (r/w/o: 1256.00/4913.99/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 69s ] thds: 64 tps: 53.00 qps: 5490.66 (r/w/o: 934.94/4555.72/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 70s ] thds: 64 tps: 55.00 qps: 6054.20 (r/w/o: 1202.04/4852.16/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 71s ] thds: 64 tps: 57.00 qps: 5931.06 (r/w/o: 1169.01/4762.04/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 72s ] thds: 64 tps: 48.00 qps: 4819.10 (r/w/o: 921.02/3898.08/0.00) lat (ms,99%): 1327.91 err/s: 0.00 reconn/s: 0.00
[ 73s ] thds: 64 tps: 60.00 qps: 5954.98 (r/w/o: 1215.00/4739.98/0.00) lat (ms,99%): 1327.91 err/s: 0.00 reconn/s: 0.00
[ 74s ] thds: 64 tps: 58.00 qps: 5409.92 (r/w/o: 1158.98/4250.94/0.00) lat (ms,99%): 1235.62 err/s: 0.00 reconn/s: 0.00
[ 75s ] thds: 64 tps: 61.00 qps: 5912.08 (r/w/o: 1253.02/4659.06/0.00) lat (ms,99%): 1258.08 err/s: 0.00 reconn/s: 0.00
[ 76s ] thds: 64 tps: 64.00 qps: 6177.71 (r/w/o: 1265.94/4911.77/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 77s ] thds: 64 tps: 63.00 qps: 6128.27 (r/w/o: 1244.05/4884.22/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 78s ] thds: 64 tps: 59.00 qps: 5949.89 (r/w/o: 1152.98/4796.91/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 79s ] thds: 64 tps: 54.00 qps: 5825.12 (r/w/o: 1108.02/4717.10/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 80s ] thds: 64 tps: 59.00 qps: 5997.87 (r/w/o: 1129.97/4867.89/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 81s ] thds: 64 tps: 53.00 qps: 5433.01 (r/w/o: 1046.00/4387.01/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 82s ] thds: 64 tps: 53.00 qps: 5915.89 (r/w/o: 1144.98/4770.92/0.00) lat (ms,99%): 1213.57 err/s: 0.00 reconn/s: 0.00
[ 83s ] thds: 64 tps: 61.00 qps: 6114.29 (r/w/o: 1195.06/4919.24/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 84s ] thds: 64 tps: 63.00 qps: 6195.93 (r/w/o: 1267.99/4927.94/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 85s ] thds: 64 tps: 63.00 qps: 6095.81 (r/w/o: 1259.96/4835.85/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 86s ] thds: 64 tps: 59.00 qps: 5876.23 (r/w/o: 1205.05/4671.18/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 87s ] thds: 64 tps: 64.00 qps: 6138.95 (r/w/o: 1279.99/4858.96/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 88s ] thds: 64 tps: 64.00 qps: 6440.88 (r/w/o: 1273.98/5166.91/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 89s ] thds: 64 tps: 60.00 qps: 5523.10 (r/w/o: 1161.02/4362.08/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 90s ] thds: 64 tps: 57.99 qps: 5782.49 (r/w/o: 1180.90/4601.60/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 91s ] thds: 64 tps: 56.00 qps: 5567.48 (r/w/o: 1125.10/4442.38/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 92s ] thds: 64 tps: 62.00 qps: 6144.81 (r/w/o: 1197.96/4946.84/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 93s ] thds: 64 tps: 65.00 qps: 6427.22 (r/w/o: 1271.04/5156.18/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 94s ] thds: 64 tps: 65.00 qps: 6477.85 (r/w/o: 1337.97/5139.88/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 95s ] thds: 64 tps: 49.00 qps: 5352.02 (r/w/o: 969.00/4383.02/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 96s ] thds: 64 tps: 59.00 qps: 6067.82 (r/w/o: 1170.97/4896.85/0.00) lat (ms,99%): 1191.92 err/s: 0.00 reconn/s: 0.00
[ 97s ] thds: 64 tps: 44.00 qps: 5023.06 (r/w/o: 922.01/4101.05/0.00) lat (ms,99%): 1280.93 err/s: 0.00 reconn/s: 0.00
[ 98s ] thds: 64 tps: 61.00 qps: 6005.91 (r/w/o: 1233.98/4771.93/0.00) lat (ms,99%): 1280.93 err/s: 0.00 reconn/s: 0.00
[ 99s ] thds: 64 tps: 63.00 qps: 6002.27 (r/w/o: 1270.06/4732.21/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 100s ] thds: 64 tps: 62.99 qps: 6108.44 (r/w/o: 1265.88/4842.56/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 101s ] thds: 64 tps: 64.01 qps: 6157.54 (r/w/o: 1280.11/4877.43/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 102s ] thds: 64 tps: 62.00 qps: 5923.97 (r/w/o: 1227.99/4695.98/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 103s ] thds: 64 tps: 48.00 qps: 4709.13 (r/w/o: 911.02/3798.10/0.00) lat (ms,99%): 1401.61 err/s: 0.00 reconn/s: 0.00
[ 104s ] thds: 64 tps: 59.00 qps: 6173.51 (r/w/o: 1211.90/4961.61/0.00) lat (ms,99%): 1327.91 err/s: 0.00 reconn/s: 0.00
[ 105s ] thds: 64 tps: 50.00 qps: 5561.05 (r/w/o: 962.01/4599.04/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 106s ] thds: 64 tps: 60.00 qps: 6146.21 (r/w/o: 1227.04/4919.17/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 107s ] thds: 64 tps: 60.00 qps: 5794.13 (r/w/o: 1197.03/4597.10/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 108s ] thds: 64 tps: 57.00 qps: 6015.06 (r/w/o: 1156.01/4859.05/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 109s ] thds: 64 tps: 61.99 qps: 6196.45 (r/w/o: 1256.89/4939.56/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 110s ] thds: 64 tps: 59.01 qps: 5798.69 (r/w/o: 1188.14/4610.55/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 111s ] thds: 64 tps: 60.00 qps: 5887.65 (r/w/o: 1184.93/4702.72/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 112s ] thds: 64 tps: 63.00 qps: 6019.13 (r/w/o: 1272.03/4747.10/0.00) lat (ms,99%): 1109.09 err/s: 0.00 reconn/s: 0.00
[ 113s ] thds: 64 tps: 62.00 qps: 5800.88 (r/w/o: 1231.98/4568.91/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 114s ] thds: 64 tps: 60.00 qps: 5872.12 (r/w/o: 1196.03/4676.10/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 115s ] thds: 64 tps: 63.00 qps: 6210.79 (r/w/o: 1248.96/4961.83/0.00) lat (ms,99%): 1129.24 err/s: 0.00 reconn/s: 0.00
[ 116s ] thds: 64 tps: 64.00 qps: 6539.21 (r/w/o: 1295.04/5244.17/0.00) lat (ms,99%): 1069.86 err/s: 0.00 reconn/s: 0.00
[ 117s ] thds: 64 tps: 61.00 qps: 6017.03 (r/w/o: 1232.01/4785.03/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
[ 118s ] thds: 64 tps: 56.00 qps: 5811.94 (r/w/o: 1063.99/4747.96/0.00) lat (ms,99%): 1149.76 err/s: 0.00 reconn/s: 0.00
[ 119s ] thds: 64 tps: 60.00 qps: 6162.90 (r/w/o: 1218.98/4943.92/0.00) lat (ms,99%): 1170.65 err/s: 0.00 reconn/s: 0.00
[ 120s ] thds: 64 tps: 54.00 qps: 5828.14 (r/w/o: 1088.03/4740.11/0.00) lat (ms,99%): 1089.30 err/s: 0.00 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            144240
        write:                           576960
        other:                           0
        total:                           721200
    transactions:                        7212   (59.78 per sec.)
    queries:                             721200 (5977.86 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          120.6445s
    total number of events:              7212

Latency (ms):
         min:                                  655.83
         avg:                                 1067.86
         max:                                 1412.34
         99th percentile:                     1280.93
         sum:                              7701441.92

Threads fairness:
    events (avg/stddev):           112.6875/0.53
    execution time (avg/stddev):   120.3350/0.20

cd ./sysbench/src/lua;sysbench --db-driver=mysql --mysql-host=172.16.251.240 --mysql-port=15306 --mysql-user=root --mysql-password=123456 --mysql-db=mydb --table_size=2000000 --tables=50 --time=120 --threads=128 --events=0 --percentile=99 --read-percent=20 --write-percent=80 --skip_trx=on --mysql-ignore-errors=1062 --db-ps-mode=disable --report-interval=1 oltp_read_write_pct run
sysbench 1.0.20 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 128
Report intermediate results every 1 second(s)
Initializing random number generator from current time


Initializing worker threads...

Threads started!

[ 1s ] thds: 128 tps: 0.00 qps: 7731.19 (r/w/o: 2555.11/5176.09/0.00) lat (ms,99%): 0.00 err/s: 0.00 reconn/s: 0.00
[ 2s ] thds: 128 tps: 112.05 qps: 6664.93 (r/w/o: 1582.70/5082.23/0.00) lat (ms,99%): 1973.38 err/s: 0.00 reconn/s: 0.00
[ 3s ] thds: 128 tps: 16.00 qps: 6231.18 (r/w/o: 978.03/5253.15/0.00) lat (ms,99%): 2159.29 err/s: 0.00 reconn/s: 0.00
[ 4s ] thds: 128 tps: 43.00 qps: 5333.08 (r/w/o: 677.01/4656.07/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 5s ] thds: 128 tps: 85.00 qps: 7149.29 (r/w/o: 1883.08/5266.21/0.00) lat (ms,99%): 2449.36 err/s: 0.00 reconn/s: 0.00
[ 6s ] thds: 128 tps: 55.00 qps: 5994.02 (r/w/o: 957.00/5037.02/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 7s ] thds: 128 tps: 73.00 qps: 6556.86 (r/w/o: 1602.97/4953.89/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 8s ] thds: 128 tps: 42.00 qps: 5629.08 (r/w/o: 591.01/5038.07/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 9s ] thds: 128 tps: 86.00 qps: 6869.63 (r/w/o: 1968.89/4900.73/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 10s ] thds: 128 tps: 25.00 qps: 5533.98 (r/w/o: 405.00/5128.98/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 11s ] thds: 128 tps: 103.00 qps: 6926.88 (r/w/o: 2154.96/4771.92/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 12s ] thds: 128 tps: 23.00 qps: 5534.47 (r/w/o: 301.03/5233.44/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 13s ] thds: 128 tps: 105.00 qps: 6028.91 (r/w/o: 2258.97/3769.94/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 14s ] thds: 128 tps: 0.00 qps: 4555.69 (r/w/o: 0.00/4555.69/0.00) lat (ms,99%): 0.00 err/s: 0.00 reconn/s: 0.00
[ 15s ] thds: 128 tps: 126.01 qps: 7173.30 (r/w/o: 2520.10/4653.19/0.00) lat (ms,99%): 2632.28 err/s: 0.00 reconn/s: 0.00
[ 16s ] thds: 128 tps: 2.00 qps: 5582.93 (r/w/o: 40.00/5542.93/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 17s ] thds: 128 tps: 126.01 qps: 7165.32 (r/w/o: 2507.11/4658.21/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 18s ] thds: 128 tps: 4.00 qps: 5703.00 (r/w/o: 58.00/5645.00/0.00) lat (ms,99%): 2045.74 err/s: 0.00 reconn/s: 0.00
[ 19s ] thds: 128 tps: 122.99 qps: 7180.40 (r/w/o: 2494.79/4685.61/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 20s ] thds: 128 tps: 5.00 qps: 5915.42 (r/w/o: 89.01/5826.42/0.00) lat (ms,99%): 2045.74 err/s: 0.00 reconn/s: 0.00
[ 21s ] thds: 128 tps: 122.99 qps: 6592.68 (r/w/o: 2449.88/4142.80/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 22s ] thds: 128 tps: 4.00 qps: 5826.16 (r/w/o: 96.00/5730.15/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 23s ] thds: 128 tps: 124.00 qps: 7030.18 (r/w/o: 2448.06/4582.12/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 24s ] thds: 128 tps: 4.00 qps: 5649.09 (r/w/o: 117.00/5532.08/0.00) lat (ms,99%): 2082.91 err/s: 0.00 reconn/s: 0.00
[ 25s ] thds: 128 tps: 115.00 qps: 6544.72 (r/w/o: 2234.90/4309.81/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 26s ] thds: 128 tps: 15.00 qps: 6508.21 (r/w/o: 334.01/6174.20/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 27s ] thds: 128 tps: 113.00 qps: 6396.90 (r/w/o: 2196.97/4199.94/0.00) lat (ms,99%): 2159.29 err/s: 0.00 reconn/s: 0.00
[ 28s ] thds: 128 tps: 13.00 qps: 5770.49 (r/w/o: 353.97/5416.52/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 29s ] thds: 128 tps: 113.01 qps: 6803.79 (r/w/o: 2179.25/4624.54/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 30s ] thds: 128 tps: 18.00 qps: 6226.88 (r/w/o: 430.99/5795.88/0.00) lat (ms,99%): 2082.91 err/s: 0.00 reconn/s: 0.00
[ 31s ] thds: 128 tps: 114.00 qps: 6840.77 (r/w/o: 2224.92/4615.84/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 32s ] thds: 128 tps: 15.00 qps: 6208.99 (r/w/o: 349.00/5859.99/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 33s ] thds: 128 tps: 108.01 qps: 6386.34 (r/w/o: 2143.11/4243.23/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 34s ] thds: 128 tps: 19.00 qps: 6067.01 (r/w/o: 381.00/5686.01/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 35s ] thds: 128 tps: 102.00 qps: 6140.99 (r/w/o: 1860.00/4280.99/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 36s ] thds: 128 tps: 24.00 qps: 5916.86 (r/w/o: 669.98/5246.88/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 37s ] thds: 128 tps: 79.00 qps: 6012.09 (r/w/o: 1464.02/4548.07/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 38s ] thds: 128 tps: 52.00 qps: 7280.97 (r/w/o: 1161.00/6119.98/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 39s ] thds: 128 tps: 100.00 qps: 6327.09 (r/w/o: 1797.02/4530.06/0.00) lat (ms,99%): 2045.74 err/s: 0.00 reconn/s: 0.00
[ 40s ] thds: 128 tps: 25.00 qps: 6088.98 (r/w/o: 708.00/5380.98/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 41s ] thds: 128 tps: 101.00 qps: 6396.71 (r/w/o: 1638.93/4757.79/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 42s ] thds: 128 tps: 28.00 qps: 6360.13 (r/w/o: 950.02/5410.11/0.00) lat (ms,99%): 2362.72 err/s: 0.00 reconn/s: 0.00
[ 43s ] thds: 128 tps: 92.00 qps: 6488.13 (r/w/o: 1636.03/4852.10/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 44s ] thds: 128 tps: 33.00 qps: 5217.03 (r/w/o: 867.01/4350.03/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 45s ] thds: 128 tps: 60.00 qps: 6200.94 (r/w/o: 1067.99/5132.95/0.00) lat (ms,99%): 2362.72 err/s: 0.00 reconn/s: 0.00
[ 46s ] thds: 128 tps: 69.00 qps: 6974.92 (r/w/o: 1492.98/5481.94/0.00) lat (ms,99%): 2405.65 err/s: 0.00 reconn/s: 0.00
[ 47s ] thds: 128 tps: 62.00 qps: 5874.11 (r/w/o: 1128.02/4746.09/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 48s ] thds: 128 tps: 67.00 qps: 6800.71 (r/w/o: 1458.94/5341.77/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 49s ] thds: 128 tps: 51.00 qps: 5632.21 (r/w/o: 952.04/4680.17/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 50s ] thds: 128 tps: 76.00 qps: 6759.04 (r/w/o: 1584.01/5175.03/0.00) lat (ms,99%): 2159.29 err/s: 0.00 reconn/s: 0.00
[ 51s ] thds: 128 tps: 53.00 qps: 6208.65 (r/w/o: 1014.94/5193.71/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 52s ] thds: 128 tps: 75.00 qps: 6934.40 (r/w/o: 1556.09/5378.31/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 53s ] thds: 128 tps: 52.00 qps: 5874.00 (r/w/o: 968.00/4906.00/0.00) lat (ms,99%): 2159.29 err/s: 0.00 reconn/s: 0.00
[ 54s ] thds: 128 tps: 76.00 qps: 6606.98 (r/w/o: 1582.00/5024.99/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 55s ] thds: 128 tps: 56.00 qps: 6298.01 (r/w/o: 1106.00/5192.01/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 56s ] thds: 128 tps: 71.00 qps: 6199.97 (r/w/o: 1448.99/4750.97/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 57s ] thds: 128 tps: 49.00 qps: 5917.04 (r/w/o: 905.01/5012.03/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 58s ] thds: 128 tps: 78.00 qps: 6573.88 (r/w/o: 1634.97/4938.91/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 59s ] thds: 128 tps: 40.00 qps: 5620.78 (r/w/o: 753.97/4866.81/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 60s ] thds: 128 tps: 88.00 qps: 6303.15 (r/w/o: 1790.04/4513.11/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 61s ] thds: 128 tps: 35.00 qps: 5955.18 (r/w/o: 622.02/5333.16/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 62s ] thds: 128 tps: 91.00 qps: 6114.00 (r/w/o: 1903.00/4211.00/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 63s ] thds: 128 tps: 26.00 qps: 5827.02 (r/w/o: 529.00/5298.01/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 64s ] thds: 128 tps: 102.00 qps: 6920.96 (r/w/o: 2037.99/4882.97/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 65s ] thds: 128 tps: 28.00 qps: 5721.95 (r/w/o: 449.00/5272.95/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 66s ] thds: 128 tps: 95.00 qps: 6690.97 (r/w/o: 1987.99/4702.98/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 67s ] thds: 128 tps: 31.00 qps: 6038.32 (r/w/o: 567.94/5470.39/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 68s ] thds: 128 tps: 97.01 qps: 6780.56 (r/w/o: 1967.16/4813.40/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 69s ] thds: 128 tps: 37.00 qps: 6298.32 (r/w/o: 720.04/5578.28/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 70s ] thds: 128 tps: 89.00 qps: 6392.97 (r/w/o: 1824.99/4567.98/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 71s ] thds: 128 tps: 41.00 qps: 6800.03 (r/w/o: 781.00/6019.02/0.00) lat (ms,99%): 2082.91 err/s: 0.00 reconn/s: 0.00
[ 72s ] thds: 128 tps: 87.00 qps: 6115.96 (r/w/o: 1795.99/4319.97/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 73s ] thds: 128 tps: 26.00 qps: 5491.03 (r/w/o: 511.00/4980.02/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 74s ] thds: 128 tps: 100.00 qps: 6845.94 (r/w/o: 1940.98/4904.96/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 75s ] thds: 128 tps: 22.00 qps: 4990.05 (r/w/o: 518.00/4472.04/0.00) lat (ms,99%): 2405.65 err/s: 0.00 reconn/s: 0.00
[ 76s ] thds: 128 tps: 89.00 qps: 6616.81 (r/w/o: 1659.95/4956.86/0.00) lat (ms,99%): 2362.72 err/s: 0.00 reconn/s: 0.00
[ 77s ] thds: 128 tps: 41.00 qps: 6340.19 (r/w/o: 908.03/5432.16/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 78s ] thds: 128 tps: 80.00 qps: 6137.94 (r/w/o: 1550.99/4586.96/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 79s ] thds: 128 tps: 46.00 qps: 6699.05 (r/w/o: 986.01/5713.04/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 80s ] thds: 128 tps: 76.00 qps: 5688.66 (r/w/o: 1490.91/4197.75/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 81s ] thds: 128 tps: 51.00 qps: 6582.35 (r/w/o: 1063.06/5519.30/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 82s ] thds: 128 tps: 72.99 qps: 5887.55 (r/w/o: 1427.89/4459.66/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 83s ] thds: 128 tps: 51.00 qps: 6210.44 (r/w/o: 986.07/5224.37/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 84s ] thds: 128 tps: 85.00 qps: 6949.07 (r/w/o: 1728.02/5221.05/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 85s ] thds: 128 tps: 45.00 qps: 6169.02 (r/w/o: 893.00/5276.02/0.00) lat (ms,99%): 2159.29 err/s: 0.00 reconn/s: 0.00
[ 86s ] thds: 128 tps: 88.00 qps: 7085.68 (r/w/o: 1792.92/5292.76/0.00) lat (ms,99%): 2082.91 err/s: 0.00 reconn/s: 0.00
[ 87s ] thds: 128 tps: 42.00 qps: 6168.17 (r/w/o: 842.02/5326.14/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 88s ] thds: 128 tps: 80.00 qps: 6350.11 (r/w/o: 1559.03/4791.09/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 89s ] thds: 128 tps: 46.00 qps: 5927.76 (r/w/o: 925.96/5001.80/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 90s ] thds: 128 tps: 82.00 qps: 6548.19 (r/w/o: 1619.05/4929.15/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 91s ] thds: 128 tps: 45.00 qps: 6034.02 (r/w/o: 948.00/5086.02/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 92s ] thds: 128 tps: 71.00 qps: 6088.04 (r/w/o: 1357.01/4731.03/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 93s ] thds: 128 tps: 58.00 qps: 6999.90 (r/w/o: 1223.98/5775.92/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 94s ] thds: 128 tps: 76.00 qps: 6173.85 (r/w/o: 1498.96/4674.88/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 95s ] thds: 128 tps: 51.00 qps: 6267.08 (r/w/o: 1059.01/5208.07/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 96s ] thds: 128 tps: 75.00 qps: 6428.09 (r/w/o: 1408.02/5020.07/0.00) lat (ms,99%): 2159.29 err/s: 0.00 reconn/s: 0.00
[ 97s ] thds: 128 tps: 53.00 qps: 6307.09 (r/w/o: 1152.02/5155.08/0.00) lat (ms,99%): 2198.52 err/s: 0.00 reconn/s: 0.00
[ 98s ] thds: 128 tps: 63.00 qps: 5789.97 (r/w/o: 1177.99/4611.98/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 99s ] thds: 128 tps: 63.99 qps: 6489.42 (r/w/o: 1280.89/5208.54/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 100s ] thds: 128 tps: 54.00 qps: 5971.67 (r/w/o: 1094.94/4876.73/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 101s ] thds: 128 tps: 70.01 qps: 6512.95 (r/w/o: 1466.21/5046.74/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 102s ] thds: 128 tps: 62.00 qps: 6382.00 (r/w/o: 1082.00/5300.00/0.00) lat (ms,99%): 2082.91 err/s: 0.00 reconn/s: 0.00
[ 103s ] thds: 128 tps: 68.99 qps: 6961.40 (r/w/o: 1537.87/5423.53/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 104s ] thds: 128 tps: 62.01 qps: 6072.53 (r/w/o: 1133.10/4939.43/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
[ 105s ] thds: 128 tps: 67.00 qps: 6626.10 (r/w/o: 1442.02/5184.08/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 106s ] thds: 128 tps: 39.00 qps: 4909.76 (r/w/o: 717.97/4191.80/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 107s ] thds: 128 tps: 84.00 qps: 6653.15 (r/w/o: 1729.04/4924.11/0.00) lat (ms,99%): 2362.72 err/s: 0.00 reconn/s: 0.00
[ 108s ] thds: 128 tps: 38.00 qps: 5721.94 (r/w/o: 665.99/5055.94/0.00) lat (ms,99%): 2405.65 err/s: 0.00 reconn/s: 0.00
[ 109s ] thds: 128 tps: 87.00 qps: 6637.03 (r/w/o: 1825.01/4812.02/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 110s ] thds: 128 tps: 39.00 qps: 6049.07 (r/w/o: 687.01/5362.07/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 111s ] thds: 128 tps: 87.00 qps: 6165.94 (r/w/o: 1824.98/4340.96/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 112s ] thds: 128 tps: 30.00 qps: 5728.99 (r/w/o: 602.00/5126.99/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 113s ] thds: 128 tps: 100.00 qps: 7242.05 (r/w/o: 2002.01/5240.04/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 114s ] thds: 128 tps: 27.00 qps: 5357.98 (r/w/o: 534.00/4823.99/0.00) lat (ms,99%): 2238.47 err/s: 0.00 reconn/s: 0.00
[ 115s ] thds: 128 tps: 87.00 qps: 6278.79 (r/w/o: 1732.94/4545.85/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 116s ] thds: 128 tps: 35.00 qps: 5879.22 (r/w/o: 739.03/5140.19/0.00) lat (ms,99%): 2320.55 err/s: 0.00 reconn/s: 0.00
[ 117s ] thds: 128 tps: 91.00 qps: 6448.99 (r/w/o: 1765.00/4683.99/0.00) lat (ms,99%): 2279.14 err/s: 0.00 reconn/s: 0.00
[ 118s ] thds: 128 tps: 42.00 qps: 6901.01 (r/w/o: 879.00/6022.01/0.00) lat (ms,99%): 2159.29 err/s: 0.00 reconn/s: 0.00
[ 119s ] thds: 128 tps: 95.00 qps: 6864.99 (r/w/o: 1864.00/5000.99/0.00) lat (ms,99%): 2045.74 err/s: 0.00 reconn/s: 0.00
[ 120s ] thds: 128 tps: 30.00 qps: 5734.97 (r/w/o: 641.00/5093.97/0.00) lat (ms,99%): 2120.76 err/s: 0.00 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            151440
        write:                           605760
        other:                           0
        total:                           757200
    transactions:                        7572   (62.65 per sec.)
    queries:                             757200 (6264.74 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          120.8662s
    total number of events:              7572

Latency (ms):
         min:                                  903.19
         avg:                                 2036.54
         max:                                 2663.08
         99th percentile:                     2362.72
         sum:                             15420658.29

Threads fairness:
    events (avg/stddev):           59.1562/0.36
    execution time (avg/stddev):   120.4739/0.23


```


