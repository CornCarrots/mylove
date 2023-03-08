CREATE DATABASE if not exists test_love collate utf8_general_ci;

create table invite
(
    id          bigint auto_increment,
    invite_id   bigint           not null comment '邀请ID',
    invite_code varchar(255)     not null comment '邀请码',
    invite_type int default 0    not null comment '邀请方式 1-私密邀请',
    is_delete   bit default b'0' not null,
    PRIMARY KEY (`id`),
    UNIQUE KEY invite_id_idx (invite_id)
)
    comment '邀请表'
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8;

create table invite_user
(
    id        bigint auto_increment,
    user_id   bigint           not null comment '用户id',
    invite_id bigint           not null comment '邀请id',
    is_delete bit default b'0' not null,
    PRIMARY KEY (`id`),
    index user_invite_id_idx (user_id, invite_id)
)
    comment '用户邀请'
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8;

create table note
(
    id          bigint auto_increment,
    note_id     bigint           not null comment '文章id',
    user_id     bigint           not null comment '用户ID',
    content     varchar(500)     not null comment '内容',
    status      int default 0    not null comment '状态 0-生效中 1-已完成',
    note_type   int default 0    not null comment '心愿类型 1-想要的礼物 2-想做的事情',
    is_delete   bit default b'0' not null,
    create_time datetime         null comment '创建时间',
    update_time datetime         null comment '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY note_id_idx (note_id),
    index user_id_idx (user_id)
)
    comment '便利贴'
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8;

create table user
(
    id          bigint auto_increment comment '主键',
    user_id     bigint                not null comment '用户id',
    passport    varchar(255)          not null comment '用户名',
    password    varchar(255)          not null comment '密码',
    nickname    varchar(255)          not null comment '昵称',
    create_time datetime DEFAULT null comment '创建时间',
    update_time datetime DEFAULT null comment '更新时间',
    is_delete   bit      default b'0' not null DEFAULT 0 comment '是否删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY user_id_idx (user_id)
) comment '用户表'
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8;

