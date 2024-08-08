INSERT INTO users (name, username, password, email, phone, role, status)
VALUES ('TungDuong', 'tungduong', 'securepassword', 'tungduong@example.com', '123-456-7890', 'admin', 'active');


INSERT INTO access_tokens (token, user_id)
VALUES ('1', 'new_gen_user_id');
