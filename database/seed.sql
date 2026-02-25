-- Seed data for gowoopi
-- Store
INSERT IGNORE INTO stores (id, name, admin_username, admin_password_hash, default_language, created_at, updated_at) 
VALUES ('00000000-0000-0000-0000-000000000001', '황금치킨', 'admin', '', 'ko', NOW(), NOW());

-- Admin (password: admin123)
INSERT IGNORE INTO admins (store_id, username, password_hash, name, created_at, updated_at) 
VALUES ('00000000-0000-0000-0000-000000000001', 'admin', '$2a$10$WlEUrq.I9xYgoBxg3KUR/OC5dtSNQK4gvCkW0BOWY34VrvZHREfJm', '관리자', NOW(), NOW());

-- Table (password: admin123)
INSERT IGNORE INTO tables (store_id, table_number, password_hash, is_active, created_at, updated_at) 
VALUES ('00000000-0000-0000-0000-000000000001', 1, '$2a$10$WlEUrq.I9xYgoBxg3KUR/OC5dtSNQK4gvCkW0BOWY34VrvZHREfJm', 1, NOW(), NOW());

-- Categories
INSERT INTO categories (store_id, name, sort_order, created_at, updated_at) VALUES 
('00000000-0000-0000-0000-000000000001', '치킨', 1, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', '사이드', 2, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', '음료', 3, NOW(), NOW())
ON DUPLICATE KEY UPDATE name=VALUES(name);

-- Get category IDs and insert menus
SET @chicken_id = (SELECT id FROM categories WHERE store_id='00000000-0000-0000-0000-000000000001' AND name='치킨' LIMIT 1);
SET @side_id = (SELECT id FROM categories WHERE store_id='00000000-0000-0000-0000-000000000001' AND name='사이드' LIMIT 1);
SET @drink_id = (SELECT id FROM categories WHERE store_id='00000000-0000-0000-0000-000000000001' AND name='음료' LIMIT 1);

-- Menus
INSERT INTO menus (store_id, category_id, name, description, price, image_url, is_available, sort_order, created_at, updated_at) VALUES 
-- 치킨
('00000000-0000-0000-0000-000000000001', @chicken_id, '후라이드치킨', '바삭바삭 황금빛 후라이드', 18000, 'https://images.unsplash.com/photo-1626082927389-6cd097cdc6ec?w=400', 1, 1, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @chicken_id, '양념치킨', '달콤 매콤한 특제 양념', 19000, 'https://images.unsplash.com/photo-1575932444877-5106bee2a599?w=400', 1, 2, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @chicken_id, '간장치킨', '짭짤 달콤한 간장 소스', 19000, 'https://images.unsplash.com/photo-1569058242567-93de6f36f8eb?w=400', 1, 3, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @chicken_id, '반반치킨', '후라이드 + 양념 반반', 19000, 'https://images.unsplash.com/photo-1567620832903-9fc6debc209f?w=400', 1, 4, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @chicken_id, '치즈볼치킨', '치즈볼이 가득한 치킨', 21000, 'https://images.unsplash.com/photo-1614398751058-eb2e0bf63e53?w=400', 1, 5, NOW(), NOW()),
-- 사이드
('00000000-0000-0000-0000-000000000001', @side_id, '치즈볼', '쭉쭉 늘어나는 치즈볼 5개', 5000, 'https://images.unsplash.com/photo-1531749668029-2db88e4276c7?w=400', 1, 1, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @side_id, '감자튀김', '바삭한 감자튀김', 4000, 'https://images.unsplash.com/photo-1573080496219-bb080dd4f877?w=400', 1, 2, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @side_id, '치킨무', '새콤달콤 치킨무', 1000, 'https://images.unsplash.com/photo-1583224994076-e3c28d83d0f5?w=400', 1, 3, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @side_id, '콘샐러드', '달콤한 콘샐러드', 3000, 'https://images.unsplash.com/photo-1512621776951-a57141f2eefd?w=400', 1, 4, NOW(), NOW()),
-- 음료
('00000000-0000-0000-0000-000000000001', @drink_id, '콜라 1.25L', '시원한 콜라', 3000, 'https://images.unsplash.com/photo-1622483767028-3f66f32aef97?w=400', 1, 1, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @drink_id, '사이다 1.25L', '청량한 사이다', 3000, 'https://images.unsplash.com/photo-1625772299848-391b6a87d7b3?w=400', 1, 2, NOW(), NOW()),
('00000000-0000-0000-0000-000000000001', @drink_id, '맥주 500ml', '시원한 생맥주', 5000, 'https://images.unsplash.com/photo-1608270586620-248524c67de9?w=400', 1, 3, NOW(), NOW())
ON DUPLICATE KEY UPDATE name=VALUES(name);
