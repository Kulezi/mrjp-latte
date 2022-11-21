#!/bin/bash
COMPILER="latc_x86_64"
RED='\033[0;31m'
GREEN='\033[0;32m' 
RESET='\033[0m'
OKFILE=$(mktemp)
echo "OK" > $OKFILE
ERRORFILE=$(mktemp)
echo "ERROR" > $ERRORFILE

echo "running make"
make
if [ $? -eq 0 ]; then 
    printf "${GREEN}make succeeded${RESET}\n"
else 
    printf "${RED}make failed${RESET}\n"
fi

PASSED_GOOD=0
TOTAL_GOOD=0

echo "running good tests"
for i in lattests/good/*.lat; do
    TOTAL_GOOD=$((TOTAL_GOOD+1))
    printf "$i: "
    compiler_output=$(mktemp)
    compiler_output_head=$(mktemp)
    ./$COMPILER $i 2> $compiler_output
    exit_code=$1
    if [ $? -eq 0 ]; then
        head -n 1 $compiler_output > $compiler_output_head
        cmp $compiler_output $OKFILE
        if [ $? -eq 0 ]; then
            printf "exit code ${GREEN}$?${RESET}\n"
        else 
            printf "${RED}expected $(cat $OKFILE) got $(cat $compiler_output) ${RESET}\n"
            continue
        fi
    else
        rm $compiler_output
        printf "exit code ${RED}$?${RESET}\n"
        continue
    fi
    rm $compiler_output

    # program_output=$(mktemp)
    # EXEC="${i%.lat}"
    # ./$EXEC > $program_output

    # cmp $program_output "${i%.lat}.output"
    # if [ $? -eq 0 ]; then
    #     printf "got ${GREEN}OK${RESET}\n"
    #     rm $program_output
    # else 
    #     printf "got ${RED}ERROR${RESET}\n"
    # fi

    PASSED_GOOD=$((PASSED_GOOD+1))
    echo ""
done


PASSED_BAD=0
TOTAL_BAD=0


echo "running bad tests"
for i in lattests/bad/*.lat; do
    TOTAL_BAD=$((TOTAL_BAD+1))
    printf "$i: "
    compiler_output=$(mktemp)
    compiler_output_head=$(mktemp)
    ./$COMPILER $i 2> $compiler_output
    exit_code=$?
    if [ $exit_code -eq 0 ]; then
        rm $compiler_output
        printf "exit code ${RED}$?${RESET}\n"
        continue
    else
        head -n 1 $compiler_output > $compiler_output_head
        cmp $compiler_output_head $ERRORFILE
        if [ $? -eq 0 ]; then
            printf "exit code ${GREEN}$exit_code${RESET}\n"
        else 
            printf "${RED}expected $(cat $ERRORFILE) got $(cat $compiler_output) ${RESET}\n"
            continue
        fi
    fi
    rm $compiler_output

    PASSED_BAD=$((PASSED_BAD+1))
    echo ""
done


PASSED_EXT=0
TOTAL_EXT=0


# echo "running bad tests"
# for i in lattests/extensions/*/*.lat; do
#     TOTAL_EXT=$((TOTAL_EXT+1))
#     printf "$i: "
#     compiler_output=$(mktemp)
#     compiler_output_head=$(mktemp)
#     ./$COMPILER $i 2> $compiler_output
#     exit_code=$?
#     if [ $exit_code -eq 0 ]; then
#         rm $compiler_output
#         printf "exit code ${RED}$?${RESET}\n"
#         continue
#     else
#         head -n 1 $compiler_output > $compiler_output_head
#         cmp $compiler_output_head $ERRORFILE
#         if [ $? -eq 0 ]; then
#             printf "exit code ${GREEN}$exit_code${RESET}\n"
#         else 
#             printf "${RED}expected $(cat $ERRORFILE) got $(cat $compiler_output) ${RESET}\n"
#         fi
#     fi
#     rm $compiler_output

#     PASSED_EXT=$((PASSED_EXT+1))
#     echo ""
# done

rm $OKFILE
rm $ERRORFILE

echo "Passed Good: $PASSED_GOOD/$TOTAL_GOOD"
echo "Passed Bad: $PASSED_BAD/$TOTAL_BAD"
