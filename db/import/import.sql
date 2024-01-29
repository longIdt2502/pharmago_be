INSERT INTO account_type (code, title) VALUES ('ADMIN', 'Tổng');
INSERT INTO account_type (code, title) VALUES ('MANAGER', 'Quản lý');
INSERT INTO account_type (code, title) VALUES ('EMPLOYEE', 'Nhân viên');

INSERT INTO company_pharma_type (code, title) VALUES ('PRODUCTION', 'Công ty sản xuất');
INSERT INTO company_pharma_type (code, title) VALUES ('REGISTERED', 'Công ty đăng ký');

INSERT INTO ticket_type (code, title) VALUES ('IMPORT', 'Nhập');
INSERT INTO ticket_type (code, title) VALUES ('EXPORT', 'Xuất');

INSERT INTO ticket_status (code, title) VALUES ('NEW', 'Chờ xác nhận');
INSERT INTO ticket_status (code, title) VALUES ('IN_PROCESS', 'Đang tiến hành');
INSERT INTO ticket_status (code, title) VALUES ('COMPLETE', 'Hoàn thành');
INSERT INTO ticket_status (code, title) VALUES ('CANCEL', 'Bị hủy');

INSERT INTO order_status (code, title) VALUES ('DRAFT', 'Nháp');
INSERT INTO order_status (code, title) VALUES ('IN_PROCESS', 'Đang tiến hành');
INSERT INTO order_status (code, title) VALUES ('COMPLETE', 'Hoàn thành');
INSERT INTO order_status (code, title) VALUES ('CANCEL', 'Huỷ');

INSERT INTO order_type (code, title) VALUES ('SELL', 'Đơn bán');
INSERT INTO order_type (code, title) VALUES ('PRESCRIPTION', 'Đơn thuốc');
INSERT INTO order_type (code, title) VALUES ('IMPORT', 'Đơn nhập');

INSERT INTO payment_item_types (code, title) VALUES ('CASH', 'Tiền mặt');
INSERT INTO payment_item_types (code, title) VALUES ('BANKING', 'Chuyển khoản');
INSERT INTO payment_item_types (code, title) VALUES ('DEBIT', 'Ghi nợ');

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý nhân viên', 'EMPLOYEE', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem danh sách nhân viên', 'EMPLOYEE-1', 'EMPLOYEE', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Tạo mới nhân viên', 'EMPLOYEE-2', 'EMPLOYEE', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem chi tiết nhân viên', 'EMPLOYEE-3', 'EMPLOYEE', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Đổi mật khẩu nhân viên', 'EMPLOYEE-4', 'EMPLOYEE', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Kích hoạt/Vô hiệu hóa nhân viên', 'EMPLOYEE-5', 'EMPLOYEE', 2);

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý đơn hàng', 'ORDER', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem danh sách đơn hàng', 'ORDER-1', 'ORDER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Tạo mới đơn hàng', 'ORDER-2', 'ORDER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem chi tiết đơn hàng', 'ORDER-3', 'ORDER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chỉnh sửa đơn hàng', 'ORDER-4', 'ORDER', 2);

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý sản phẩm', 'PRODUCT', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem danh sách sản phẩm', 'PRODUCT-1', 'PRODUCT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Tạo mới sản phẩm', 'PRODUCT-2', 'PRODUCT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem chi tiết sản phẩm', 'PRODUCT-3', 'PRODUCT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chỉnh sửa sản phẩm', 'PRODUCT-4', 'PRODUCT', 2);
