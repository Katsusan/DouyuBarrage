# DouyuBarrage
======================   
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/toomore/gogrs/master/LICENSE)

A tool that gets Douyu barrages and show them on standard output.

Install
--------------

    go get -u -x github.com/Katsusan/DouyuBarrage

Usage
---------------------

    dybarrage [options]

currently only entrance message and chat message will be shown.


Options
---------------

- `-rid` room id in Douyu. such as : if stream url is "https://douyu.com/97376", then room id will be 97376, default is 97376.


Examples
---------------

Get barrages from douyu.com/60937.

    dybarrage -rid 60937

Snapshot
---------------

![image](snapshot/exmp.png)

TODOs
---------------

- Support more messages shown(such as gifts).
- Support more ways of display.(such as GUI)


Related Project
---------------

A douyu barrage tool(python) also here. [https://github.com/rieuse/DouyuTV](https://github.com/rieuse/DouyuTV)


License
---------------

This package is licensed under MIT license. See LICENSE for details.


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/kkdai/gofbpages/trend.png)](https://bitdeli.com/free "Bitdeli Badge")