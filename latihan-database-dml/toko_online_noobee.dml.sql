insert into users(username, email, nama_lengkap)
values('user1', 'user1@example.com', 'User Satu'),
	  ('user2', 'user2@example.com', 'User Dua'),
	  ('user3', 'user3@example.com', 'User Tiga');
	  
insert into orders(user_id, status)
values(1, 'Dalam Pengiriman'),
	  (2, 'Selesai'),
	  (3, 'Dibatalkan');
	  
insert into order_items(order_id, product_name, quantity, harga_per_item)
values(1, 'Produk A', 2, 50000),
	  (1, 'Produk B', 3, 30000),
	  (2, 'Produk C', 1, 75000),
	  (2, 'Produk D', 2, 60000);
	  	  
select o.order_id, u.username
from orders o
join users u on u.user_id = o.user_id; 

select o.order_id, u.username, sum(oi.quantity * oi.harga_per_item) as total_harga
from orders o
join users u on o.user_id = u.user_id
join order_items oi on o.order_id = oi.order_id
group by o.order_id, u.username;

