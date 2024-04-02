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

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý khách hàng', 'CUSTOMER', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem danh sách khách hàng', 'CUSTOMER-1', 'CUSTOMER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Tạo mới khách hàng', 'CUSTOMER-2', 'CUSTOMER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem chi tiết khách hàng', 'CUSTOMER-3', 'CUSTOMER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chỉnh sửa khách hàng', 'CUSTOMER-4', 'CUSTOMER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xoá khách hàng', 'CUSTOMER-5', 'CUSTOMER', 2);

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý nhóm khách hàng', 'G-CUSTOMER', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem danh sách nhóm khách hàng', 'G-CUSTOMER-1', 'G-CUSTOMER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Tạo mới nhóm khách hàng', 'G-CUSTOMER-2', 'G-CUSTOMER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem chi tiết nhóm khách hàng', 'G-CUSTOMER-3', 'G-CUSTOMER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chỉnh sửa nhóm khách hàng', 'G-CUSTOMER-4', 'G-CUSTOMER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xoá nhóm khách hàng', 'G-CUSTOMER-5', 'G-CUSTOMER', 2);

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý tồn kho', 'INVENTORY', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Điều chỉnh tồn kho', 'INVENTORY-1', 'INVENTORY', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xuất dữ liệu', 'INVENTORY-2', 'INVENTORY', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Danh sách tồn kho', 'INVENTORY-3', 'INVENTORY', 2);

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý nhập kho', 'IMPORT', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Tạo mới nhập kho', 'IMPORT-1', 'IMPORT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chi tiết nhập kho', 'IMPORT-2', 'IMPORT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chỉnh sửa nhập kho', 'IMPORT-3', 'IMPORT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Huỷ đơn nhập kho', 'IMPORT-4', 'IMPORT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xác nhận đơn nhập kho', 'IMPORT-5', 'IMPORT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Danh sách nhập kho', 'IMPORT-6', 'IMPORT', 2);

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý xuất kho', 'EXPORT', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Danh sách đơn xuất kho', 'EXPORT-1', 'EXPORT', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chi tiết đơn xuất kho', 'EXPORT-2', 'EXPORT', 2);

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý nhà cung cấp', 'SUPPLIER', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem danh sách nhà cung cấp', 'SUPPLIER-1', 'SUPPLIER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Tạo mới nhà cung cấp', 'SUPPLIER-2', 'SUPPLIER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem chi tiết nhà cung cấp', 'SUPPLIER-3', 'SUPPLIER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chỉnh sửa nhà cung cấp', 'SUPPLIER-4', 'SUPPLIER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xoá nhà cung cấp', 'SUPPLIER-5', 'SUPPLIER', 2);

INSERT INTO apps (title, code, parent, level) VALUES ('Quản lý nhóm nhà cung cấp', 'G-SUPPLIER', null, 1);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem danh sách nhóm nhà cung cấp', 'G-SUPPLIER-1', 'G-SUPPLIER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Tạo mới nhóm nhà cung cấp', 'G-SUPPLIER-2', 'G-SUPPLIER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xem chi tiết nhóm nhà cung cấp', 'G-SUPPLIER-3', 'G-SUPPLIER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Chỉnh sửa nhóm nhà cung cấp', 'G-SUPPLIER-4', 'G-SUPPLIER', 2);
INSERT INTO apps (title, code, parent, level) VALUES ('Xoá nhóm nhà cung cấp', 'G-SUPPLIER-5', 'G-SUPPLIER', 2);

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

INSERT INTO debt_note_type (code, title) VALUES ('REVENUE', 'Khoản thu');
INSERT INTO debt_note_type (code, title) VALUES ('EXPENSE', 'Khoản chi');

INSERT INTO debt_note_status (code, title) VALUES ('OPEN', 'Chưa thanh toán');
INSERT INTO debt_note_status (code, title) VALUES ('REPAYING', 'Thanh toán 1 phần');
INSERT INTO debt_note_status (code, title) VALUES ('SETTLED', 'Hoàn thành');
INSERT INTO debt_note_status (code, title) VALUES ('OVERDUE', 'Quá hạn');