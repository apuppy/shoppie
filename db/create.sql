CREATE DATABASE `shoppie` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

-- brands
CREATE TABLE brands(
    id int(11) unsigned not null auto_increment primary key,
    name varchar(30) not null default '' comment 'brand name',
    tel char(11) not null default '' comment 'brand telephone',
    address varchar(60) not null default '' comment 'brand address',
    logo varchar(60) not null default '' comment 'brand logo',
    slogan varchar(30) not null default '' comment 'brand slogan',
    is_opening tinyint(1) unsigned not null default 0 comment 'brand opening status 0:closed 1:opening',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment 'brands';

-- shops
CREATE TABLE shops(
    id int(11) unsigned not null auto_increment primary key,
    brand_id int(11) unsigned not null default 0 comment 'brand id',
    name varchar(30) not null default '' comment 'shop name',
    sn varchar(30) not null default '' comment 'shop serial number',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset=utf8mb4 comment 'shop';

-- shop_staffs
CREATE TABLE shop_staffs(
    id int(11) unsigned not null auto_increment primary key,
    realname varchar(15) not null default '' comment 'staff realname',
    nickname varchar(15) not null default '' comment 'staff nickname',
    mobile varchar(11) not null default '' comment 'mobile number',
    avatar varchar(60) not null default '' comment 'staff avatar',
    motto varchar(30) not null default '' comment 'staff motto',
    bio varchar(60) not null default '' comment 'personal short bio',
    email varchar(30) not null default '' comment 'email',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset=utf8mb4 comment 'shop staff';

-- product_categories
CREATE TABLE product_categories(
    id int(11) unsigned not null auto_increment primary key,
    name varchar(20) not null default '' comment 'category name',
    is_opening tinyint(1) not null default 0 comment 'category opening status 0:closed 1:opening',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='product category';

-- product 
CREATE TABLE products(
    id int(11) unsigned not null auto_increment primary key,
    name varchar(30) not null default '' comment 'product name',
    cost_price int(11) unsigned not null default 0 comment 'cost price,unit is cent/分',
    selling_price int(11) unsigned not null default 0 comment 'selling price, unit is cent/分',
    left_count int(11) unsigned not null default 0 comment 'product left count',
    sold_count int(11) unsigned not null default 0 comment 'product sold count',
    is_opening tinyint(1) not null default 0 comment 'category opening status 0:closed 1:opening',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='product';

-- product_details
CREATE TABLE product_details(
    id int(11) unsigned not null auto_increment primary key,
    product_id int(11) unsigned not null default 0 comment 'product id',
    content text comment 'product detail content',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='product detail';

-- product_pictures
CREATE TABLE product_pictures(
    id int(11) unsigned not null auto_increment primary key,
    product_id int(11) unsigned not null default 0 comment 'product id',
    picture_id int(11) unsigned not null default 0 comment 'picture id',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='product pictures';

-- pictures
CREATE TABLE pictures(
    id int(11) unsigned not null auto_increment primary key,
    pic_type tinyint(1) unsigned not null default 0 comment 'picture type,0:normal/product picture 1:logo 2:shop picture',
    filename varchar(60) not null default '' comment 'picture filename',
    is_delete tinyint(1) unsigned not null default 0 comment 'picture is deleted 0:not deleted 1:deleted',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='pictures';

-- customers
CREATE TABLE customers(
    id int(11) unsigned not null auto_increment primary key,
    nickname varchar(20) not null default '' comment 'nickname',
    realname varchar(20) not null default '' comment 'realname',
    mobile char(11) not null default '' comment 'mobile',
    email varchar(30) not null default '' comment 'email',
    avatar_id int(11) unsigned not null default 0 comment 'avatar id',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='customer';

-- avatars
CREATE TABLE avatars(
    id int(11) unsigned not null auto_increment primary key,
    filename varchar(60) not null default '' comment 'avatar filename',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='avatar';

-- delivery_addresses
CREATE TABLE delivery_addresses(
    id int(11) unsigned not null auto_increment primary key,
    address_id int(11) unsigned not null default 0 comment 'address id',
    province_id int(11) unsigned not null default 0 comment 'province id',
    city_id int(11) unsigned not null default 0 comment 'city id',
    district_id int(11) unsigned not null default 0 comment 'district id',
    detailed_address varchar(60) not null default '' comment 'detailed address',
    is_default tinyint(1) unsigned not null default 0 comment 'default delivery address 0:not default 1:default',
    is_delete tinyint(1) unsigned not null default 0 comment 'delivery address is deleted 0:not deleted 1:deleted',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='delivery address';

-- TODO
-- CREATE TABLE city;

-- shopping_carts
CREATE TABLE shopping_carts(
    id int(11) unsigned not null auto_increment primary key,
    customer_id int(11) unsigned not null default 0 comment 'customer id',
    product_id int(11) unsigned not null default 0 comment 'product id',
    unit_price int(11) unsigned not null default 0 comment 'unit price,unit is cent/分',
    purchase_count int(11) unsigned not null default 0 comment 'purchase count',
    is_delete tinyint(1) unsigned not null default 0 comment 'product deleted,0:not deleted 1:deleted',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='shopping cart';

-- customer_orders
CREATE TABLE customer_orders(
    id int(11) unsigned not null auto_increment primary key,
    customer_id int(11) unsigned not null default 0 comment 'customer id',
    product_id int(11) unsigned not null default 0 comment 'product id',
    unit_price int(11) unsigned not null default 0 comment 'unit price,unit is cent/分',
    purchase_count int(11) unsigned not null default 0 comment 'purchase count',
    discount_price int(11) unsigned not null default 0 comment 'discount price,unit is cent/分',
    pay_progress tinyint(1) unsigned not null default 0 comment 'pay progress, 0:unpaid 1:paying 2:paid',
    delivery_progress tinyint(1) unsigned not null default 0 comment 'delivery progress, 0:none 1:stocking 2:deliverying 3:received',
    pay_price int(11) unsigned not null default 0 comment 'should pay actual price,unit is cent/分',
    is_delete tinyint(1) unsigned not null default 0 comment 'customer order deleted,0:not deleted 1:deleted',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='customer order';

-- customer_pay_orders
CREATE TABLE customer_pay_orders(
    id int(11) unsigned not null auto_increment primary key,
    customer_id int(11) unsigned not null default 0 comment 'customer id',
    delivery_address_id int(11) unsigned not null default 0 comment 'delivery_address id',
    pay_channel tinyint(1) unsigned not null default 0 comment 'pay channel 0:wechat pay 1:alipay 2:bank card 3:cash 4:offline',
    pay_price int(11) unsigned not null default 0 comment 'should pay total price,unit is cent/分',
    pay_progress tinyint(1) unsigned not null default 0 comment 'pay progress, 0:unpaid 1:paying 2:paid',
    paid_at DATETIME not null default CURRENT_TIMESTAMP comment 'successfully paid time',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='customer pay order';

-- customer_pay_order_logs
CREATE TABLE customer_pay_order_logs(
    id int(11) unsigned not null auto_increment primary key,
    customer_id int(11) unsigned not null default 0 comment 'customer id',
    customer_pay_order_id int(11) unsigned not null default 0 comment 'customer_pay_order id',
    callback_message text comment 'pay channel callback message',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    deleted_at DATETIME comment 'deleted time'
)engine=InnoDB default charset = utf8mb4 comment='customer pay order log';

-- customer_pay_order_item
CREATE TABLE customer_pay_order_item(
    id int(11) unsigned not null auto_increment primary key,
    customer_id int(11) unsigned not null default 0 comment 'customer id',
    customer_pay_order_id int(11) unsigned not null default 0 comment 'customer pay order id',
    customer_order_id int(11) unsigned not null default 0 comment 'customer order id',
    pay_price int(11) unsigned not null default 0 comment 'should pay actual price,unit is cent/分',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8mb4 comment='customer pay order item';


-- customer_order_delivery_log
CREATE TABLE customer_order_delivery_log(
    id int(11) unsigned not null auto_increment primary key,
    customer_id int(11) unsigned not null default 0 comment 'customer id',
    customer_order_id int(11) unsigned not null default 0 comment 'customer_order id',
    address_name varchar(16) not null default '' comment 'delivery station address name',
    lat DECIMAL(10, 8) NOT NULL comment 'latitude',
    lng DECIMAL(11, 8) NOT NULL comment 'longitude',
    message varchar(64) not null default '' comment 'delivery message',
    remark varchar(64) not null default '' comment 'remark',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8mb4 comment='customer order delivery log';

-- shop_income
CREATE TABLE shop_income(
    id int(11) unsigned not null auto_increment primary key,
    shop_id int(11) unsigned not null default 0 comment 'shop id',
    total_income int(11) unsigned not null default 0 comment 'total income',
    left_income int(11) unsigned not null default 0 comment 'left income',
    withdrawn_income int(11) unsigned not null default 0 comment 'withdrawn income',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time'
)engine=InnoDB default charset = utf8mb4 comment 'shop income';

-- shop_income_log
CREATE TABLE shop_income_log(
    id int(11) unsigned not null auto_increment primary key,
    customer_pay_order_id int(11) unsigned not null default 0 comment 'relate pay order_id',
    price int(11) unsigned not null default 0 comment 'income price',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time'
)engine=InnoDB default charset = utf8mb4 comment 'shop income history';

-- bank
CREATE TABLE bank(
    id int(11) unsigned not null auto_increment primary key,
    bank_name varchar(16) not null default '' comment 'bank name',
    logo varchar(32) not null default '' comment 'bank logo image',
    is_delete tinyint(1) not null default 0 comment 'is delete,0:not delete 1:deleted'
)engine=InnoDB default charset = utf8mb4 comment 'banks';

-- bank_card
CREATE TABLE bank_card(
    id int(11) unsigned not null auto_increment primary key,
    bank_id int(11) unsigned not null default 0 comment 'bank id',
    card_number varchar(19) not null default '' comment 'card number',
    account_name varchar(16) not null default '' comment 'bank card account name',
    account_mobile varchar(11) not null default '' comment 'account mobile',
    is_default tinyint(1) unsigned not null default 0 comment 'is default,card 0:no 1:yes',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'created time'
)engine=InnoDB default charset = utf8mb4 comment 'bank card or payment tool';