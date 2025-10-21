> [!WARNING]
> Work in progress.

# ztu
An implementation of the ZIPFS compression/decompression specification

```shell
NAME:
   ztu - An implementation of the ZIPFS compression/decompression specification

USAGE:
   ztu [global options] file

DESCRIPTION:
   Examples of use:

   - to compress a file 
   $ztu -o fileCompressed -i dictionaryCID -c file

   - to decompress a file 
   $ztu -o fileDecompressed -d file

   ztu stands for Zeta Tucanae, which is a solar-type star in the constellation Tucana.

GLOBAL OPTIONS:
   --output string, -o string  Output file name
   --compress, -c              compress file
   --decompress, -d            decompress file
   --cid string, -i string     CID of dictionary for compression
   --help, -h                  show help
```
