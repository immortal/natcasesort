<?php


$array = array(
    "z",
    "p",
    "14",
    "40",
    "30",
    "P",
    "b",
    "f",
    "g",
    "1",
    "x",
    "v",
    "S",
    "s",
    "13",
    "9",
    "7",
    "a",
    "e",
    "E",
    "zz",
    "A",
    "aa");

natcasesort($array);

foreach ($array as $v) {
    echo "$v ";
}
echo PHP_EOL;
