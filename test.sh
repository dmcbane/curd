#!/bin/sh

echo "./curd -h" > results.txt
./curd -h >> results.txt
echo "./curd clean" >> results.txt
./curd clean >> results.txt
echo "./curd clean --config fred" >> results.txt
./curd clean --config fred >> results.txt
echo "./curd clean --config fred --verbose" >> results.txt
./curd clean --config fred --verbose >> results.txt
echo "./curd clean --verbose" >> results.txt
./curd clean --verbose >> results.txt
echo "./curd ls" >> results.txt
./curd ls >> results.txt
echo "./curd list" >> results.txt
./curd list >> results.txt
echo "./curd ls --config fred" >> results.txt
./curd ls --config fred >> results.txt
echo "./curd list --config fred" >> results.txt
./curd list --config fred >> results.txt
echo "./curd ls --config fred --verbose" >> results.txt
./curd ls --config fred --verbose >> results.txt
echo "./curd list --config fred --verbose" >> results.txt
./curd list --config fred --verbose >> results.txt
echo "./curd ls --verbose" >> results.txt
./curd ls --verbose >> results.txt
echo "./curd list --verbose" >> results.txt
./curd list --verbose >> results.txt
echo "./curd remove" >> results.txt
./curd remove >> results.txt
echo "./curd remove keyword" >> results.txt
./curd remove keyword >> results.txt
echo "./curd remove --config fred" >> results.txt
./curd remove --config fred >> results.txt
echo "./curd remove keyword --config fred" >> results.txt
./curd remove keyword --config fred >> results.txt
echo "./curd remove --config fred --verbose" >> results.txt
./curd remove --config fred --verbose >> results.txt
echo "./curd remove keyword --config fred --verbose" >> results.txt
./curd remove keyword --config fred --verbose >> results.txt
echo "./curd remove keyword --verbose" >> results.txt
./curd remove keyword --verbose >> results.txt
echo "./curd remove --verbose" >> results.txt
./curd remove --verbose >> results.txt
echo "./curd save" >> results.txt
./curd save >> results.txt
echo "./curd save keyword" >> results.txt
./curd save keyword >> results.txt
echo "./curd save --dir tom" >> results.txt
./curd save --dir tom >> results.txt
echo "./curd save keyword --dir tom" >> results.txt
./curd save keyword --dir tom >> results.txt
echo "./curd save --config fred" >> results.txt
./curd save --config fred >> results.txt
echo "./curd save keyword --config fred" >> results.txt
./curd save keyword --config fred >> results.txt
echo "./curd save --config fred --dir tom" >> results.txt
./curd save --config fred --dir tom >> results.txt
echo "./curd save keyword --config fred --dir tom" >> results.txt
./curd save keyword --config fred --dir tom >> results.txt
echo "./curd save --config fred --dir tom --verbose" >> results.txt
./curd save --config fred --dir tom --verbose >> results.txt
echo "./curd save keyword --config fred --dir tom --verbose" >> results.txt
./curd save keyword --config fred --dir tom --verbose >> results.txt
echo "./curd save --dir tom --verbose" >> results.txt
./curd save --dir tom --verbose >> results.txt
echo "./curd save keyword --dir tom --verbose" >> results.txt
./curd save keyword --dir tom --verbose >> results.txt
echo "./curd save keyword --config fred --verbose" >> results.txt
./curd save keyword --config fred --verbose >> results.txt
echo "./curd save --config fred --verbose" >> results.txt
./curd save --config fred --verbose >> results.txt
echo "./curd save keyword --verbose" >> results.txt
./curd save keyword --verbose >> results.txt
echo "./curd save --verbose" >> results.txt
./curd save --verbose >> results.txt
echo "./curd" >> results.txt
./curd >> results.txt
echo "./curd --config fred" >> results.txt
./curd --config fred >> results.txt
echo "./curd --config fred --verbose" >> results.txt
./curd --config fred --verbose >> results.txt
echo "./curd --verbose" >> results.txt
./curd --verbose >> results.txt
echo "./curd john" >> results.txt
./curd john >> results.txt
echo "./curd john --config fred" >> results.txt
./curd john --config fred >> results.txt
echo "./curd john --config fred --verbose" >> results.txt
./curd john --config fred --verbose >> results.txt
echo "./curd john --verbose" >> results.txt
./curd john --verbose >> results.txt
