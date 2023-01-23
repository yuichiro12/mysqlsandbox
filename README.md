# mysqlsandbox

## usage

### initialize

```
$ docker-compose up -d
```

### invoke deadlock

```
$ go run deadlock.go

## results in deadlock
Error 1213 (40001): Deadlock found when trying to get lock; try restarting transaction
```

### observe deadlock
```
$ make c

mysql> show engine innodb status\G
```

Check out the `LATEST DETECTED DEADLOCK` section and you can see deadlock occurs with PRIMARY key.

```
------------------------
LATEST DETECTED DEADLOCK
------------------------

*** (1) TRANSACTION:
TRANSACTION 4398, ACTIVE 0 sec inserting
mysql tables in use 1, locked 1
LOCK WAIT 6 lock struct(s), heap size 1136, 9 row lock(s), undo log entries 1
MySQL thread id 230, OS thread handle 277761988352, query id 171 172.23.0.1 root update
INSERT IGNORE INTO t1 (name) VALUES ('00411460f7c92d2124a67ea0f4cb5f85'),('00411460f7c92d2124a67ea0f4cb5f85'),('006f52e9102a8d3be2fe5614f42ba989'),('01161aaa0b6d1345dd8fe4e481144d84'),('01161aaa0b6d1345dd8fe4e481144d84'),('01386bd6d8e091c2ab4c7c7de644d37b'),('013a006f03dbc5392effeb8f18fda755'),('013a006f03dbc5392effeb8f18fda755'),('01f78be6f7cad02658508fe4616098a9'),('02522a2b2726fb0a03bb19f2d8d9524d'),('0266e33d3f546cb5436a10798e657d97'),('02e74f10e0327ad868d138f2b4fdd6f0'),('02e74f10e0327ad868d138f2b4fdd6f0'),('0353ab4cbed5beae847a7ff6e220b5cf'),('0353ab4cbed5beae847a7ff6e220b5cf'),('03afdbd66e7929b125f8597834fa83a4'),('03afdbd66e7929b125f8597834fa83a4'),('03afdbd66e7929b125f8597834fa83a4'),('03afdbd66e7929b125f8597834fa83a4'),('03c6b06952c750899bb03d998e631860'),('04025959b191f8f9de3f924f0940515f'),('04025959b191f8f9de3f924f0940515f'),('043c3d7e489c69b48737cc0c92d0f3a2'),('045117b0e0a11a242b9765e79cbf113f'),('04ecb1fa2

*** (1) HOLDS THE LOCK(S):
RECORD LOCKS space id 14 page no 13 n bits 136 index PRIMARY of table `test`.`t1` trx id 4398 lock_mode X locks gap before rec
Record lock, heap no 64 PHYSICAL RECORD: n_fields 4; compact format; info bits 0
 0: len 8; hex 8000000000002329; asc       #);;
 1: len 6; hex 00000000112e; asc      .;;
 2: len 7; hex 82000000f70110; asc        ;;
 3: len 30; hex 303235323261326232373236666230613033626231396632643864393532; asc 02522a2b2726fb0a03bb19f2d8d952; (total 32 bytes);


*** (1) WAITING FOR THIS LOCK TO BE GRANTED:
RECORD LOCKS space id 14 page no 8 n bits 416 index name_unique of table `test`.`t1` trx id 4398 lock mode S waiting
Record lock, heap no 244 PHYSICAL RECORD: n_fields 2; compact format; info bits 0
 0: len 30; hex 303235323261326232373236666230613033626231396632643864393532; asc 02522a2b2726fb0a03bb19f2d8d952; (total 32 bytes);
 1: len 8; hex 80000000000007d1; asc         ;;


*** (2) TRANSACTION:
TRANSACTION 4391, ACTIVE 0 sec inserting
mysql tables in use 1, locked 1
LOCK WAIT 41 lock struct(s), heap size 24784, 25 row lock(s), undo log entries 1
MySQL thread id 223, OS thread handle 277772637952, query id 164 172.23.0.1 root update
INSERT IGNORE INTO t1 (name) VALUES ('00411460f7c92d2124a67ea0f4cb5f85'),('00411460f7c92d2124a67ea0f4cb5f85'),('00411460f7c92d2124a67ea0f4cb5f85'),('006f52e9102a8d3be2fe5614f42ba989'),('00ac8ed3b4327bdd4ebbebcb2ba10a00'),('01161aaa0b6d1345dd8fe4e481144d84'),('013a006f03dbc5392effeb8f18fda755'),('013d407166ec4fa56eb1e1f8cbe183b9'),('01882513d5fa7c329e940dda99b12147'),('01882513d5fa7c329e940dda99b12147'),('01882513d5fa7c329e940dda99b12147'),('019d385eb67632a7e958e23f24bd07d7'),('01f78be6f7cad02658508fe4616098a9'),('024d7f84fff11dd7e8d9c510137a2381'),('02522a2b2726fb0a03bb19f2d8d9524d'),('02e74f10e0327ad868d138f2b4fdd6f0'),('03c6b06952c750899bb03d998e631860'),('043c3d7e489c69b48737cc0c92d0f3a2'),('043c3d7e489c69b48737cc0c92d0f3a2'),('04ecb1fa28506ccb6f72b12c0245ddbc'),('05049e90fa4f5039a8cadc6acbb4b2cc'),('051e4e127b92f5d98d3c79b195f2b291'),('0584ce565c824b7b7f50282d9a19945b'),('05f971b5ec196b8c65b75d2ef8267331'),('05f971b5e

*** (2) HOLDS THE LOCK(S):
RECORD LOCKS space id 14 page no 8 n bits 416 index name_unique of table `test`.`t1` trx id 4391 lock_mode X locks rec but not gap
Record lock, heap no 244 PHYSICAL RECORD: n_fields 2; compact format; info bits 0
 0: len 30; hex 303235323261326232373236666230613033626231396632643864393532; asc 02522a2b2726fb0a03bb19f2d8d952; (total 32 bytes);
 1: len 8; hex 80000000000007d1; asc         ;;


*** (2) WAITING FOR THIS LOCK TO BE GRANTED:
RECORD LOCKS space id 14 page no 13 n bits 136 index PRIMARY of table `test`.`t1` trx id 4391 lock_mode X locks gap before rec insert intention waiting
Record lock, heap no 64 PHYSICAL RECORD: n_fields 4; compact format; info bits 0
 0: len 8; hex 8000000000002329; asc       #);;
 1: len 6; hex 00000000112e; asc      .;;
 2: len 7; hex 82000000f70110; asc        ;;
 3: len 30; hex 303235323261326232373236666230613033626231396632643864393532; asc 02522a2b2726fb0a03bb19f2d8d952; (total 32 bytes);

```
