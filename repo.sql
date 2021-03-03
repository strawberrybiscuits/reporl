CREATE DATABASE IF NOT EXISTS reporl DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
/*
存放获取到的release版本信息
*/
CREATE TABLE `repo_rl` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `repoName` varchar(255) NOT NULL COMMENT '/apache/shiro',
  `title` varchar(255) NOT NULL,
  `link` varchar(512) NOT NULL,
  `description` varchar(512) NOT NULL COMMENT 'content tag',
  `createTime` datetime NOT NULL,
  `simpleName` varchar(255) NOT NULL COMMENT 'shiro',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*
存放rss订阅仓库信息
*/
CREATE TABLE `repo_rss` (
  `id` int(20) NOT NULL,
  `repoUrl` varchar(255) NOT NULL COMMENT 'https://github.com/apache/shiro',
  `repoName` varchar(255) NOT NULL COMMENT 'apache/shiro',
  `simpleName` varchar(255) NOT NULL COMMENT 'shiro',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;