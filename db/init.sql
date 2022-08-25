CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL,
    name VARCHAR(255),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    id_role BIGINT NOT NULL,
    enable BOOLEAN NOT NULL,
    verification_token VARCHAR(255),
    verification_token_exp TIMESTAMP WITH TIME ZONE,
    otp VARCHAR(6),
    otp_exp TIMESTAMP WITH TIME ZONE,
    locked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY(id),
    CONSTRAINT fk_accounts_role
        FOREIGN KEY(id_role)
            REFERENCES roles(id)
);

CREATE TABLE IF NOT EXISTS genders (
    id BIGSERIAL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS users_details (
    id BIGSERIAL,
    id_user BIGINT,
    fullname VARCHAR(255),
    id_gender BIGINT,
    phone VARCHAR(20),
    date_of_birth DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY(id),
    CONSTRAINT fk_users_details_accounts
        FOREIGN KEY(id_user)
            REFERENCES accounts(id),
    CONSTRAINT fk_users_details_genders
        FOREIGN KEY(id_gender)
            REFERENCES genders(id)
);

CREATE TABLE IF NOT EXISTS articles (
    id BIGSERIAL,
    id_user BIGINT,
    title VARCHAR(255),
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY(id),
    CONSTRAINT fk_articles_accounts
        FOREIGN KEY(id_user)
            REFERENCES accounts(id)
);

CREATE TABLE IF NOT EXISTS tags (
    id BIGSERIAL,
    id_article BIGINT,
    name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY(id),
    CONSTRAINT fk_tags_articles
        FOREIGN KEY(id_article)
            REFERENCES articles(id)
);

CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL,
    name VARCHAR(255),
    PRIMARY KEY(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);


CREATE TABLE IF NOT EXISTS articles_categories (
    id BIGSERIAL,
    id_article BIGINT,
    id_category BIGINT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_articles_categories_articles
        FOREIGN KEY(id_article)
            REFERENCES articles(id),
    CONSTRAINT fk_articles_categories_categories
        FOREIGN KEY(id_category)
            REFERENCES categories(id)    
);

INSERT INTO roles (id, name) VALUES (1, 'admin'), (2, 'user') ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name;
