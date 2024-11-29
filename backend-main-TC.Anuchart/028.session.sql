DROP VIEW IF EXISTS v_older_total_more_500;

CREATE VIEW v_older_total_more_500 AS
SELECT c.cus_firstname as ชื่อ, 
       c.cus_lastname as นามสกุล, 
       o.order_id AS รหัสสั่งซื้อ, 
       o.order_date AS วันสั่งซื้อ, 
       o.order_total AS จำนวนที่สั่งซื้อ 
FROM tbl_customer c 
INNER JOIN tbl_orders o 
ON c.cus_id = o.cus_id 
WHERE o.order_total >= 500;
