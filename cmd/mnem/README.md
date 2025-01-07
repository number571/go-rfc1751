# Mnem

### Install

```bash
go install github.com/number571/go-rfc1751/cmd/mnem@latest
```

### Usage

Generate mnemonic

```bash
usage: 
    mnem [-size=256] [-raw=false]
```

Convert mnemonic

```bash
usage: 
    mnem --conv [-raw=false]
stdin:
    [string]EOL
```

### Examples

#### 1. Generate

```bash
mnem
ARCH CUE FILM AWL FINK MUDD LARK GURU DIP ALAN LEEK MOOD FELT HUG MAW CASH TOLD AIDS BLOC LOSE WE ROW THEY REAR
```

```bash
mnem -size=128
NEAT CELL DIN GRIT AFAR BETH FRAY FOAL MUCK MEG DANA MYTH
```

```bash
mnem -size=128 -raw
4dc48252b1836e616a2b2a6154c5eadc
```

#### 2. Convert

```bash
echo "MAIL HERS YAP SULK LACK SHAW BODE SAY SWUM IOWA HOP GINA" | mnem -conv
b612e91a727a91b45aa75f9a4f81bd17
```

```bash
echo "b612e91a727a91b45aa75f9a4f81bd17" | mnem -conv -raw
MAIL HERS YAP SULK LACK SHAW BODE SAY SWUM IOWA HOP GINA
```
