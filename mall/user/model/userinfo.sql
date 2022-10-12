CREATE TABLE `userinfo` (
  `userid` varchar(60) NOT NULL,
  `password` varchar(45) NOT NULL,
  `del` tinyint DEFAULT '0',
  PRIMARY KEY (`userid`),
  KEY `pwindex` (`password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci