/*========================>database user <===================================*/
CREATE DATABASE user;
USE user;

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '电子邮箱',
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `avatar_id` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `question` varchar(100) NOT NULL DEFAULT '' COMMENT '找回密码问题',
  `answer` varchar(100) NOT NULL DEFAULT '' COMMENT '找回密码答案',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_phone` (`phone`),
  UNIQUE KEY `uniq_username` (`username`),
  KEY `ix_update_at` (`update_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `user_receive_address` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '收货人名称',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否为默认地址',
  `post_code` varchar(100) NOT NULL DEFAULT '' COMMENT '邮政编码',
  `province` varchar(100) NOT NULL DEFAULT '' COMMENT '省份/直辖市',
  `city` varchar(100) NOT NULL DEFAULT '' COMMENT '城市',
  `district` varchar(100) NOT NULL DEFAULT '' COMMENT '区',
  `detail_address` varchar(128) NOT NULL DEFAULT '' COMMENT '详细地址(街道)',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '数据创建时间[禁止在代码中赋值]',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '数据更新时间[禁止在代码中赋值]',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户收货地址表';

/*========================>database product <===================================*/
CREATE DATABASE product;
USE product;

CREATE TABLE `image`(
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '图片id',
    `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品Id',
    `url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片地址',
    `type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '图片类型1-单图 2-详情图',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品图片表';

CREATE TABLE `product` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品id',
    `category_id` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类别Id',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
    `subtitle` varchar(200) NOT NULL DEFAULT '' COMMENT '商品副标题',
    `images` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片地址,逗号分隔',
    `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '商品详情',
    `price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '价格,单位-元保留两位小数',
    `stock` int(11) NOT NULL DEFAULT 0 COMMENT '库存数量',
    `status` int(6) NOT NULL DEFAULT 1 COMMENT '商品状态.1-在售 2-下架 3-删除',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_category_id` (`category_id`),
    KEY `ix_update_at` (`update_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';


CREATE TABLE `product_category` (
    `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
    `parent_id` smallint(6) NOT NULL DEFAULT 0 COMMENT '父类别id当id=0时说明是根节点,一级类别',
    `name` varchar(50) NOT NULL DEFAULT '' COMMENT '类别名称',
    `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '类别状态1-正常,2-已废弃',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品类别表';

CREATE TABLE `banner` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '轮播图id',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '轮播图名称',
    `image_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片id',
    `link_url` varchar(255) NOT NULL DEFAULT '' COMMENT '链接地址',
    `description` varchar(255) NOT NULL DEFAULT '' COMMENT '轮播图描述',
    `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态1-显示 2-隐藏',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='轮播图表';
/*========================>database order<===================================*/
CREATE DATABASE `order`;
USE `order`;

CREATE TABLE `order` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单id',
    `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `total_price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '订单总价',
    `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '订单状态1-待支付 2-已支付 3-已发货 4-已完成 5-已取消',
    `address_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收货地址',
    `post_cost` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '邮费',
    `payment` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '实付金额',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

/*每个订单可能包含多个订单项，比如一个订单多个商品*/
CREATE TABLE `order_item` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单项id',
    `order_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
    `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
    `product_name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
    `product_image_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品图片',
    `unit_price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '商品单价',
    `quantity` int(11) NOT NULL DEFAULT 0 COMMENT '商品数量',
    `total_price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '商品总价',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_order_id` (`order_id`),
    KEY `ix_product_id` (`product_id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单项表';

CREATE TABLE `shipping` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '收货信息表id',
    `order_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
    `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `receiver_name` varchar(20) NOT NULL DEFAULT '' COMMENT '收货姓名',
    `receiver_phone` varchar(20) NOT NULL DEFAULT '' COMMENT '收货固定电话',
    `receiver_mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '收货移动电话',
    `receiver_province` varchar(20) NOT NULL DEFAULT '' COMMENT '省份',
    `receiver_city` varchar(20) NOT NULL DEFAULT '' COMMENT '城市',
    `receiver_district` varchar(20) NOT NULL DEFAULT '' COMMENT '区/县',
    `receiver_address` varchar(200) NOT NULL DEFAULT '' COMMENT '详细地址',
    `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_order_id` (`order_id`),
    KEY `ix_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收货信息表';


/*========================>database cart<===================================*/
CREATE DATABASE cart;
USE cart;

CREATE TABLE `cart` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '购物车id',
    `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
    `quantity` int(11) NOT NULL DEFAULT 0 COMMENT '商品数量',
    `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_user_id` (`user_id`),
    KEY `ix_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车表';


/*========================>database payment<===================================*/
CREATE DATABASE payment;
USE payment;

CREATE TABLE `payment` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '支付表id',
    `order_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
    `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `amount` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '支付金额',
    `payment_method` tinyint(4) NOT NULL DEFAULT 1 COMMENT '支付方式1-支付宝 2-微信 3-银行卡',
    `platform_number` varchar(200) NOT NULL DEFAULT '' COMMENT '支付流水号',
    `platform_status` varchar(20) NOT NULL DEFAULT '' COMMENT '支付状态',
    `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '支付状态1-待支付 2-已支付 3-支付失败',
    `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_order_id` (`order_id`),
    KEY `ix_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付表';


/*========================>database comment<===================================*/
CREATE DATABASE comment;
USE comment;

CREATE TABLE `image`(
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '图片id',
    `comment_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论Id',
    `url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片地址',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_comment_id` (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论图片表';


CREATE TABLE `comment` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论id',
    `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
    `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级评论id',
    `content` text NOT NULL COMMENT '评论内容',
    `rating` tinyint(4) NOT NULL DEFAULT 0 COMMENT '评分',
    `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_user_id` (`user_id`),
    KEY `ix_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';
