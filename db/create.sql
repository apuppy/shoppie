-- brand
CREATE TABLE brand(
    id int unsigned not null auto_increment primary key,
    name varchar(32) not null default '' comment 'brand name',
    tel char(11) not null default '' comment 'brand telephone',
    address varchar(64) not null default '' comment 'brand address',
    logo varchar(64) not null default '' comment 'brand logo',
    slogan varchar(30) not null default '' comment 'brand slogan',
    is_opening tinyint(1) not null default 0 comment 'brand opening status 0:closed 1:opening'
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment 'brand';

-- shop
CREATE TABLE shop(
    id int unsigned not null auto_increment primary key,
    brand_id int unsigned not null default 0 comment 'brand id',
    name varchar(32) not null default '' comment 'shop name',
    SN varchar(32) not null default '' comment 'shop serial number',
    create_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset=utf8 comment '店铺';

-- product_category
CREATE TABLE product_category(
    id int unsigned not null auto_increment primary key,
    name varchar(16) not null default '' comment 'category name',
    is_opening tinyint(1) not null default 0 comment 'category opening status 0:closed 1:opening',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment='product category';

-- product 
CREATE TABLE product(
    id int unsigned not null auto_increment primary key,
    name varchar(32) not null default '' comment 'product name',
    cost_price int unsigned not null default 0 comment 'cost price,unit is cent/分',
    selling_price int unsigned not null default 0 comment 'selling price, unit is cent/分',
    left_count int unsigned not null default 0 comment 'product left count',
    sold_count int unsigned not null default 0 comment 'product sold count',
    is_opening tinyint(1) not null default 0 comment 'category opening status 0:closed 1:opening',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment='product';

-- product_detail

CREATE TABLE product_detail(
    id int unsigned not null auto_increment primary key,
    product_id unsigned not null default 0 comment 'product id',
    content text comment 'product detail content',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment='product detail';

-- product_pictures
CREATE TABLE product_picture(
    id int unsigned not null auto_increment primary key,
    product_id unsigned not null default 0 comment 'product id',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment='product pictures';

-- pictures
CREATE TABLE picture(
    id int unsigned not null auto_increment primary key,
    pic_type tinyint(1) unsigned not null default 0 'picture type,0:normal/product picture 1:logo 2:shop picture',
    filename varchar(64) not null default '' comment 'picture filename',
    is_delete tinyint(1) unsigned not null default 0 comment 'picture is deleted 0:not deleted 1:deleted',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment='pictures';




-- customer
CREATE TABLE customer(
    id int unsigned not null auto_increment primary key,
    nickname varchar(16) not null default '' comment 'nickname',
    realname varchar(16) not null default '' comment 'realname',
    mobile char(11) not null default '' comment 'mobile',
    email varchar(32) not null default '' comment 'email',
    avatar_id int unsigned not null default 0 comemnt 'avatar id',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment='customer';

-- avatar
CREATE TABLE avatar(
    id int unsigned not null auto_increment primary key,
    filename varchar(64) not null default '' comment 'avatar filename',
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment='avatar';

-- delivery_address
CREATE TABLE delivery_address(
    id int unsigned not null auto_increment primary key,
    address_id int unsigned not null default 0 comment 'address id',
    province_id int unsigned not null default 0 comment 'province id',
    city_id int unsigned not null default 0 comment 'city id',
    district_id int unsigned not null default 0 comment 'district id',
    detailed_address varchar(64) not null default '' comment 'detailed address',
    is_default tinyint(1) unsigned not null default 0 comment 'default delivery address 0:not default 1:default',
    is_delete tinyint(1) unsigned not null default 0 comment 'delivery address is deleted 0:not deleted 1:deleted',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time'
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
)engine=InnoDB default charset = utf8 comment='delivery address';

-- TODO
-- CREATE TABLE city;

-- shopping cart
CREATE TABLE shopping_cart(
    id int unsigned not null auto_increment primary key,
    customer_id int unsigned not null default 0 comment 'customer id',
    product_id int unsigned not null default 0 comment 'product id',
    unit_price int unsigned not null default 0 comment 'unit price,unit is cent/分',
    purchase_count int unsigned not null default 0 comment 'purchase count',
    is_delete tinyint(1) unsigned not null default 0 comment 'product deleted,0:not deleted 1:deleted',
    updated_at DATETIME not null default CURRENT_TIMESTAMP comment 'updated time'
    created_at DATETIME not null default CURRENT_TIMESTAMP comment 'added time'
);