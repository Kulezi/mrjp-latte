#include <ctype.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void printInt(int64_t x) { printf("%d\n", x); }

void error() {
    puts("runtime error");
    exit(1);
}

void printString(const char *x) { puts(x); }

int64_t compare(const char *s1, const char *s2) { return strcmp(s1, s2); }

char *concat(const char *s1, const char *s2) {
    size_t new_len = strlen(s1) + strlen(s2) + 1;
    char *res = malloc(sizeof(char) * new_len);
    strcpy(res, s1);
    strcat(res, s2);
    return res;
}

int readInt() {
    int n;
    scanf("%d ", &n);
    return n;
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