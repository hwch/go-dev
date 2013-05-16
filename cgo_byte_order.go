package main

/*
  #include <stdio.h>
  unsigned int isBigEndian(void)
  {
          typedef union {
                  int i;
                  unsigned char b;
          } orderByte;
          orderByte tp;

          tp.i = 0x12345678;

          if (tp.b == 0x12) {
                  return 1;
          }
          return 0;
  }
*/
import "C"
import "fmt"

func main() {
        if C.isBigEndian() != 0x00 {
                fmt.Printf("Big Endian\n")
        } else {
                fmt.Printf("Little Endian\n")
        }
}
