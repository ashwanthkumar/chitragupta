--  Default Schema of the DB which is available as part of the repo.

CREATE TABLE IF NOT EXISTS users (
  `uid` INTEGER PRIMARY KEY AUTOINCREMENT, -- Just for Visual Display on the page
  `name` TEXT
);

CREATE TABLE IF NOT EXISTS accounts (
  `aid` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` TEXT,
  `user_id` INTEGER,
  `type` INTEGER,
  FOREIGN KEY (`user_id`) REFERENCES users (`uid`)
);

CREATE TABLE IF NOT EXISTS transactions (
  `tid` INTEGER PRIMARY KEY AUTOINCREMENT,
  `date` INTEGER,
  `amount` REAL,
  `account_id` INTEGER, -- Store JSON serialized blob of Key-Value pairs
  `metadata` TEXT,
  FOREIGN KEY (`account_id`) REFERENCES accounts (`aid`)
);

CREATE TABLE IF NOT EXISTS tags (
  `tag_id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` TEXT,
  `user_id` INTEGER,
  FOREIGN KEY (`user_id`) REFERENCES users (`uid`)
);

CREATE TABLE IF NOT EXISTS transactions_has_tags (
  `transaction_id` INTEGER,
  `tag_id` INTEGER,
  FOREIGN KEY (`transaction_id`) REFERENCES transactions (`tid`),
  FOREIGN KEY (`tag_id`) REFERENCES tags (`tag_id`)
);
