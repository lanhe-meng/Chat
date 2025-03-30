# Chat
user

```mysql
CREATE TABLE `user` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,   -- 唯一用户ID（主键）
  `username` VARCHAR(50) NOT NULL UNIQUE,       -- 用户名（唯一）
  `email` VARCHAR(100) NOT NULL UNIQUE,         -- 邮箱（唯一，可选）
  `password_hash` VARCHAR(255) NOT NULL,        -- 密码哈希值（非明文存储）
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间（自动生成）
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 最后更新时间
  `last_login_at` TIMESTAMP NULL,               -- 最后登录时间
  `is_active` TINYINT(1) DEFAULT 1,             -- 账户状态（1=激活，0=禁用）
  `avatar_url` VARCHAR(255) DEFAULT NULL,       -- 用户头像地址
  PRIMARY KEY (`id`),
  INDEX `idx_username` (`username`),
  INDEX `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

