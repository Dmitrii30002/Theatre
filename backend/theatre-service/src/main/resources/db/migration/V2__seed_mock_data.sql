-- ============================================
-- МИГРАЦИЯ: Тестовые данные для API спектаклей
-- ============================================

-- Спектакли из mockSpectaclesResponse.
INSERT INTO theatre.theatres (id, name)
VALUES (1, 'Большой театр')
ON CONFLICT (id) DO NOTHING;

INSERT INTO theatre.spectacles (
    id,
    name,
    description,
    genre,
    age_limit,
    duration_minutes,
    pushkin_card,
    theatre_id,
    preview_url
)
VALUES
    (1, 'Гамлет', 'Трагедия Уильяма Шекспира о принце Датском, стоящем перед выбором между местью, долгом и собственной совестью.', 'Драма', 16, 180, TRUE, 1, 'https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600'),
    (2, 'Щелкунчик', NULL, NULL, NULL, NULL, FALSE, 1, 'https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600'),
    (3, 'Мастер и Маргарита', NULL, NULL, NULL, NULL, FALSE, 1, 'https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600'),
    (4, 'Гамлет2', NULL, NULL, NULL, NULL, FALSE, 1, 'https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600'),
    (5, 'Щелкунчик2', NULL, NULL, NULL, NULL, FALSE, 1, 'https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600'),
    (6, 'Мастер и Маргарита2', NULL, NULL, NULL, NULL, FALSE, 1, 'https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600')
ON CONFLICT (id) DO NOTHING;

-- Площадка и ценовые группы из mockShowDetail.
INSERT INTO theatre.platforms (id, name, scheme)
VALUES (1, 'Большой зал', 'СХЕМА ПОСАДКИ')
ON CONFLICT (id) DO NOTHING;

INSERT INTO theatre.seats_groups (id, color, price)
VALUES
    (1, '#4CAF50', 5000),
    (2, '#2196F3', 3500)
ON CONFLICT (id) DO NOTHING;

-- UUID соответствуют show_001, show_002 и show_003 из mockSpectacleDetail.
INSERT INTO theatre.shows (id, spectacle_id, platform_id, date)
VALUES
    ('00000000-0000-0000-0000-000000000001', 1, 1, '2026-08-25 19:00:00'),
    ('00000000-0000-0000-0000-000000000002', 1, 1, '2026-08-27 19:00:00'),
    ('00000000-0000-0000-0000-000000000003', 1, 1, '2026-08-30 14:00:00')
ON CONFLICT (id) DO NOTHING;

-- В текущем enum booking_status используются доменные эквиваленты:
-- SOLD -> BOOKED, HELD -> RESERVED, BLOCKED -> UNAVAILABLE.
INSERT INTO theatre.seats (id, number, "row", type, group_id, show_id, status)
VALUES
    (10001, 1, 1, 'STANDARD', 1, '00000000-0000-0000-0000-000000000001', 'AVAILABLE'),
    (10002, 2, 1, 'STANDARD', 1, '00000000-0000-0000-0000-000000000001', 'BOOKED'),
    (10003, 3, 1, 'STANDARD', 1, '00000000-0000-0000-0000-000000000001', 'AVAILABLE'),
    (10004, 4, 1, 'STANDARD', 1, '00000000-0000-0000-0000-000000000001', 'RESERVED'),
    (10005, 5, 1, 'STANDARD', 1, '00000000-0000-0000-0000-000000000001', 'AVAILABLE'),
    (10006, 1, 2, 'STANDARD', 2, '00000000-0000-0000-0000-000000000001', 'AVAILABLE'),
    (10007, 2, 2, 'STANDARD', 2, '00000000-0000-0000-0000-000000000001', 'AVAILABLE'),
    (10008, 3, 2, 'STANDARD', 2, '00000000-0000-0000-0000-000000000001', 'BOOKED'),
    (10009, 4, 2, 'STANDARD', 2, '00000000-0000-0000-0000-000000000001', 'AVAILABLE'),
    (10010, 5, 2, 'STANDARD', 2, '00000000-0000-0000-0000-000000000001', 'UNAVAILABLE')
ON CONFLICT (id) DO NOTHING;

INSERT INTO theatre.spectacle_images (id, url, spectacle_id)
VALUES
    (1, 'https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600', 1),
    (2, 'https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600', 1)
ON CONFLICT (id) DO NOTHING;

-- Синхронизируем sequences после явных тестовых ID.
SELECT setval('theatre.theatres_id_seq', GREATEST((SELECT MAX(id) FROM theatre.theatres), 1));
SELECT setval('theatre.spectacles_id_seq', GREATEST((SELECT MAX(id) FROM theatre.spectacles), 1));
SELECT setval('theatre.platforms_id_seq', GREATEST((SELECT MAX(id) FROM theatre.platforms), 1));
SELECT setval('theatre.seats_groups_id_seq', GREATEST((SELECT MAX(id) FROM theatre.seats_groups), 1));
SELECT setval('theatre.seats_id_seq', GREATEST((SELECT MAX(id) FROM theatre.seats), 1));
SELECT setval('theatre.spectacle_images_id_seq', GREATEST((SELECT MAX(id) FROM theatre.spectacle_images), 1));