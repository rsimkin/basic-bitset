# Basic bitset [![Build Status](https://travis-ci.org/rsimkin/basic-bitset.svg?branch=master)](https://travis-ci.org/rsimkin/basic-bitset.svg?branch=master) [![codecov](https://codecov.io/gh/rsimkin/basic-bitset/branch/master/graph/badge.svg)](https://codecov.io/gh/rsimkin/basic-bitset)

Решение написано по мотивам highloadcup 2018, чтобы освоить принцип работы bitset'а, а также попробовать возможности golang.

Использование памяти при хранении _четных_ чисел от 0 до 1 300 000.

| Способ хранения | Использование памяти |
|-----------------|----------------------|
| Basic Bitset    | 68M                  |
| map[uint32]bool | 128M                 |
|[Roaring bitmap](https://roaringbitmap.org/)| 66M|
