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

$array = array("z1", "z10", "z100", "z101", "z102", "z11", "z12", "z13", "z14", "z15", "z16", "z17", "z18", "z19", "z2", "z20", "z3", "z4", "z5", "z6", "z7", "z8", "z9");
natcasesort($array);
foreach ($array as $v) {
    echo "$v ";
}
echo PHP_EOL;
