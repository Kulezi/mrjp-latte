#include <ctype.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

void printInt(int64_t x) { printf("%d\n", x); }

void error() {
    printf("runtime error\n");
    exit(1);
}

void printString(char *x) { printf("%s\n", x); }

int readInt() {
    int res = 0;
    int c = 0;
    while (isspace(c)) {
        c = getchar();
    }

    int mult = 1;
    if (c == '-') {
        mult = -1;
        c = getchar();
    }

    do {
        res = (res * 10) + c - '0';
    } while (c >= '0' && c <= '9');

    if (!isspace(c)) error();

    return res;
}

char extend(char **buf, size_t cap) {
    char *tmp = realloc(buf, sizeof(char) * cap);
    if (tmp == NULL) {
        free(buf);
        error();
    }
    *buf = tmp;
}

char *readString() {
    size_t len = 0;
    size_t cap = 1;
    char *buf = malloc(sizeof(char) * cap);
    if (buf == NULL) error();
    int c = getchar();
    while (c != '\n') {
        if (cap <= len) {
            cap = cap * 2;
            extend(&buf, cap);
        }

        buf[len] = c;
        len++;
        c = getchar();
    }

    if (cap <= len) {
        cap++;
        extend(&buf, cap + 1);
        buf[len] = '\0';
    }

    return buf;
}