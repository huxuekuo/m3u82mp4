CREATE DATABASE video_server CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
  
create table v_user(
	id bigint unsigned primary key comment "主键",
	account varchar(30) comment "账号",
	password varchar(10) default "" comment "密码,md5加密",
	username varchar(30) default "" not null comment "名称",
	avatar varchar(255) default "" not null comment "头像",
	is_del tinyint unsigned default 0 not null comment "逻辑删除",
	createtime bigint default null comment "创建时间",
	updatetime bigint default null comment "修改时间"
)CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;