CREATE TABLE IF NOT EXISTS `douyu_barrage` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `nickname` varchar(63) NOT NULL DEFAULT 'anoymous' COMMENT '用户昵称',
    `barrage` varchar(513) NOT NULL DEFAULT '' COMMENT '弹幕',
    `ctime` datatime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY('id')
) ENGINE=InnoDB DRFAULT CHARSET=utf8 COMMENT='';