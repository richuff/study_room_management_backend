-- 创建数据库
drop database if exists student_management;
create database student_management;
use student_management;

-- 用户表
drop table if exists `student`;
drop table if exists `appointment`;
drop table if exists `user`;
create table `user` (
        `user_id` int auto_increment primary key comment "用户id",
        `email` varchar(255) not null unique comment "邮箱",
        `name` varchar(100) not null comment "姓名",
        `password` varchar(255) not null ,
        `is_delete` boolean default false,
        `created_at` datetime default CURRENT_TIMESTAMP
);
-- 学生信息表
create table `student`(
    `stu_id` int auto_increment primary key comment "学号",
    `user_id` int comment "用户id",
    `gender` enum('男', '女', '未知') default '未知',
    `major` varchar(100) comment "专业",
    foreign key student_user(user_id) references user(user_id)
);

-- 自习室信息表
drop table if exists `room`;
create table `room`(
    `room_id` bigint primary key auto_increment,
    `room_name` varchar(60) unique not null comment "自习室名称",
    `floor` tinyint not null comment "楼层",
    `seat_total`	smallint comment "总座位数",
    `seat_available`	smallint comment 	"可用座位数",
    `room_type`	tinyint	comment "类型 1普通区 2静音区 3讨论区",
    `desc`	varchar(500) comment "简介",
    `open_time`	time	comment "当日开门时间",
    `close_time`	time	comment "关门时间",
    `status`	tinyint	comment "状态",
    `created_at` datetime default CURRENT_TIMESTAMP comment "创建时间",
    `update_at` datetime default CURRENT_TIMESTAMP comment "更新时间"
);

insert into room (
    room_name,
    floor,
    seat_total,
    seat_available,
    room_type,
    `desc`,
    open_time,
    close_time,
    status
) values (
             '静思阁·302',
             3,
             50,
             50,
             2,
             '安静区，提供免费 WiFi、空调与台灯',
             '08:00:00',
             '22:30:00',
             1
         );


-- 预约信息表
create table `appointment`(
   `id` bigint auto_increment primary key comment "预约信息id",
   `user_id` int comment "用户id",
   `room_id` bigint comment "房间id",
  `created_at` datetime default CURRENT_TIMESTAMP comment "创建时间",
   foreign key appointment_user(user_id) references user(user_id),
   foreign key appointment_room(room_id) references room(room_id)
);