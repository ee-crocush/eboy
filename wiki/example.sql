-- Удаление таблиц, если они существуют
DROP TABLE IF EXISTS
    world_state, actions, areas_neutrals, areas_buildings, areas_heroes,
    areas_units, areas_enemies, areas, enemies, units, heroes, abilities, hero_ability,
    buildings, neutrals, user_resources, resources, users, leagues;

------------------------------------------------------------
-- 1. Лиги, пользователи, ресурсы и их связи
------------------------------------------------------------

-- Таблица leagues (лиги)
CREATE TABLE leagues (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    authority INTEGER DEFAULT 0
);

-- Таблица users (пользователи)
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    league_id INT REFERENCES leagues(id) ON DELETE SET NULL,
    balance DECIMAL(19, 2) DEFAULT 0,
    level SMALLINT DEFAULT 1 CHECK(level > 0)
);

-- Таблица ресурсов
CREATE TABLE resources (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Таблица связей user_resources (ресурсы пользователей)
CREATE TABLE user_resources (
    resource_id INT NOT NULL REFERENCES resources(id),
    user_id INT NOT NULL REFERENCES users(user_id),
    value DECIMAL(19, 2) DEFAULT 0,
    PRIMARY KEY (resource_id, user_id)
);

------------------------------------------------------------
-- 2. Объекты и подтипы
------------------------------------------------------------

-- Таблица objects (базовая информация о всех игровых объектах)
CREATE TABLE objects (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL CHECK(type IN ('neutrals', 'buildings', 'heroes', 'units', 'enemies'))
);

-- Нейтральные объекты (например, точки ресурсов, нейтральные постройки и т.п.)
CREATE TABLE neutrals (
    id SERIAL PRIMARY KEY,
    object_id INT NOT NULL REFERENCES objects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    product_id INT, -- ресурс, который производит
    prod_cof DECIMAL(2, 2), -- коэффициент производительности
    capacity DECIMAL(19, 2) NOT NULL,
    can_gathered BOOL NOT NULL DEFAULT True
);

-- Здания
CREATE TABLE buildings (
    id SERIAL PRIMARY KEY,
    object_id INT NOT NULL REFERENCES objects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    characteristics JSONB NOT NULL,
    level SMALLINT DEFAULT 1 CHECK(level > 0),
    upgrade_price DECIMAL(19, 2) DEFAULT 0
);

-- Способности героев
CREATE TABLE abilities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    characteristics JSONB NOT NULL,
    level SMALLINT DEFAULT 1 CHECK(level > 0)
);

-- Герои
CREATE TABLE heroes (
    id SERIAL PRIMARY KEY,
    object_id INT NOT NULL REFERENCES objects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    characteristics JSONB NOT NULL,
    experience_to_up DECIMAL(19, 2) DEFAULT 0,
    level SMALLINT DEFAULT 1 CHECK(level > 0)
);

-- Таблица hero_ability (связь героев и способностей)
CREATE TABLE hero_ability (
    hero_id INT NOT NULL REFERENCES heroes(id) ON DELETE CASCADE,
    ability_id INT NOT NULL REFERENCES abilities(id) ON DELETE CASCADE,
    PRIMARY KEY (hero_id, ability_id)
);

-- Юниты
CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    object_id INT NOT NULL REFERENCES objects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    characteristics JSONB NOT NULL,
    experience_to_up DECIMAL(19, 2) DEFAULT 0,
    level SMALLINT DEFAULT 1 CHECK(level > 0)
);

-- Враги
CREATE TABLE enemies (
    id SERIAL PRIMARY KEY,
    object_id INT NOT NULL REFERENCES objects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    characteristics JSONB NOT NULL,
    level SMALLINT DEFAULT 1 CHECK(level > 0)
);

------------------------------------------------------------
-- 3. Арены и объекты, находящиеся на арене
------------------------------------------------------------

-- Таблица areas (арены)
CREATE TABLE areas (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    radius SMALLINT NOT NULL CHECK(radius > 0)
);

-- Нейтральные объекты на арене
CREATE TABLE areas_neutrals (
    area_id INT NOT NULL REFERENCES areas(id) ON DELETE CASCADE,
    neutral_id INT NOT NULL REFERENCES neutrals(id) ON DELETE CASCADE,
    characteristics JSONB, -- динамические характеристики, которые могут отличаться от базовых
    coordinates JSONB, -- координаты на арене
    PRIMARY KEY (area_id, neutral_id)
);

-- Здания на арене
CREATE TABLE areas_buildings (
    area_id INT NOT NULL REFERENCES areas(id) ON DELETE CASCADE,
    building_id INT NOT NULL REFERENCES buildings(id) ON DELETE CASCADE,
    characteristics JSONB,
    coordinates JSONB,
    PRIMARY KEY (area_id, building_id)
);

-- Герои на арене
CREATE TABLE areas_heroes (
    area_id INT NOT NULL REFERENCES areas(id) ON DELETE CASCADE,
    hero_id INT NOT NULL REFERENCES heroes(id) ON DELETE CASCADE,
    characteristics JSONB,
    experience DECIMAL(19, 2) DEFAULT 0,
    coordinates JSONB,
    PRIMARY KEY (area_id, hero_id)
);

-- Юниты на арене
CREATE TABLE areas_units (
    area_id INT NOT NULL REFERENCES areas(id) ON DELETE CASCADE,
    unit_id INT NOT NULL REFERENCES units(id) ON DELETE CASCADE,
    characteristics JSONB,
    experience DECIMAL(19, 2) DEFAULT 0,
    coordinates JSONB,
    PRIMARY KEY (area_id, unit_id)
);

-- Враги на арене
CREATE TABLE areas_enemies (
    area_id INT NOT NULL REFERENCES areas(id) ON DELETE CASCADE,
    enemy_id INT NOT NULL REFERENCES enemies(id) ON DELETE CASCADE,
    characteristics JSONB,
    coordinates JSONB,
    PRIMARY KEY (area_id, enemy_id)
);

------------------------------------------------------------
-- 4. Действия и состояние мира
------------------------------------------------------------

-- Таблица actions (действия на арене)
CREATE TABLE actions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    area_id INT NOT NULL REFERENCES areas(id) ON DELETE CASCADE,
    object_source_id INT,  -- ID источника действия (не обязательно внешний ключ, т.к. может ссылаться на разные типы)
    source_type VARCHAR(50) CHECK(source_type IN ('neutrals', 'buildings', 'heroes', 'units', 'enemies')),
    object_dest_id INT,    -- ID цели действия
    dest_type VARCHAR(50) CHECK(dest_type IN ('neutrals', 'buildings', 'heroes', 'units', 'enemies')),
    action_type SMALLINT NOT NULL CHECK(action_type BETWEEN 1 AND 4), -- 1 - MOVE, 2 - ATTACK, 3 - BUILD, 4 - UPGRADE
    start_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    duration INTERVAL DEFAULT '00:00:00',
    status VARCHAR(50) CHECK(status IN ('DONE', 'PROCESS', 'CANCELED'))
);

-- Таблица world_state (состояние мира)
-- Предполагается, что состояние мира будем получать по user_id и последнему time_stamp
CREATE TABLE world_state (
    session_id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    area_id INT NOT NULL REFERENCES areas(id) ON DELETE CASCADE,
    action_id INT NOT NULL REFERENCES actions(id) ON DELETE CASCADE,
    time_stamp TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
