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
TRANSACTION 4438, ACTIVE 1 sec inserting
mysql tables in use 1, locked 1
LOCK WAIT 6 lock struct(s), heap size 1136, 5 row lock(s)
MySQL thread id 245, OS thread handle 277783058176, query id 200 172.23.0.1 root update
INSERT IGNORE INTO t1 (name) VALUES ('00411460f7c92d2124a67ea0f4cb5f85'),('00411460f7c92d2124a67ea0f4cb5f85'),('00ac8ed3b4327bdd4ebbebcb2ba10a00'),('00ac8ed3b4327bdd4ebbebcb2ba10a00'),('00ac8ed3b4327bdd4ebbebcb2ba10a00'),('00ac8ed3b4327bdd4ebbebcb2ba10a00'),('013a006f03dbc5392effeb8f18fda755'),('013a006f03dbc5392effeb8f18fda755'),('013d407166ec4fa56eb1e1f8cbe183b9'),('01882513d5fa7c329e940dda99b12147'),('024d7f84fff11dd7e8d9c510137a2381'),('0266e33d3f546cb5436a10798e657d97'),('02a32ad2669e6fe298e607fe7cc0e1a0'),('02a32ad2669e6fe298e607fe7cc0e1a0'),('02a32ad2669e6fe298e607fe7cc0e1a0'),('0336dcbab05b9d5ad24f4333c7658a0e'),('0336dcbab05b9d5ad24f4333c7658a0e'),('0336dcbab05b9d5ad24f4333c7658a0e'),('0353ab4cbed5beae847a7ff6e220b5cf'),('03afdbd66e7929b125f8597834fa83a4'),('03c6b06952c750899bb03d998e631860'),('043c3d7e489c69b48737cc0c92d0f3a2'),('045117b0e0a11a242b9765e79cbf113f'),('045117b0e0a11a242b9765e79cbf113f'),('04ecb1fa2

*** (1) HOLDS THE LOCK(S):
RECORD LOCKS space id 15 page no 15 n bits 176 index PRIMARY of table `test`.`t1` trx id 4438 lock_mode X locks gap before rec
Record lock, heap no 19 PHYSICAL RECORD: n_fields 4; compact format; info bits 0
 0: len 8; hex 8000000000003a99; asc       : ;;
 1: len 6; hex 000000001157; asc      W;;
 2: len 7; hex 82000001120110; asc        ;;
 3: len 30; hex 303133643430373136366563346661353665623165316638636265313833; asc 013d407166ec4fa56eb1e1f8cbe183; (total 32 bytes);


*** (1) WAITING FOR THIS LOCK TO BE GRANTED:
RECORD LOCKS space id 15 page no 15 n bits 176 index PRIMARY of table `test`.`t1` trx id 4438 lock_mode X locks gap before rec insert intention waiting
Record lock, heap no 19 PHYSICAL RECORD: n_fields 4; compact format; info bits 0
 0: len 8; hex 8000000000003a99; asc       : ;;
 1: len 6; hex 000000001157; asc      W;;
 2: len 7; hex 82000001120110; asc        ;;
 3: len 30; hex 303133643430373136366563346661353665623165316638636265313833; asc 013d407166ec4fa56eb1e1f8cbe183; (total 32 bytes);


*** (2) TRANSACTION:
TRANSACTION 4435, ACTIVE 1 sec inserting
mysql tables in use 1, locked 1
LOCK WAIT 8 lock struct(s), heap size 1136, 4 row lock(s)
MySQL thread id 243, OS thread handle 277774407424, query id 198 172.23.0.1 root update
INSERT IGNORE INTO t1 (name) VALUES ('00411460f7c92d2124a67ea0f4cb5f85'),('006f52e9102a8d3be2fe5614f42ba989'),('006f52e9102a8d3be2fe5614f42ba989'),('01161aaa0b6d1345dd8fe4e481144d84'),('01386bd6d8e091c2ab4c7c7de644d37b'),('01386bd6d8e091c2ab4c7c7de644d37b'),('01386bd6d8e091c2ab4c7c7de644d37b'),('013a006f03dbc5392effeb8f18fda755'),('013a006f03dbc5392effeb8f18fda755'),('01882513d5fa7c329e940dda99b12147'),('01882513d5fa7c329e940dda99b12147'),('01882513d5fa7c329e940dda99b12147'),('019d385eb67632a7e958e23f24bd07d7'),('019d385eb67632a7e958e23f24bd07d7'),('019d385eb67632a7e958e23f24bd07d7'),('024d7f84fff11dd7e8d9c510137a2381'),('0266e33d3f546cb5436a10798e657d97'),('0353ab4cbed5beae847a7ff6e220b5cf'),('03c6b06952c750899bb03d998e631860'),('03c6b06952c750899bb03d998e631860'),('04025959b191f8f9de3f924f0940515f'),('045117b0e0a11a242b9765e79cbf113f'),('04ecb1fa28506ccb6f72b12c0245ddbc'),('04ecb1fa28506ccb6f72b12c0245ddbc'),('051e4e127

*** (2) HOLDS THE LOCK(S):
RECORD LOCKS space id 15 page no 15 n bits 176 index PRIMARY of table `test`.`t1` trx id 4435 lock_mode X locks gap before rec
Record lock, heap no 19 PHYSICAL RECORD: n_fields 4; compact format; info bits 0
 0: len 8; hex 8000000000003a99; asc       : ;;
 1: len 6; hex 000000001157; asc      W;;
 2: len 7; hex 82000001120110; asc        ;;
 3: len 30; hex 303133643430373136366563346661353665623165316638636265313833; asc 013d407166ec4fa56eb1e1f8cbe183; (total 32 bytes);


*** (2) WAITING FOR THIS LOCK TO BE GRANTED:
RECORD LOCKS space id 15 page no 15 n bits 176 index PRIMARY of table `test`.`t1` trx id 4435 lock_mode X locks gap before rec insert intention waiting
Record lock, heap no 19 PHYSICAL RECORD: n_fields 4; compact format; info bits 0
 0: len 8; hex 8000000000003a99; asc       : ;;
 1: len 6; hex 000000001157; asc      W;;
 2: len 7; hex 82000001120110; asc        ;;
 3: len 30; hex 303133643430373136366563346661353665623165316638636265313833; asc 013d407166ec4fa56eb1e1f8cbe183; (total 32 bytes);

```
