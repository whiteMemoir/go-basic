insert into menus (name, category, description, price, image_url)
values('sate', 'makanan', 'enak', 12000, 'https://www.sasa.co.id/medias/page_medias/cara_bikin_sate_ayam_empuk.jpg'),
	  ('sogem', 'minuman', 'seger', 10000, ''),
	  ('kopi', 'minuman', 'pusing', 15000, 'https://images.immediate.co.uk/production/volatile/sites/30/2020/08/flat-white-3402c4f.jpg?quality=90&webp=true&resize=500,454');
	  
select * from menus;


select name, category, price, image_url 
from menus;

select * from menus
where id = 1;

update menus 
set name = 'Ketoprak telur', price = 14000
where id = 1

delete from menus 
where id = 3;

select * from employees e ;

insert into employees (nip, name, address)
values('0001', 'Budi', 'Jalan A'),
	  ('0002', 'Rina', 'Jalan B');
	  
select * from employees e ;

insert into transactions(menu_id, employee_id, quantity, total)
values(1, 1, 2, 28000),
	  (2, 2, 3, 36000);
	  
select t.id, t.menu_id, t.employee_id, t.quantity, t.total,
	   m.id, m.name, m.price,
	   e.id, e.name
	   from transactions t
	   join employees e on e.id = t.employee_id 
	   join menus m on m.id = t.menu_id