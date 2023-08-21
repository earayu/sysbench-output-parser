# Sysbench Report Parser

[![Coverage Status](https://coveralls.io/repos/github/yokogawa-k/sysbench-output-parser/badge.svg?branch=master)](https://coveralls.io/github/yokogawa-k/sysbench-output-parser?branch=master)
[![Actions Status](https://github.com/yokogawa-k/sysbench-output-parser/workflows/tests/badge.svg)](https://github.com/yokogawa-k/sysbench-output-parser/actions)

## Overview

```bash
make build

./sysparser --file=./sample/kb_sysbench.log

# detail info
./sysparser --file=./sample/kb_sysbench.log --detail=true

# batch process
./transform.sh $path

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

