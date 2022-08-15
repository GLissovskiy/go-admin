package models

type Permission struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

/*
INSERT INTO `go_admin`.`permissions` (`id`, `name`) VALUES ('1', 'view_users');
INSERT INTO `go_admin`.`permissions` (`id`, `name`) VALUES ('2', 'edit_users');
INSERT INTO `go_admin`.`permissions` (`id`, `name`) VALUES ('3', 'view_roles');
INSERT INTO `go_admin`.`permissions` (`id`, `name`) VALUES ('4', 'edit_roles');
INSERT INTO `go_admin`.`permissions` (`id`, `name`) VALUES ('5', 'view_products');
INSERT INTO `go_admin`.`permissions` (`id`, `name`) VALUES ('6', 'edit_products');
INSERT INTO `go_admin`.`permissions` (`id`, `name`) VALUES ('7', 'view_orders');
INSERT INTO `go_admin`.`permissions` (`id`, `name`) VALUES ('8', 'edit_orders');

*/
