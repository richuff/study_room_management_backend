-- 创建数据库（可选）
CREATE DATABASE IF NOT EXISTS student_management;
USE student_management;

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
                                  `user_id` INT AUTO_INCREMENT PRIMARY KEY,
                                  `email` VARCHAR(255) NOT NULL UNIQUE,
`password` VARCHAR(255) NOT NULL,
`is_delete` BOOLEAN DEFAULT FALSE,
`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 学生信息表
CREATE TABLE IF NOT EXISTS `student` (
                                     `id` INT AUTO_INCREMENT PRIMARY KEY,
                                     `academy` VARCHAR(100),
`gender` ENUM('男', '女', '未知') DEFAULT '未知',
`major` VARCHAR(100)
);