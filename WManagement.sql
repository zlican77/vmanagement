/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50520
Source Host           : localhost:3306
Source Database       : dbtest

Target Server Type    : MYSQL
Target Server Version : 50520
File Encoding         : 65001

Date: 2023-11-30 15:12:59
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `acid`
-- ----------------------------
DROP TABLE IF EXISTS `acid`;
CREATE TABLE `acid` (
  `ID` char(4) NOT NULL DEFAULT '',
  `SL` decimal(7,2) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of acid
-- ----------------------------
INSERT INTO `acid` VALUES ('1', '4428.00');

-- ----------------------------
-- Table structure for `table1`
-- ----------------------------
DROP TABLE IF EXISTS `table1`;
CREATE TABLE `table1` (
  `username` char(20) NOT NULL DEFAULT '',
  `password` int(11) DEFAULT NULL,
  `phone` float DEFAULT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of table1
-- ----------------------------
INSERT INTO `table1` VALUES ('user123', '123', '213');
INSERT INTO `table1` VALUES ('user23', '23', '23');
INSERT INTO `table1` VALUES ('张三', '1', '-50000');
INSERT INTO `table1` VALUES ('李四', '2', '4000');

-- ----------------------------
-- Table structure for `table2`
-- ----------------------------
DROP TABLE IF EXISTS `table2`;
CREATE TABLE `table2` (
  `productID` char(10) NOT NULL DEFAULT '',
  `PDate` date DEFAULT NULL,
  `RKnum` decimal(8,2) DEFAULT NULL,
  `RKRY` char(10) DEFAULT NULL,
  PRIMARY KEY (`productID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of table2
-- ----------------------------
INSERT INTO `table2` VALUES ('1', '2022-07-05', '100.00', 'zhangsan');

-- ----------------------------
-- Table structure for `入库表`
-- ----------------------------
DROP TABLE IF EXISTS `入库表`;
CREATE TABLE `入库表` (
  `入库时间` datetime NOT NULL DEFAULT '2024-11-12 15:00:00',
  `货物号` char(4) NOT NULL DEFAULT '',
  `入库数量` decimal(9,2) DEFAULT NULL,
  `经办人` char(10) DEFAULT NULL,
  PRIMARY KEY (`入库时间`,`货物号`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of 入库表
-- ----------------------------
INSERT INTO `入库表` VALUES ('2022-07-24 03:00:00', '100', '1000.00', '张三2');
INSERT INTO `入库表` VALUES ('2022-08-26 15:59:00', '200', '200.00', '张三1');
INSERT INTO `入库表` VALUES ('2022-09-06 14:04:00', '200', '10.00', '');
INSERT INTO `入库表` VALUES ('2023-11-30 08:55:00', '100', '10.00', '张三1');
INSERT INTO `入库表` VALUES ('2023-11-30 10:18:00', '100', '20.00', '张三3');

-- ----------------------------
-- Table structure for `出库表`
-- ----------------------------
DROP TABLE IF EXISTS `出库表`;
CREATE TABLE `出库表` (
  `出库时间` datetime NOT NULL DEFAULT '2024-11-13 16:00:00',
  `货物号` char(4) NOT NULL DEFAULT '',
  `出库数量` decimal(9,2) DEFAULT NULL,
  `经办人` char(10) DEFAULT NULL,
  PRIMARY KEY (`出库时间`,`货物号`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of 出库表
-- ----------------------------
INSERT INTO `出库表` VALUES ('2022-07-28 03:00:00', '100', '10.00', '张三3');
INSERT INTO `出库表` VALUES ('2022-08-26 15:00:00', '200', '50.00', '张三1');
INSERT INTO `出库表` VALUES ('2022-09-06 14:07:00', '100', '190.00', '张三1');

-- ----------------------------
-- Table structure for `库存表`
-- ----------------------------
DROP TABLE IF EXISTS `库存表`;
CREATE TABLE `库存表` (
  `货物号` char(4) NOT NULL DEFAULT '',
  `更新日期` datetime DEFAULT NULL,
  `库存量` decimal(9,2) DEFAULT NULL,
  PRIMARY KEY (`货物号`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of 库存表
-- ----------------------------
INSERT INTO `库存表` VALUES ('100', '2023-11-30 10:18:00', '830.00');
INSERT INTO `库存表` VALUES ('200', '2022-09-06 14:04:00', '160.00');
INSERT INTO `库存表` VALUES ('210', '2022-07-27 23:59:10', '0.00');
INSERT INTO `库存表` VALUES ('220', '2022-07-27 23:59:35', '0.00');
INSERT INTO `库存表` VALUES ('400', '2022-09-06 14:06:23', '-400.00');

-- ----------------------------
-- Table structure for `货物信息表`
-- ----------------------------
DROP TABLE IF EXISTS `货物信息表`;
CREATE TABLE `货物信息表` (
  `货物号` char(4) NOT NULL DEFAULT '',
  `货物名` char(20) DEFAULT NULL,
  `规格` char(20) DEFAULT NULL,
  `型号` char(20) DEFAULT NULL,
  `说明` char(20) DEFAULT NULL,
  PRIMARY KEY (`货物号`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of 货物信息表
-- ----------------------------
INSERT INTO `货物信息表` VALUES ('100', '钢笔', 'HERO', '', '');
INSERT INTO `货物信息表` VALUES ('200', '铅笔', 'HB', '中华', '');
INSERT INTO `货物信息表` VALUES ('210', '铅笔', '2B', '', '考试用');
INSERT INTO `货物信息表` VALUES ('220', '444', '', '', '');
INSERT INTO `货物信息表` VALUES ('400', '签字笔', '0.5', '黑色', '');
INSERT INTO `货物信息表` VALUES ('500', '新货物', 'new', 'xin', '');

-- ----------------------------
-- View structure for `入库信息`
-- ----------------------------
DROP VIEW IF EXISTS `入库信息`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `入库信息` AS select `入库表`.`入库时间` AS `入库时间`,`入库表`.`货物号` AS `货物号`,`入库表`.`入库数量` AS `入库数量`,`入库表`.`经办人` AS `经办人`,`货物信息表`.`货物名` AS `货物名`,`货物信息表`.`规格` AS `规格`,`货物信息表`.`型号` AS `型号` from (`入库表` left join `货物信息表` on((`货物信息表`.`货物号` = `入库表`.`货物号`))) ;

-- ----------------------------
-- View structure for `出库信息`
-- ----------------------------
DROP VIEW IF EXISTS `出库信息`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `出库信息` AS select `出库表`.`出库时间` AS `出库时间`,`出库表`.`货物号` AS `货物号`,`出库表`.`出库数量` AS `出库数量`,`出库表`.`经办人` AS `经办人`,`货物信息表`.`货物名` AS `货物名`,`货物信息表`.`规格` AS `规格`,`货物信息表`.`型号` AS `型号` from (`出库表` left join `货物信息表` on((`货物信息表`.`货物号` = `出库表`.`货物号`))) ;

-- ----------------------------
-- View structure for `库存信息`
-- ----------------------------
DROP VIEW IF EXISTS `库存信息`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `库存信息` AS select `库存表`.`货物号` AS `货物号`,`库存表`.`更新日期` AS `更新日期`,`库存表`.`库存量` AS `库存量`,`货物信息表`.`货物名` AS `货物名`,`货物信息表`.`规格` AS `规格`,`货物信息表`.`型号` AS `型号` from (`库存表` left join `货物信息表` on((`货物信息表`.`货物号` = `库存表`.`货物号`))) ;
DROP TRIGGER IF EXISTS `TT1`;
DELIMITER ;;
CREATE TRIGGER `TT1` AFTER INSERT ON `入库表` FOR EACH ROW BEGIN
      SET  @hwh=new.货物号;
      SET  @ruSJ=new.入库时间;
      SET  @rksl=new.入库数量  ;
   IF EXISTS(SELECT * FROM 库存表 WHERE 货物号=@hwh) THEN
      UPDATE 库存表 SET 库存量=库存量+@rksl,更新日期=@rusj WHERE 货物号=@hwh ; 
    ELSE
      INSERT INTO 库存表 VALUES(@hwh,@rusj,@rksl) ;
   END IF;
  END

;;
DELIMITER ;
