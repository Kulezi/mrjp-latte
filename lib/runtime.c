#include <ctype.h>
#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void printInt(int32_t x) {
    printf("%d\n", x);
    fflush(stdout);
}

void error() {
    puts("runtime error");
    fflush(stdout);
    exit(1);
}

void printString(const char *x) {
    if (x == NULL) x = "";
    puts(x);
    fflush(stdout);
}

int32_t compare(const char *s1, const char *s2) {
    if (s1 == NULL) s1 = "";
    if (s2 == NULL) s2 = "";

    if (strcmp(s1, s2)) {
        return 1;
    } else {
        return 0;
    }
}

int32_t *newArray(int32_t size) {
    int32_t *arr = calloc(size + 1, 4);
    arr[0] = size;
    return arr;
}

int32_t *newClass(int32_t size, int32_t vtable) {
    int32_t *class = calloc(size + 1, 4);
    class[0] = vtable;
    return class;
}

char *concat(const char *s1, const char *s2) {
    if (s1 == NULL) s1 = "";
    if (s2 == NULL) s2 = "";
    size_t new_len = strlen(s1) + strlen(s2) + 1;
    char *res = malloc(sizeof(char) * new_len);
    strcpy(res, s1);
    strcat(res, s2);
    return res;
}

int32_t readInt() {
    int32_t n;
    scanf("%d ", &n);
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