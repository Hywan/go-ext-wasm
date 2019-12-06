// Compiled with `wasicc wasi_c.cpp -o wasi_c`.

#include <stdio.h>

extern "C" {
  int32_t sum(int32_t, int32_t) __attribute__((used));
  int32_t sum(int32_t x, int32_t y) {
    return x + y;
  }
}

int main(int argc, char **argv) {
  if (argc < 2) {
    printf("Hello, WASI!\n");
  } else {
    printf("Hello, %s!\n", argv[1]);
  }
}
