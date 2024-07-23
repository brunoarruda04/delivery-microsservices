CREATE TYPE user_role AS ENUM (
  'RESTAURANT_ADMIN',
  'RESTAURANT_USER',
  'USER',
  'DELIVERY'
);

CREATE TABLE users (
  id UUID PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role user_role NOT NULL
);