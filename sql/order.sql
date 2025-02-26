-- 创建商品表
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    version INT NOT NULL DEFAULT 0
);

-- 创建订单表
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- 向 products 表中插入商品 A，库存为 500
INSERT INTO products (name, stock) VALUES ('商品A', 500);

UPDATE products SET name = '商品D' WHERE id = 4;

-- 在 products 表的 id 字段上创建索引
CREATE INDEX idx_products_id ON products (id);
-- 在 products 表的 version 字段上创建索引
CREATE INDEX idx_products_version ON products (version);

select * from products;
+----+---------+-------+---------+
| id | name    | stock | version |
+----+---------+-------+---------+
|  1 | 商品A   |     0 |     500 |
|  2 | 商品B   |     0 |     500 |
|  3 | 商品C   |     0 |    1000 |
|  4 | 商品C   |  2000 |       0 |
+----+---------+-------+---------+

