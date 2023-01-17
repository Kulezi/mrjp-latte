#include <ctype.h>
#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void printInt(int64_t x) {
    printf("%d\n", x);
    fflush(stdout);
}

void error() {
    puts("runtime error");
    fflush(stdout);
    exit(1);
}

void printString(const char *x) {
    puts(x);
    fflush(stdout);
}

int64_t compare(const char *s1, const char *s2) { return strcmp(s1, s2); }

int64_t *newArray(int64_t size) {
    int64_t *arr = calloc(8, size + 1);
    arr[0] = size;
    for (int i = 1; i <= size; i++) arr[i] = i;
    return arr;
}

char *concat(const char *s1, const char *s2) {
    size_t new_len = strlen(s1) + strlen(s2) + 1;
    char *res = malloc(sizeof(char) * new_len);
    strcpy(res, s1);
    strcat(res, s2);
    return res;
}

int64_t readInt() {
    int64_t n;
    scanf("%" SCNd64 " ", &n);
    return n;
}

char *readString() {
    char *buf = NULL;
    size_t len;
    size_t ret = getline(&buf, &len, stdin);
    if (ret == -1) return "";
    if (ret > 0 && buf[ret - 1] == '\n') {
        buf[ret - 1] = '\0';
    }
    return buf;
}