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
