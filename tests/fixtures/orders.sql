DELETE FROM `order`;
DELETE FROM `client`;

INSERT INTO client (id, name, email, created_at, updated_at) VALUES
  (1, 'Test Client', 'test@example.com', NOW(), NOW());

INSERT INTO `order` (id, client_id, status, created_at, updated_at) VALUES
  (1, 1, 'created', NOW(), NOW()); 